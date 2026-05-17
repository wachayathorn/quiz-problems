# Booking System API Design

> **Context**: ระบบจองห้องประชุมภายในองค์กร ผู้ใช้ login ก่อนใช้งาน 1 user จองได้หลายห้อง 1 ห้องห้ามเวลาทับซ้อน

---

## 1. Assumptions & Clarified Requirements

ก่อนออกแบบ ต้อง clarify กับ BA/SA — เอกสารฉบับนี้สมมุติคำตอบไว้แล้ว ดังนี้

| ประเด็น | สมมุติฐาน |
|---|---|
| Identity | ใช้ JWT, `userId` ดึงจาก token เท่านั้น (ไม่ใช่จาก body) |
| Time format | ISO 8601 UTC ทุก field, client เป็นคนแปลง timezone |
| Time semantic | Half-open interval `[startTime, endTime)` — `end == start` ของอีกอันไม่ถือว่าทับ |
| ขั้นต่ำ/สูงสุด | ขั้นต่ำ 15 นาที สูงสุด 8 ชั่วโมง จองล่วงหน้าได้ 60 วัน |
| Cancel policy | เจ้าของ booking cancel ได้ก่อน start time, admin cancel ได้ตลอด |
| Delete | ไม่มี hard delete — ใช้ soft cancel (`status = CANCELLED`) เพื่อ audit |
| Concurrency | 2 request พร้อมกัน → ใช้ DB transaction + unique constraint กัน race |
| Recurring booking | ออกนอก scope MVP |

---

## 2. Resource Model

```
Booking
├── id            string (UUID)
├── roomId        string
├── userId        string  (จาก JWT)
├── startTime     timestamp (UTC)
├── endTime       timestamp (UTC)
├── status        enum: ACTIVE | CANCELLED
├── note          string (optional, max 500)
├── createdAt     timestamp
├── updatedAt     timestamp
├── cancelledAt   timestamp (nullable)
└── cancelReason  string    (nullable)
```

---

## 3. Endpoints

ทุก endpoint ต้องมี `Authorization: Bearer <jwt>` ยกเว้นระบุไว้

### 3.1 Create Booking

```
POST /api/v1/bookings
```

Request body:
```json
{
  "roomId": "room-456",
  "startTime": "2025-10-15T10:00:00Z",
  "endTime":   "2025-10-15T11:00:00Z",
  "note": "Sprint planning"
}
```

Success: `201 Created`
```json
{
  "id": "bkg-789",
  "roomId": "room-456",
  "userId": "usr-123",
  "startTime": "2025-10-15T10:00:00Z",
  "endTime":   "2025-10-15T11:00:00Z",
  "status": "ACTIVE",
  "note": "Sprint planning",
  "createdAt": "2025-10-10T08:00:00Z",
  "updatedAt": "2025-10-10T08:00:00Z"
}
```

Error: `400`, `401`, `409 BOOKING_OVERLAP`, `422 INVALID_TIME_RANGE`, `429`, `500`

### 3.2 List Bookings

```
GET /api/v1/bookings?roomId=room-456&from=2025-10-15&to=2025-10-16&status=ACTIVE&page=1&pageSize=20&sort=-startTime
```

Query params (ทั้งหมด optional):

| Param | Type | Note |
|---|---|---|
| `roomId` | string | filter ตามห้อง |
| `userId` | string | filter ตามเจ้าของ (admin เท่านั้น; user ปกติเห็นแค่ของตัวเอง) |
| `from`, `to` | ISO date | ช่วงเวลาที่ booking ทับกับ window นี้ |
| `status` | enum | default `ACTIVE` |
| `page` | int | default 1 |
| `pageSize` | int | default 20, max 100 |
| `sort` | string | เช่น `-startTime`, `createdAt` |

Success: `200 OK`
```json
{
  "data": [ { ...booking } ],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "total": 134
  }
}
```

### 3.3 Get Booking by ID

```
GET /api/v1/bookings/{id}
```

Success: `200`, return Booking object
Error: `401`, `403` (ไม่ใช่เจ้าของและไม่ใช่ admin), `404`

### 3.4 Update Booking (partial)

```
PATCH /api/v1/bookings/{id}
```

ใช้ `PATCH` เพราะมัก update บางฟิลด์ (เช่น เลื่อนเวลา) — `PUT` จะบังคับส่งทุก field

Request body (ทุก field optional):
```json
{
  "startTime": "2025-10-15T10:30:00Z",
  "endTime":   "2025-10-15T11:30:00Z",
  "note": "Postponed by 30 min"
}
```

Rule: เจ้าของเท่านั้น, เปลี่ยน `roomId` ไม่ได้ (ให้ cancel แล้วจองใหม่)

Success: `200`, return updated Booking
Error: `400`, `401`, `403`, `404`, `409`, `422`

### 3.5 Cancel Booking

```
POST /api/v1/bookings/{id}/cancel
```

ใช้ action endpoint แทน `DELETE` เพราะต้องเก็บ audit (soft cancel)

Request body:
```json
{ "reason": "Meeting cancelled by organizer" }
```

Success: `200`, return Booking with `status: CANCELLED`
Error: `401`, `403`, `404`, `409` (ถ้า cancel หลัง start time แล้ว)

---

## 4. Error Response Contract

ทุก error response ใช้ shape เดียวกัน เพื่อให้ client switch logic ได้ด้วย `code`

```json
{
  "error": {
    "code": "BOOKING_OVERLAP",
    "message": "room is already booked in this time range",
    "details": [
      { "field": "startTime", "issue": "overlaps with booking bkg-111" }
    ],
    "requestId": "req-abc-123"
  }
}
```

### Error Code Mapping (`errcode`)

| HTTP | Code | เคสที่เกิด |
|---|---|---|
| 400 | `BAD_REQUEST` | JSON parse fail, field type ผิด, required missing |
| 401 | `UNAUTHORIZED` | ไม่มี / token หมดอายุ / token invalid |
| 403 | `FORBIDDEN` | แก้/ดู booking ของคนอื่นโดยไม่ใช่ admin |
| 404 | `NOT_FOUND` | booking id ไม่มี |
| 409 | `BOOKING_OVERLAP` | เวลาทับกับ booking active อื่น |
| 409 | `ALREADY_CANCELLED` | cancel booking ที่ cancel ไปแล้ว |
| 422 | `INVALID_TIME_RANGE` | `end <= start`, duration นอกขอบเขต, จองในอดีต |
| 429 | `TOO_MANY_REQUESTS` | rate limit |
| 500 | `INTERNAL_SERVER_ERROR` | bug / DB / dependency ล่ม |

> **หลักการ**: `code` คือ contract ห้ามเปลี่ยน, `message` เปลี่ยนได้ (i18n), `details` ช่วย frontend ชี้ field, `requestId` ช่วย correlate log ตอน debug SIT/UAT

---

## 5. Database Schema

```sql
CREATE TABLE bookings (
    id            VARCHAR(36)  NOT NULL PRIMARY KEY,
    room_id       VARCHAR(36)  NOT NULL,
    user_id       VARCHAR(36)  NOT NULL,
    start_time    TIMESTAMP    NOT NULL,
    end_time      TIMESTAMP    NOT NULL,
    status        VARCHAR(16)  NOT NULL DEFAULT 'ACTIVE',
    note          VARCHAR(500),
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL,
    cancelled_at  TIMESTAMP,
    cancel_reason VARCHAR(500),
    CONSTRAINT chk_time CHECK (end_time > start_time)
);

-- composite index สำหรับ overlap check (เป็น query หลัก)
CREATE INDEX idx_bookings_room_time
  ON bookings (room_id, status, start_time, end_time);

-- index สำหรับ list ของ user
CREATE INDEX idx_bookings_user_time
  ON bookings (user_id, start_time);
```

---

## 6. Overlap Detection Logic

สูตร: two intervals `[a, b)` และ `[c, d)` **ทับกัน** เมื่อ `a < d AND c < b`

```sql
-- ตรวจก่อน insert / update
SELECT id
FROM bookings
WHERE room_id    = :roomId
  AND status    = 'ACTIVE'
  AND start_time < :newEndTime    -- existing.start < new.end
  AND end_time   > :newStartTime  -- existing.end   > new.start
  AND id        != :currentId     -- exclude self ตอน update (ส่ง '' ถ้า create)
FOR UPDATE;                       -- lock row กัน race
```

### Concurrency Handling

```
BEGIN TRANSACTION;
  SELECT ... FOR UPDATE;          -- ถ้าเจอ → rollback → 409
  INSERT INTO bookings ...;
COMMIT;
```

**ทางเลือก** (Oracle/MSSQL):

- Oracle: `SELECT ... FOR UPDATE` ใช้ได้
- MSSQL: ใช้ `WITH (UPDLOCK, HOLDLOCK)` บน SELECT
- หรือใช้ **exclusion constraint** (Postgres) / **application-level Redis lock** ถ้า DB ไม่รองรับ

---

## 7. Non-Functional Requirements

| Area | Spec |
|---|---|
| Auth | JWT, expire 1 ชม., refresh token 7 วัน |
| Rate limit | 60 req/min/user สำหรับ POST, 300 req/min/user สำหรับ GET |
| Timeout | request timeout 5 วินาที, DB query timeout 2 วินาที (`context.WithTimeout`) |
| Idempotency | POST รับ header `Idempotency-Key` (UUID จาก client) — กัน double-submit ตอน retry |
| Pagination | offset-based สำหรับ MVP, พิจารณา cursor-based ถ้าข้อมูลโตเกิน 100k |
| Logging | structured JSON log: `requestId`, `userId`, `endpoint`, `status`, `latencyMs` — **ห้าม** log token/PII |
| Metrics | RED metrics: Request rate, Error rate, Duration (p50/p95/p99) |
| Tracing | propagate `traceparent` header สำหรับ distributed tracing |
| API Spec | OpenAPI 3 doc generated + commit ใน repo (ใช้ร่วมกับ BA/QA/Frontend) |

---

## 8. Deployment Considerations

| Env | Note |
|---|---|
| Dev | local DB, mock auth |
| SIT | shared DB, integrate กับ auth service ตัวจริง, ใช้ test data set |
| UAT | data shape ใกล้เคียง prod, ให้ business team ทดสอบ |
| Prod | enable rate limit, alerting, structured log ไป central log (ELK/Datadog) |

Config แยกตาม env ผ่าน environment variables ห้าม hardcode connection string / secret

---

## 9. Open Questions (ฝาก BA/SA)

1. Buffer time ระหว่าง booking ติดกัน (เช่น clean room 15 นาที) ต้องบังคับไหม?
2. ห้องมี capacity ต่างกัน ต้องเช็คจำนวนผู้เข้าร่วมไหม?
3. Recurring booking (จองทุกจันทร์ × 10 สัปดาห์) อยู่ใน scope phase ไหน?
4. Notification (อีเมล / LINE / Teams) หลังสร้าง/ยกเลิก booking — ต้องมีไหม?
5. มี role อะไรบ้าง? (user, admin, room manager?)
6. SLA: API ต้องตอบภายในกี่ ms? expected concurrent user เท่าไหร่?