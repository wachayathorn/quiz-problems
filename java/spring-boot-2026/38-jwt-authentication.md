# #38 — JWT: ของใน token กับ Stateless Security

> **วิดีโอ:** [Spring JWT Authentication | JWT Internals & Stateless Security | Spring Boot Full Course #38](https://youtu.be/pUWp1EuVnGs) — Coder Army, ~82 นาที
> **แหล่งที่ใช้สรุป:** ซับไตเติลภาษาอังกฤษที่ YouTube สร้างอัตโนมัติ
> **เวอร์ชันในโค้ด:** Spring Boot 4.1 + Java 21 (ใส่หมายเหตุไว้ทุกจุดที่ Spring Boot 3.x ต่างกัน)

ตอนนี้ตอบคำถามว่า หลัง login แล้วจะจำผู้ใช้ได้อย่างไรโดยไม่ส่ง username กับ password ทุก request ผู้สอนเริ่มจาก session แบบ stateful ที่ Spring Security ทำให้อยู่แล้ว แล้วค่อยไล่จาก JSON เปล่า ๆ ไปจนถึงลายเซ็นของ JWT และปิดท้ายด้วย login endpoint ที่แจก access token

ส่วนที่ติดป้าย `💡 เพิ่มเติม` คือของที่ผมเติมเอง ไม่ได้อยู่ในวิดีโอ ส่วนที่ติด `⚠️ แก้จากวิดีโอ` คือจุดที่วิดีโอพูดคลาดเคลื่อน

---

## สารบัญ

| เวลา | หัวข้อ |
|---|---|
| [0:00](https://youtu.be/pUWp1EuVnGs?t=0) | [1. โปรเจกต์เดิม และปัญหาของ Basic Auth](#1-โปรเจกต์เดิม-และปัญหาของ-basic-auth) |
| [6:46](https://youtu.be/pUWp1EuVnGs?t=406) | [2. Authentication สองก้อน](#2-authentication-สองก้อน) |
| [8:47](https://youtu.be/pUWp1EuVnGs?t=527) | [3. Stateful: session id ใน cookie](#3-stateful-session-id-ใน-cookie) |
| [15:02](https://youtu.be/pUWp1EuVnGs?t=902) | [4. SessionCreationPolicy](#4-sessioncreationpolicy) |
| [16:23](https://youtu.be/pUWp1EuVnGs?t=983) | [5. ทำไม session ถึงลำบากเมื่อมีหลายเครื่อง](#5-ทำไม-session-ถึงลำบากเมื่อมีหลายเครื่อง) |
| [17:38](https://youtu.be/pUWp1EuVnGs?t=1058) | [6. Stateless: ส่ง token กลับมาทุกครั้ง](#6-stateless-ส่ง-token-กลับมาทุกครั้ง) |
| [21:23](https://youtu.be/pUWp1EuVnGs?t=1283) | [7. ทำไม Base64 อย่างเดียวถึงไม่พอ](#7-ทำไม-base64-อย่างเดียวถึงไม่พอ) |
| [30:17](https://youtu.be/pUWp1EuVnGs?t=1817) | [8. Hash ก็ยังปลอมได้](#8-hash-ก็ยังปลอมได้) |
| [40:43](https://youtu.be/pUWp1EuVnGs?t=2443) | [9. Encryption กับ Signature](#9-encryption-กับ-signature) |
| [49:36](https://youtu.be/pUWp1EuVnGs?t=2976) | [10. JWT สามส่วน และกุญแจสองแบบ](#10-jwt-สามส่วน-และกุญแจสองแบบ) |
| [55:03](https://youtu.be/pUWp1EuVnGs?t=3303) | [11. Header, claims และการประกอบ token](#11-header-claims-และการประกอบ-token) |
| [59:47](https://youtu.be/pUWp1EuVnGs?t=3587) | [12. Dependency, secret และ Bean](#12-dependency-secret-และ-bean) |
| [1:06:26](https://youtu.be/pUWp1EuVnGs?t=3986) | [13. SecurityFilterChain แบบ stateless](#13-securityfilterchain-แบบ-stateless) |
| [1:08:27](https://youtu.be/pUWp1EuVnGs?t=4107) | [14. Login endpoint กับ JwtService](#14-login-endpoint-กับ-jwtservice) |
| [1:18:35](https://youtu.be/pUWp1EuVnGs?t=4715) | [15. ยิง Bearer token](#15-ยิง-bearer-token) |
| — | [16. สรุปท้ายตอน](#16-สรุปท้ายตอน) |

---

## 1. โปรเจกต์เดิม และปัญหาของ Basic Auth

[▶ 0:00](https://youtu.be/pUWp1EuVnGs?t=0)

ผู้สอนย้ำว่าตอนแรกของ Spring Security อธิบาย stateful กับ stateless ไว้แล้ว และตอนที่แล้วทำ login flow ไว้แล้ว ตอนนี้ต่อจากโปรเจกต์นั้น

ในโปรเจกต์มี Entity สองตัวที่ผูก many-to-many กัน

| Entity | ฟิลด์ | ตาราง |
|---|---|---|
| `User` | id, username, password, enabled, roles | `user` |
| `Role` | id, name | `role` |
| ตารางเชื่อม | user id กับ role id | `user_roles` |

`RoleController` ที่ `POST /api/roles` มีหน้าที่เดียวคือบันทึก role ลงฐานข้อมูล ตอนก่อนหน้าผู้สอนใส่ `ROLE_USER` กับ `ROLE_ADMIN` ไว้แล้ว `RoleService.addRole` เรียก `RoleRepository.save` โดยไม่ตรวจค่าเพิ่ม

`UserController` อยู่ที่ `/api/user` มีสอง endpoint ที่ยังใช้ในตอนนี้

- `GET /hello` คืนข้อความทักทาย
- `POST /register` รับ username กับ password แล้วคืน username กับข้อความว่าสมัครสำเร็จ

`AuthService.register` สร้าง `User` ใส่ username จาก request เข้ารหัส password ด้วย `PasswordEncoder` ตั้ง `enabled` เป็น true ดึง role ชื่อ `ROLE_USER` จาก repository แล้ว `save` ไม่รับ role จาก client เพราะ client จะอ้างเองว่าเป็น admin ได้

`CustomUserDetails` ในแพ็กเกจ entity ห่อ `User` ให้เป็น `UserDetails` ของ Spring Security โดยส่งต่อ `getUsername`, `getPassword`, `isEnabled` และแปลงแต่ละ `Role.name` เป็น `SimpleGrantedAuthority` method ที่เหลือของ `UserDetails` คืน true อยู่แล้ว จึงไม่ override

`CustomUserDetailsService` เป็น `@Service` ที่ implement `UserDetailsService` method เดียวคือ `loadUserByUsername` มันเรียก `UserRepository.findByUsername` ถ้าไม่เจอจะโยน `UsernameNotFoundException` พร้อมข้อความ `User not found` ถ้าเจอจะคืน `new CustomUserDetails(user)`

> 💡 **เพิ่มเติม:** many-to-many ของ `roles` เป็น lazy ตามค่า default ของ JPA ตอนที่แล้ว `getAuthorities` เลยเจอ `LazyInitializationException` ผู้สอนใส่ `@EntityGraph(attributePaths = "roles")` ที่ `findByUsername` เพื่อให้โหลด role พร้อม user โค้ดด้านล่างใส่ annotation นี้ไว้ด้วย แม้ตอนนี้จะไม่ได้พูดถึงซ้ำ

`SecurityConfig` จากตอนที่แล้วมี `PasswordEncoder` เป็น `BCryptPasswordEncoder`, `DaoAuthenticationProvider` ที่รับ `UserDetailsService` กับ encoder และ `SecurityFilterChain` ที่ปิด CSRF, เปิด `formLogin` กับ `httpBasic`, `permitAll` เฉพาะ `POST /api/user/register` และให้ที่เหลือ `authenticated` ผู้สอนบอกว่า bean ของ provider ไม่จำเป็น เพราะถ้ามี `UserDetailsService` กับ `PasswordEncoder` อยู่แล้ว Spring สร้างให้เอง

[▶ 5:13](https://youtu.be/pUWp1EuVnGs?t=313)

เขาแก้ `GET /hello` ให้รับ `Authentication` แล้วคืนชื่อคนที่ login อยู่ `Authentication` ก้อนนี้ถูกใส่เข้า `SecurityContext` โดย filter ตอนที่ตรวจตัวตนผ่าน

```java
@GetMapping("/hello")
public String hello(Authentication authentication) {
    return "Hello, you're logged in as " + authentication.getName();
}
```

จากนั้นยิง `GET /api/user/hello` ใน Postman ด้วย Basic Auth ของผู้ใช้ `Aditya` รหัส `secret123` ได้ข้อความว่า login เป็น Aditya ถอด Basic Auth ออกแล้วยิงซ้ำ ได้ **401 Unauthorized** เพราะแอปไม่ได้จำ request ก่อนหน้า

> 💡 **เพิ่มเติม:** ซับเขียนว่า "locked in" ทั้งตอน คำที่ผู้สอนพูดและที่ endpoint คืนคือ logged in

---

## 2. Authentication สองก้อน

[▶ 6:46](https://youtu.be/pUWp1EuVnGs?t=406)

ทุกครั้งที่ client เรียก endpoint ที่ต้องยืนยันตัว server ต้องมี `Authentication` ให้ request นั้น ผู้สอนแยกเป็นสองช่วง

**Authentication request** เกิดตอน client เพิ่งอ้างตัว ยังไม่ผ่านการตรวจ ในตัวอย่างของ Basic Auth ก้อนนี้มี

| ส่วน | ค่าในตัวอย่าง |
|---|---|
| principal | `Aditya` คนที่อ้างว่าเป็นตัวเอง |
| credentials | `secret123` |
| authorities | ว่าง |
| authenticated | `false` |

**Authentication response** เกิดหลัง server ตรวจผ่านแล้ว ก้อนนี้ไม่เก็บรหัสผ่าน มี `UserDetails` ของ Aditya, authorities ที่ใส่ role แล้ว เช่น `ROLE_USER` และ `authenticated = true` จากนั้น request จึงเข้า `GET /hello` ได้

HTTP แต่ละครั้งไม่รู้เรื่องครั้งก่อน server จึงมองทุก request เป็นของใหม่ ไม่ว่าจะ login ไปแล้วหรือยัง นี่คือเหตุผลที่ Basic Auth ต้องแนบ username กับ password ทุกครั้ง

---

## 3. Stateful: session id ใน cookie

[▶ 8:47](https://youtu.be/pUWp1EuVnGs?t=527)

ทาง stateful คือหลังตรวจผ่านแล้ว server เก็บ Authentication response ไว้เอง แล้วออก **session id** ที่สุ่มขึ้นมา เช่น `ASDF123` ส่งกลับให้ client ใน cookie ครั้งถัดไป client ส่ง cookie นั้นมาด้วย server เอา id ไปหา Authentication ก้อนเดิม ถ้าเจอ ก็ไม่ต้องให้กรอกรหัสอีก

```mermaid
sequenceDiagram
    participant C as Client
    participant S as Server
    C->>S: username + password
    Note over S: สร้าง Authentication<br/>เก็บคู่กับ session id
    S-->>C: Set-Cookie: session id
    C->>S: GET /api/user/hello + cookie
    Note over S: เจอ id เดิม<br/>โหลด Authentication ก้อนเก่า
    S-->>C: Hello, you're logged in as Aditya
```

สถาปัตยกรรมนี้เรียก stateful เพราะ server ยังถือ state ของคนที่ login ไว้ ไม่ว่าจะอยู่ใน memory หรือฐานข้อมูล session id เองเป็นสตริงสุ่ม อ่านชื่อหรือ role จากมันไม่ได้ ของที่แปลความหมายได้อยู่ฝั่ง server

[▶ 11:06](https://youtu.be/pUWp1EuVnGs?t=666)

ผู้สอนบอกว่าค่า default ของ Spring Security **สร้าง session อยู่แล้ว** แค่ Postman ไม่ได้ส่ง cookie กลับไป จึงถูกขอรหัสทุกครั้ง เปิดแท็บ Cookies ของ request ที่ยิงสำเร็จจะเห็น session id

[▶ 11:54](https://youtu.be/pUWp1EuVnGs?t=714)

ใน browser เปิด `GET /api/user/hello` โดยยังไม่ login จะถูกพาไปหน้า login ที่ Spring Security สร้างให้ กรอก `Aditya` / `secret123` แล้วเข้า hello ได้ เปิดแท็บใหม่ที่ endpoint เดิม browser แนบ cookie ให้เอง จึงไม่ถามรหัสอีก ลบ cookie นั้นแล้ว refresh จะกลับไปหน้า login กด login ใหม่จะได้ session id ก้อนใหม่

[▶ 14:16](https://youtu.be/pUWp1EuVnGs?t=856)

อีกทางคือ endpoint `logout` ที่ Spring มีให้ พอเรียกแล้วผู้ใช้ถูกส่งกลับหน้า login และ hello ถามรหัสอีกครั้ง cookie เก่ายังโผล่ในเครื่องมือของ browser ได้ แต่ใช้ login ต่อไม่ได้ session ใหม่จะผูกกับ Authentication รอบใหม่

> ⚠️ **แก้จากวิดีโอ:** ผู้สอนเปิด `/logout` จากแถบที่อยู่ ซึ่งเป็น GET แล้วถูกพากลับหน้า login ใน Spring Security 6 และ 7 matcher ค่า default ของ `LogoutConfigurer` คือ **POST** `/logout` ไม่ใช่ GET การพิมพ์ URL ในแถบที่อยู่จึงไม่ได้เรียก logout handler โดยตรง ฟอร์ม logout ของหน้า login เป็น POST

> 💡 **เพิ่มเติม:** ชื่อ cookie ของ session บน Tomcat คือ `JSESSIONID` ผู้สอนเรียกมันว่า session id

ผู้สอนสรุปว่าพฤติกรรมทั้งก้อนนี้ไม่ได้เขียนเอง มันมาจากค่า default ของ Spring Security

---

## 4. SessionCreationPolicy

[▶ 15:02](https://youtu.be/pUWp1EuVnGs?t=902)

ใน `SecurityFilterChain` เรียก `sessionManagement` แล้วตั้ง `sessionCreationPolicy` ได้ enum `SessionCreationPolicy` มีสี่ค่า

| ค่า | ความหมายที่ผู้สอนใช้ |
|---|---|
| `ALWAYS` | สร้าง session ทุกกรณี |
| `NEVER` | ไม่สร้าง session id Authentication ถูกทิ้งเมื่อ request จบ เพราะไม่มีที่เก็บ |
| `IF_REQUIRED` | สร้างเมื่อการยืนยันตัวตนต้องการ ค่านี้คือ default |
| `STATELESS` | ไม่ใช้ session ใช้คู่กับ JWT |

```java
.sessionManagement(session -> session
        .sessionCreationPolicy(SessionCreationPolicy.IF_REQUIRED))
```

ตอนยิงจาก browser ระบบสร้าง session เพราะ form login ต้องการ session ถ้าเปลี่ยนเป็น `NEVER` จะไม่สร้าง session id ให้ request นั้น

> 💡 **เพิ่มเติม:** `IF_REQUIRED` เป็นค่า default จริงใน `SessionManagementConfigurer` ของ Spring Security 7.1 ส่วน `NEVER` ยังใช้ session ที่มีอยู่แล้วได้ ถ้า client ส่ง cookie เก่ามา `STATELESS` ไม่สร้างและไม่ใช้ session เลย คู่นี้คือค่าที่ตอนหลังตั้งให้ JWT

---

## 5. ทำไม session ถึงลำบากเมื่อมีหลายเครื่อง

[▶ 16:23](https://youtu.be/pUWp1EuVnGs?t=983)

session ทำให้ server ต้องจำ Authentication ไว้ ถ้าแอปมีหลาย instance แล้ว request ถัดไปไปตกอีกเครื่อง เครื่องนั้นไม่มีก้อน Authentication นี้ การยัดคนเดิมให้กลับมาเครื่องเดิมเรียกว่า sticky session อีกทางคือเอา session ไปไว้ที่เก็บส่วนกลาง เพื่อให้ทุกเครื่องอ่านได้

> ⚠️ **แก้จากวิดีโอ:** ผู้สอนเรียกการรันหลายเครื่องว่า data replication ของที่เขาอธิบายคือปัญหา session ไม่ตาม request ข้าม instance ไม่ใช่ชื่อเรียกของการมี server หลายตัว

เขาบอกว่า stateful ไม่ได้ด้อยกว่า stateless เสมอไป ข้อดีของ stateful คือเร็วกว่า เพราะไม่ต้องสร้าง Authentication ใหม่ทุก request ของใน session ถูกหยิบกลับมาใช้

---

## 6. Stateless: ส่ง token กลับมาทุกครั้ง

[▶ 17:38](https://youtu.be/pUWp1EuVnGs?t=1058)

แบบ stateless ไม่เก็บ Authentication หลัง request จบ ตอน login server สร้าง Authentication แล้วคืน **token** ให้ client เก็บเอง request ถัดไป client ส่ง token นั้นใน header server อ่านค่าใน token แล้วสร้าง Authentication ใหม่ ทิ้งก้อนเก่าไปแล้ว ไม่ได้เอา token ไปแมปกับของที่เก็บไว้

```mermaid
sequenceDiagram
    participant C as Client
    participant A as Server เครื่องไหนก็ได้
    C->>A: POST /auth/login username + password
    Note over A: ตรวจรหัส สร้าง token แล้วทิ้ง Authentication
    A-->>C: access token
    C->>A: GET /api/user/hello<br/>Authorization: Bearer token
    Note over A: อ่าน token แล้วสร้าง Authentication ใหม่
    A-->>C: Hello, you're logged in as Aditya
```

token จึงต้องมีข้อมูลพอให้สร้าง Authentication ได้ เช่น ชื่อและ role และต้องมีวิธีรู้ว่า client ไม่ได้แก้ค่าเหล่านั้น ชื่อที่ใช้อยู่ทั่วไปของ token แบบนี้คือ **JWT** ย่อจาก JSON Web Token

เพราะไม่มี state ฝั่ง server client ส่ง request ไป replica ไหนก็ได้ ทุกเครื่องสร้าง Authentication จาก token ก้อนเดียวกัน

---

## 7. ทำไม Base64 อย่างเดียวถึงไม่พอ

[▶ 21:23](https://youtu.be/pUWp1EuVnGs?t=1283)

ผู้สอนไล่จากหลักการ ไม่เริ่มที่ไลบรารี login ที่เขาใช้เป็นตัวอย่างคือ `POST /login` พร้อม JSON ของ username `Aditya` และ password `secret123` server ตรวจแล้วต้องคืน token ที่สร้าง Authentication กลับมาได้

ก้อนแรกที่นึกออกคือ JSON

```json
{"sub": "Aditya", "role": "User"}
```

`sub` คือ subject คนที่เป็นเจ้าของ token `role` คือสิทธิ์ ตอนสร้าง Authentication ฝั่ง response server ใส่ `sub` เป็น principal, ใส่ role เป็น authorities, ไม่ใส่ credentials และตั้ง `authenticated` เป็น true

JSON ทั้งก้อนยัดใส่ header ไม่สะดวก เพราะมีช่องว่างและรูปแบบที่ไม่คงที่ ทางต่อคือเข้ารหัสเป็น **Base64URL** ได้สตริงที่ใส่ header ได้

Base64URL เป็น **encoding** ไม่ใช่การเข้ารหัสลับ แปลงไปและแปลงกลับได้โดยไม่ต้องมีกุญแจ ใครได้สตริงก็อ่าน JSON ได้ และใครมี JSON ก็สร้างสตริงเดิมได้ server ถอดกลับเป็น JSON แล้วสร้าง Authentication ได้ ปัญหาแรกเรื่องการขน JSON จึงจบ

[▶ 27:13](https://youtu.be/pUWp1EuVnGs?t=1633)

ปัญหาถัดไปคือ server ไม่ได้เก็บ token ไว้เทียบ ถ้า client ถอด Base64URL เปลี่ยน `role` จาก `User` เป็น `Admin` แล้วเข้ารหัสใหม่ สตริงจะเปลี่ยน server ถอดได้ JSON ที่ client เขียนเอง และไม่มีหลักฐานว่าค่าถูกแก้ เพราะตอนนี้ยังไม่มีอะไรให้ตรวจ claim ที่ว่า "ฉันคือ Aditya และเป็น admin"

---

## 8. Hash ก็ยังปลอมได้

[▶ 30:17](https://youtu.be/pUWp1EuVnGs?t=1817)

ทางที่ลองต่อคือ hash เพราะ hash เป็นทางเดียว ย้อนจากสตริงกลับเป็น JSON ไม่ได้ แต่ถ้าส่งแต่ hash server ก็อ่านชื่อกับ role ไม่ได้เช่นกัน

เขาเลยส่งสองส่วนคั่นด้วยจุด

```text
base64url(json) + "." + hash(json)
```

ตอนตรวจ server ถอดส่วนหน้า แล้ว hash ใหม่ ถ้าตรงกับส่วนหลัง แสดงว่า payload ไม่ถูกแก้ ถ้า client เปลี่ยน role แล้วเข้ารหัสใหม่โดยไม่แตะ hash เดิม ค่า hash ที่ server คำนวณจะไม่ตรง server ปฏิเสธได้

[▶ 37:34](https://youtu.be/pUWp1EuVnGs?t=2254)

วิธีนี้ยังพัง เพราะอัลกอริทึม hash มีอยู่จำกัด เช่น SHA-256 และใครก็คำนวณได้โดยไม่ต้องมีกุญแจ client ที่รู้ว่า server ใช้ SHA-256 จะเปลี่ยน role, เข้ารหัสใหม่, แล้ว hash ใหม่ทั้งก้อน ส่วนหลังจึงตรงกับส่วนหน้าที่ถูกแก้ server เทียบแล้วคิดว่าไม่มีใครแตะ token

hash จึงพิสูจน์ไม่ได้ว่าคนที่แก้คือคนที่ถือความลับของ server

---

## 9. Encryption กับ Signature

[▶ 40:43](https://youtu.be/pUWp1EuVnGs?t=2443)

ขั้นถัดไปคือ encryption ล็อกข้อความด้วยกุญแจ ได้สตริงที่อ่านกลับได้ก็ต่อเมื่อมีกุญแจถอด ถ้า server เข้ารหัส JSON แล้วส่งให้ client, client ส่งสตริงเดิมกลับมา server ถอดด้วยกุญแจของตัวเองได้ client ที่ไม่มีกุญแจแก้ข้างในไม่ได้

ผู้สอนบอกว่าแบบนี้คือ **JWE** (JSON Web Encryption) ใช้ได้ แต่ JWT ที่ใช้กันทั่วไปไม่ได้เข้ารหัสทั้งก้อน มันใช้ **signature**

ลายเซ็นต่างจาก encryption ตรงที่ payload ยังอ่านได้ แต่ถ้าข้อความเปลี่ยน server จะรู้ ขั้นตอนคือ

1. เข้ารหัส JSON เป็น Base64URL
2. เซ็นสตริงนั้นด้วย secret key ได้ลายเซ็นอีกสตริง
3. ส่ง `payloadที่เข้ารหัส + "." + ลายเซ็น`

client อ่านและแก้ส่วนหน้าได้ แต่คำนวณลายเซ็นใหม่ไม่ได้เพราะไม่มี secret key กุญแจแบบนี้มีพื้นที่ให้เดาใหญ่มาก ผู้สอนยกตัวอย่างความยาว 16 ตัวอักษรเป็นภาพ ไม่ใช่ขนาดที่โค้ดใช้

พอ client ส่งส่วนหน้าที่ถูกแก้คู่กับลายเซ็นเก่า server เซ็นส่วนหน้าด้วยกุญแจตัวเอง ได้ลายเซ็นคนละค่า จึงรู้ว่า token ถูกแก้ และให้ไปยืนยันตัวใหม่

นี่คือจุดที่ต่างจาก hash เปล่า ๆ hash ใครก็ทำใหม่ได้ ลายเซ็นต้องมีกุญแจที่ server ไม่แจก

---

## 10. JWT สามส่วน และกุญแจสองแบบ

[▶ 49:36](https://youtu.be/pUWp1EuVnGs?t=2976)

JWT ที่ใช้จริงมีสามส่วนคั่นด้วยจุด สองส่วนแรกที่อธิบายไปคือ payload กับ signature JWT เติม header ไว้ข้างหน้า

```text
base64url(header).base64url(payload).base64url(signature)
```

ทั้งสามส่วนเป็น Base64URL header และ payload เป็น JSON ที่เข้ารหัส signature เป็นไบต์ของลายเซ็นที่เข้ารหัสต่อ ไม่มีส่วนไหนถูก encrypt ในแบบ HS256 ที่ตอนนี้ใช้

[▶ 50:43](https://youtu.be/pUWp1EuVnGs?t=3043)

การเซ็นมีสองแบบ

**Symmetric** ใช้ secret key ตัวเดียวทั้งเซ็นและตรวจ เร็ว เพราะมีกุญแจเดียว ปัญหาคือทุก service ที่ต้องตรวจ token ต้องถือกุญแจตัวเดียวกัน ถ้า service หนึ่งรั่ว ผู้โจมตีเซ็น token ปลอมได้ แล้วกลับไปปัญหาเดียวกับ hash

**Asymmetric** ใช้คนละดอก private key เซ็นอย่างเดียว public key ตรวจอย่างเดียว public key แจกได้ เซ็นด้วยมันไม่ได้ ภาพที่ผู้สอนวาดคือ service แรกถือ private key ส่วน service ที่เหลือถือแต่ public key แล้วตรวจลายเซ็น

```mermaid
flowchart LR
    S1["Service ที่ login<br/>ถือ private key"] -->|เซ็น token| C["Client"]
    C -->|Bearer token| S2["Service อื่น<br/>ถือ public key"]
    C -->|Bearer token| S3["Service อื่น<br/>ถือ public key"]
```

เขาบอกว่าของ production มักเลือก asymmetric เพราะปลอดภัยกว่าเมื่อมีหลาย service ส่วน symmetric เร็วกว่าและดูแลกุญแจดอกเดียว เดโมนี้ใช้ symmetric เพื่อให้สั้น แล้วสลับไป asymmetric ภายหลังได้โดยเปลี่ยนที่การถือกุญแจ

---

## 11. Header, claims และการประกอบ token

[▶ 55:03](https://youtu.be/pUWp1EuVnGs?t=3303)

Header มีคำสั่งว่าเซ็นด้วยอัลกอริทึมอะไร ตัวอย่างของ symmetric คือ

```json
{"alg": "HS256", "typ": "JWT"}
```

`alg` คือ algorithm `typ` คือ type ซึ่งโดยมากเป็น `JWT` ใส่ `kid` (key id) เป็นสตริงสุ่มได้ แต่เขาบอกว่าไม่จำเป็น เดโมนี้ใช้แค่สองฟิลด์นี้

[▶ 55:48](https://youtu.be/pUWp1EuVnGs?t=3348)

Payload คือข้อมูลที่เอาไปสร้าง Authentication ได้ คีย์เหล่านี้ JWT กำหนดชื่อไว้ให้ ใช้หรือไม่ใช้ก็ได้

| คีย์ | ความหมายในวิดีโอ | ตัวอย่าง |
|---|---|---|
| `sub` | subject ตัวผู้ใช้ | `Aditya` |
| `iss` | issuer คนออก token | `Coder Army API` |
| `aud` | audience token นี้ออกให้ใคร | course API |
| `role` หรือ authorities | สิทธิ์ | `ROLE_USER` |
| `iat` | issued at เวลาที่ออก เป็น epoch | ตัวเลขวินาที |
| `exp` | เวลาหมดอายุ เป็น epoch | ตัวเลขวินาที |

เพิ่มคีย์ของตัวเองได้ แต่เขาเตือนว่าอย่าใส่เยอะ และห้ามใส่รหัสผ่านหรือข้อมูลลับ payload ถูกแค่เข้ารหัส ใครก็อ่านได้ ประโยชน์ของ JWT ไม่ใช่การปิดบังข้อมูล แต่คือลายเซ็น ถ้า header หรือ payload ถูกแก้ การตรวจลายเซ็นจะไม่ผ่าน แล้วให้ client ไป login ใหม่

[▶ 57:34](https://youtu.be/pUWp1EuVnGs?t=3454)

ลำดับประกอบ token ตามที่เขาเขียนบนกระดาน

1. `base64url(header) + "." + base64url(payload)` คู่นี้เรียกว่า signing input
2. เซ็น signing input ด้วยกุญแจ ถ้า symmetric ใช้ secret key ถ้า asymmetric ใช้ private key
3. เข้ารหัสลายเซ็นเป็น Base64URL แล้วต่อท้ายด้วยจุด
4. สตริงสามส่วนนี้คือ JWT ส่งให้ client แล้วให้ client ส่งกลับมา

ใน Spring ไม่ต้องเขียน Base64 กับ HMAC เอง คลาสของ framework ทำให้ ส่วนที่ต้องรู้คือคลาสไหนเซ็น และคลาสไหนตรวจ

---

## 12. Dependency, secret และ Bean

[▶ 59:47](https://youtu.be/pUWp1EuVnGs?t=3587)

ผู้สอนเพิ่ม dependency ชื่อ OAuth2 Resource Server จาก Maven เขาบอกตรง ๆ ว่าตอนนี้ยังไม่ทำ OAuth จะทำในวิดีโอถัดไป dependency นี้ถูกใช้เพราะมีตัวเข้ารหัสและตัวตรวจ JWT ให้

> ⚠️ **Boot 4 vs 3:** ชื่อใน Boot 3 คือ `spring-boot-starter-oauth2-resource-server` Boot 4 เปลี่ยนเป็น `spring-boot-starter-security-oauth2-resource-server` ชื่อเก่ายังอยู่แต่ถูก deprecate โค้ดด้านล่างใช้ชื่อใหม่ starter ของเว็บในโปรเจกต์นี้เป็น `spring-boot-starter-webmvc` ตาม Boot 4 เช่นกัน
> ([Spring Boot 4.0 Migration Guide](https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-4.0-Migration-Guide))

[▶ 1:00:21](https://youtu.be/pUWp1EuVnGs?t=3621)

ใน `application.properties` เขาใส่สามค่า

- `jwt.secret` สร้างด้วย `openssl rand -base64 32`
- `jwt.issuer` เป็น `Coder Army API`
- `jwt.expiry` เป็น `900` หน่วยวินาที คือ 15 นาที

> ⚠️ **แก้จากวิดีโอ:** เขากล่าวถึงคำสั่งนั้นว่าอยากได้จำนวน 32 บิต คำสั่ง `openssl rand -base64 32` สุ่ม **32 ไบต์** ซึ่งคือ 256 บิต พอดีกับขั้นต่ำของ HS256 กุญแจ 32 บิตใช้เซ็น HS256 ไม่ได้

> ⚠️ **แก้จากวิดีโอ:** ตอนลงมือเขาเรียกขั้นนี้ว่า symmetric encryption ของที่ใช้จริงคือลายเซ็น HMAC-SHA256 payload ยังอ่านได้ ไม่ได้ถูกเข้ารหัสแบบ JWE ที่อธิบายไว้ก่อนหน้า

เขาบอกว่าของจริงห้ามใส่ secret นี้ในไฟล์ที่ถูก push ควรอยู่ใน environment variable ใส่ใน properties เฉพาะตอนฝึก

> 💡 **เพิ่มเติม:** วิดีโอใช้ MySQL ผ่าน DBeaver พอร์ต 8080 เดโมด้านล่างใช้ H2 ใน memory และพอร์ต `9038` ตาราง `user` บน H2 ใช้ไม่ได้เพราะ `USER` เป็นคำสงวน จึงตั้งชื่อตารางเป็น `users` วิดีโอให้ Hibernate สร้างตาราง `user` บน MySQL role `ROLE_USER` กับ `ROLE_ADMIN` ถูกใส่ไว้แล้วจากตอนก่อน เดโมนี้สร้างสอง role นั้นตอนแอปสตาร์ต

```properties
server.port=9038

spring.datasource.url=jdbc:h2:mem:studentdb
spring.datasource.driver-class-name=org.h2.Driver
spring.jpa.hibernate.ddl-auto=update

jwt.secret=YDOBRNw9tQCjrsdw7P470psu9oM1mcT+gRWFKARjbKc=
jwt.issuer=Coder Army API
jwt.expiry=900
```

[▶ 1:01:52](https://youtu.be/pUWp1EuVnGs?t=3712)

Bean แรกแปลงสตริงใน properties เป็น `SecretKey` ถอด Base64 แล้วห่อด้วย `SecretKeySpec` อัลกอริทึมของ JCA คือ `HmacSHA256`

```java
@Bean
public SecretKey jwtSecretKey(@Value("${jwt.secret}") String secret) {
    byte[] decodedKey = Base64.getDecoder().decode(secret);
    return new SecretKeySpec(decodedKey, "HmacSHA256");
}
```

[▶ 1:03:13](https://youtu.be/pUWp1EuVnGs?t=3793)

Bean ถัดไปคือ `AuthenticationManager` เขาบอกว่ามันเป็น interface ที่มี provider หลายตัว ตอนที่แล้วใช้ `DaoAuthenticationProvider` สำหรับ username กับ password ตอนนี้ต้องเรียก `authenticate` เองจาก login endpoint จึงเปิด bean นี้จาก `AuthenticationConfiguration`

```java
@Bean
public AuthenticationManager authenticationManager(AuthenticationConfiguration configuration) throws Exception {
    return configuration.getAuthenticationManager();
}
```

[▶ 1:03:35](https://youtu.be/pUWp1EuVnGs?t=3815)

`JwtAuthenticationConverter` บอกว่าจะอ่าน role จาก claim ไหน ผู้สอนสร้าง `JwtGrantedAuthoritiesConverter` แล้ว

- `setAuthoritiesClaimName("authorities")` อ่านลิสต์จาก claim ชื่อ `authorities`
- `setAuthorityPrefix("")` ไม่เติม prefix เอง
- ส่ง converter นี้เข้า `JwtAuthenticationConverter.setJwtGrantedAuthoritiesConverter`

เขาบอกว่าถ้าไม่ทำแบบนี้ framework จะเติม `ROLE_` ข้างหน้า role ของเรา เช่น `Admin` จะกลายเป็น `ROLE_Admin` ถ้าแอปไม่ได้ใช้ role ก็ข้าม bean นี้ได้

> ⚠️ **แก้จากวิดีโอ:** ค่า default ของ `JwtGrantedAuthoritiesConverter` ใน Spring Security 7.1 ไม่ได้เติม `ROLE_` และไม่ได้หยิบ claim ที่เป็น role อยู่แล้ว prefix ค่า default คือ `SCOPE_` และ claim ที่มองหาคือ `scope` กับ `scp` ถ้าไม่ตั้งชื่อ claim เป็น `authorities` สิทธิ์ใน token จะไม่ถูกอ่าน ไม่ได้ถูกเขียนใหม่เป็น `ROLE_Admin` ค่า `SCOPE_` กับ claim `scope`/`scp` เป็นแบบนี้มาตั้งแต่ resource server ของ Spring Security 6 ด้วย

[▶ 1:04:55](https://youtu.be/pUWp1EuVnGs?t=3895)

อีกสอง bean คือคนเซ็นกับคนตรวจ ทั้งคู่รับ `SecretKey` ก้อนเดียวกัน

- `JwtEncoder` เซ็น token implementation ที่เขาเปิดให้ดูคือ `NimbusJwtEncoder`
- `JwtDecoder` ตรวจ token implementation คือ `NimbusJwtDecoder`

พอ bean สองตัวนี้อยู่ใน context โค้ดที่เซ็นและ filter ที่ตรวจจะหยิบไปใช้ เขาบอกว่าเข้าไปอ่านคลาสข้างในได้ แต่ไม่จำเป็น ถ้ารู้ว่า encoder เซ็นและ decoder ตรวจก็พอ

> 💡 **เพิ่มเติม:** ผู้สอนพูดถึง `new NimbusJwtEncoder` ที่รับ secret key แล้วให้ `encode` จาก claims อย่างเดียว บน Spring Security 7.1.1 constructor `NimbusJwtEncoder(JWKSource)` ตั้ง default algorithm เป็น **RS256** ลองแบบนั้นแล้วได้ `JwtEncodingException: Failed to select a JWK signing key` เพราะกุญแจ HMAC จับกับ RS256 ไม่ได้ โค้ดที่รันได้จึงใช้ `NimbusJwtEncoder.withSecretKey` และส่ง `JwsHeader` ของ `HS256` ตอนเข้ารหัส `NimbusJwtDecoder.withSecretKey` ใช้ HS256 เป็นค่า default อยู่แล้ว factory นี้เติม `kid` ให้ header ด้วย แม้บนกระดานจะมีแค่ `alg` กับ `typ`

---

## 13. SecurityFilterChain แบบ stateless

[▶ 1:06:26](https://youtu.be/pUWp1EuVnGs?t=3986)

เขาเอา form login กับ HTTP Basic ออก เพราะต่อไปผู้ใช้ถือ JWT ไม่ได้ส่งรหัสทุก request Authentication ที่ resource server สร้างจาก token คือ JWT authentication token และ provider ของมันคือ JWT authentication provider `DaoAuthenticationProvider` ที่ประกาศเองจึงถอดออกได้ ของที่ Spring สร้างจาก `UserDetailsService` ยังอยู่ เพื่อให้ `POST /auth/login` ตรวจรหัสได้

ใน `authorizeHttpRequests` เพิ่ม `permitAll` ให้ `POST /auth/login` คู่กับ register ที่มีอยู่แล้ว ที่เหลือยัง `authenticated` คนที่ยังไม่ login ต้องเข้า endpoint นี้ได้

`sessionCreationPolicy` เปลี่ยนจาก `IF_REQUIRED` เป็น `STATELESS`

แล้วต่อ `oauth2ResourceServer` แบบ JWT และส่ง `JwtAuthenticationConverter` เข้าไป เพื่อไม่ให้ framework แต่ง role เอง

`SecurityConfig` ทั้งไฟล์หลังรวม bean ข้างต้น

```java
package com.spring.springpracticaldemo.config;

import java.util.Base64;
import javax.crypto.SecretKey;
import javax.crypto.spec.SecretKeySpec;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.configuration.AuthenticationConfiguration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.oauth2.jwt.JwtDecoder;
import org.springframework.security.oauth2.jwt.JwtEncoder;
import org.springframework.security.oauth2.jwt.NimbusJwtDecoder;
import org.springframework.security.oauth2.jwt.NimbusJwtEncoder;
import org.springframework.security.oauth2.server.resource.authentication.JwtAuthenticationConverter;
import org.springframework.security.oauth2.server.resource.authentication.JwtGrantedAuthoritiesConverter;
import org.springframework.security.web.SecurityFilterChain;

@Configuration
public class SecurityConfig {

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }

    @Bean
    public SecretKey jwtSecretKey(@Value("${jwt.secret}") String secret) {
        byte[] decodedKey = Base64.getDecoder().decode(secret);
        return new SecretKeySpec(decodedKey, "HmacSHA256");
    }

    @Bean
    public AuthenticationManager authenticationManager(AuthenticationConfiguration configuration) throws Exception {
        return configuration.getAuthenticationManager();
    }

    @Bean
    public JwtAuthenticationConverter jwtAuthenticationConverter() {
        JwtGrantedAuthoritiesConverter grantedAuthoritiesConverter = new JwtGrantedAuthoritiesConverter();
        grantedAuthoritiesConverter.setAuthoritiesClaimName("authorities");
        grantedAuthoritiesConverter.setAuthorityPrefix("");
        JwtAuthenticationConverter converter = new JwtAuthenticationConverter();
        converter.setJwtGrantedAuthoritiesConverter(grantedAuthoritiesConverter);
        return converter;
    }

    @Bean
    public JwtEncoder jwtEncoder(SecretKey jwtSecretKey) {
        return NimbusJwtEncoder.withSecretKey(jwtSecretKey).build();
    }

    @Bean
    public JwtDecoder jwtDecoder(SecretKey jwtSecretKey) {
        return NimbusJwtDecoder.withSecretKey(jwtSecretKey).build();
    }

    @Bean
    public SecurityFilterChain securityFilterChain(
            HttpSecurity http, JwtAuthenticationConverter jwtAuthenticationConverter) throws Exception {
        http.csrf(csrf -> csrf.disable())
                .authorizeHttpRequests(auth -> auth
                        .requestMatchers("/api/user/register", "/auth/login").permitAll()
                        .anyRequest().authenticated())
                .sessionManagement(session -> session.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
                .oauth2ResourceServer(oauth2 -> oauth2
                        .jwt(jwt -> jwt.jwtAuthenticationConverter(jwtAuthenticationConverter)));
        return http.build();
    }
}
```

> 💡 **เพิ่มเติม:** วิดีโอฉีด dependency ด้วย `@Autowired` ที่ฟิลด์ แล้วบอกเองว่าไม่ควรทำแบบนั้น โค้ดในโน้ตใช้ constructor

ลำดับของ request ที่ต้องมี Bearer token เป็นแบบนี้

```mermaid
sequenceDiagram
    participant C as Client
    participant F as BearerTokenAuthenticationFilter
    participant D as JwtDecoder
    participant H as GET /api/user/hello
    C->>F: Authorization: Bearer JWT
    F->>D: ตรวจลายเซ็นและ exp
    D-->>F: claims
    Note over F: สร้าง Authentication<br/>จาก sub และ authorities
    F->>H: เรียก controller
    H-->>C: Hello, you're logged in as Aditya
```

---

## 14. Login endpoint กับ JwtService

[▶ 1:08:27](https://youtu.be/pUWp1EuVnGs?t=4107)

เขาไม่ใส่ login ไว้ใน `UserController` เพราะไม่อยากให้ controller นั้นรู้จักคลาสของ security จึงสร้าง `AuthController` ที่ `@RequestMapping("/auth")` และ `@PostMapping("/login")`

DTO ของ register มี username กับ password อยู่แล้ว แต่ชื่อมันบอกว่าเป็นของ register เขาจึงแยก `LoginRequestDTO` (username, password) กับ `LoginResponseDTO` ฟิลด์เดียวชื่อ `accessToken` วิดีโอใช้ Lombok `@AllArgsConstructor` กับ getter/setter

> 💡 **เพิ่มเติม:** เดโมใช้ record ของ Java 21 แทน Lombok ได้ JSON ชื่อฟิลด์เดียวกัน

```java
public record LoginRequestDTO(String username, String password) {
}

public record LoginResponseDTO(String accessToken) {
}
```

[▶ 1:10:40](https://youtu.be/pUWp1EuVnGs?t=4240)

เพราะ login เป็น endpoint ของเราเอง Spring จะไม่สร้าง Authentication ให้ ต้องทำสองก้าวที่ filter เคยทำให้

1. สร้าง authentication request ด้วย `UsernamePasswordAuthenticationToken.unauthenticated(username, password)` ก้อนนี้ยังไม่ผ่านการตรวจ
2. ส่งเข้า `authenticationManager.authenticate` manager เป็น interface มันเลือก provider ที่ `supports` token ชนิดนี้ ที่นี่คือ `DaoAuthenticationProvider` ผลที่ได้คือ authentication response ที่ตรวจผ่านแล้ว

จากนั้นส่งก้อนนี้ให้ `JwtService.generateToken` แล้วห่อสตริงที่ได้ด้วย `LoginResponseDTO`

```java
package com.spring.springpracticaldemo.controller;

import com.spring.springpracticaldemo.dto.LoginRequestDTO;
import com.spring.springpracticaldemo.dto.LoginResponseDTO;
import com.spring.springpracticaldemo.service.JwtService;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/auth")
public class AuthController {

    private final AuthenticationManager authenticationManager;
    private final JwtService jwtService;

    public AuthController(AuthenticationManager authenticationManager, JwtService jwtService) {
        this.authenticationManager = authenticationManager;
        this.jwtService = jwtService;
    }

    @PostMapping("/login")
    public LoginResponseDTO login(@RequestBody LoginRequestDTO request) {
        Authentication authenticationRequest = UsernamePasswordAuthenticationToken.unauthenticated(
                request.username(), request.password());
        Authentication authentication = authenticationManager.authenticate(authenticationRequest);
        String token = jwtService.generateToken(authentication);
        return new LoginResponseDTO(token);
    }
}
```

[▶ 1:12:56](https://youtu.be/pUWp1EuVnGs?t=4376)

`JwtService` เป็น `@Service` method `generateToken` รับ `Authentication` แล้วทำตามลำดับนี้

1. `Instant.now()` เป็นเวลาออก token
2. แปลง `getAuthorities()` เป็น `List<String>` เช่น `ROLE_USER`
3. สร้าง claims ด้วย builder ของ `JwtClaimsSet`
   - `issuer` จาก `jwt.issuer`
   - `issuedAt` คือเวลาปัจจุบัน คีย์ `iat`
   - `expiresAt` คือตอนนี้บวก `jwt.expiry` วินาที คีย์ `exp`
   - `subject` จาก `authentication.getName()` คีย์ `sub`
   - claim ชื่อ `authorities` เป็นลิสต์ด้านบน
4. ไม่ใส่คีย์อื่น
5. ให้ `JwtEncoder.encode` เซ็น header กับ payload แล้วคืน `getTokenValue()`

header `alg` ถูกกำหนดตอนสร้าง encoder และ `typ` เป็น `JWT` เขาจึงไม่ได้ประกอบ Base64 เองในเมธอดนี้ ออบเจกต์ `Jwt` ที่ encode คืนมามี header, claims และ token value ส่วนที่ส่งให้ client คือ token value

```java
package com.spring.springpracticaldemo.service;

import java.time.Instant;
import java.util.List;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.oauth2.jose.jws.MacAlgorithm;
import org.springframework.security.oauth2.jwt.JwsHeader;
import org.springframework.security.oauth2.jwt.JwtClaimsSet;
import org.springframework.security.oauth2.jwt.JwtEncoder;
import org.springframework.security.oauth2.jwt.JwtEncoderParameters;
import org.springframework.stereotype.Service;

@Service
public class JwtService {

    private final String issuer;
    private final long expiry;
    private final JwtEncoder jwtEncoder;

    public JwtService(
            @Value("${jwt.issuer}") String issuer,
            @Value("${jwt.expiry}") long expiry,
            JwtEncoder jwtEncoder) {
        this.issuer = issuer;
        this.expiry = expiry;
        this.jwtEncoder = jwtEncoder;
    }

    public String generateToken(Authentication authentication) {
        Instant now = Instant.now();
        List<String> authorities = authentication.getAuthorities().stream()
                .map(GrantedAuthority::getAuthority)
                .toList();

        JwtClaimsSet claims = JwtClaimsSet.builder()
                .issuer(issuer)
                .issuedAt(now)
                .expiresAt(now.plusSeconds(expiry))
                .subject(authentication.getName())
                .claim("authorities", authorities)
                .build();

        JwsHeader header = JwsHeader.with(MacAlgorithm.HS256).type("JWT").build();
        return jwtEncoder.encode(JwtEncoderParameters.from(header, claims)).getTokenValue();
    }
}
```

> ⚠️ **Boot 4 vs 3:** หลัง login ด้วยรหัสผ่าน Spring Security 7 ใส่ `FactorGrantedAuthority` เพิ่มใน `Authentication` ค่าของมันคือ `FACTOR_PASSWORD` โค้ดที่ก็อป authorities ทุกตัวลง claim จึงได้ทั้ง `ROLE_USER` และ `FACTOR_PASSWORD` บน Spring Security 6 ซึ่งมากับ Boot 3 ลิสต์นี้มีแต่ role
> ([Multi-Factor Authentication in Spring Security 7](https://spring.io/blog/2025/10/21/multi-factor-authentication-in-spring-security-7))

`AuthController` เรียก `generateToken` แล้ว `return new LoginResponseDTO(token)` ไม่คืน token เป็นสตริงเปล่า

---

## 15. ยิง Bearer token

[▶ 1:18:35](https://youtu.be/pUWp1EuVnGs?t=4715)

ลำดับใน Postman ของวิดีโอ

1. `POST /auth/login` เลือก No Auth ใส่ body เป็น username `Aditya` และ password `secret123` endpoint นี้ถูก `permitAll` จึงเข้าได้โดยยังไม่มี token
2. ได้ JSON ที่มี `accessToken` แล้วก็อปไว้
3. `GET /api/user/hello` เลือก Bearer Token แล้ววางค่าที่ก็อป หรือใส่ header เองเป็น `Authorization: Bearer <token>`
4. ได้ `Hello, you're logged in as Aditya` โดยไม่กรอกรหัสอีก
5. เรียก hello ซ้ำได้เรื่อย ๆ จนกว่า token จะหมดอายุ

> 💡 **เพิ่มเติม:** ผลด้านล่างมาจาก Boot 4.1.1 พอร์ต 9038 หลัง `mvn -q -DskipTests package` ผ่าน ต้องสมัครผู้ใช้ก่อนเพราะฐาน H2 ว่างทุกครั้งที่เปิดแอป

```bash
curl -s -X POST http://localhost:9038/api/user/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"Aditya","password":"secret123"}'
```

```json
{"username":"Aditya","message":"Aditya user saved successfully"}
```

```bash
curl -s -D - -X POST http://localhost:9038/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"Aditya","password":"secret123"}'
```

```http
HTTP/1.1 200
Content-Type: application/json
Content-Length: 328
```

```json
{"accessToken":"eyJraWQiOiJ6SVVPdHJ0ckk5VkhsVENYN3hoa3dzVWRGeTRzYzE0TG9CLUZNcTBmY1BVIiwidHlwIjoiSldUIiwiYWxnIjoiSFMyNTYifQ.eyJpc3MiOiJDb2RlciBBcm15IEFQSSIsInN1YiI6IkFkaXR5YSIsImV4cCI6MTc5MjE1NjI0NywiaWF0IjoxNzkyMTU1MzQ3LCJhdXRob3JpdGllcyI6WyJST0xFX1VTRVIiLCJGQUNUT1JfUEFTU1dPUkQiXX0.N2MBgY2UdwQyjINJRn2MFDsShJS29sXEgo0aTjmr08U"}
```

response ของ login ไม่มี `Set-Cookie` และไม่มี `JSESSIONID` สอดคล้องกับ `STATELESS`

ถอดสองส่วนแรกของ token นี้ได้ header กับ payload จริง `exp - iat` เท่ากับ 900

```json
{"kid":"zIUOtrtrI9VHlTCX7xhkwsUdFy4sc14LoB-FMq0fcPU","typ":"JWT","alg":"HS256"}
```

```json
{"iss":"Coder Army API","sub":"Aditya","exp":1792156247,"iat":1792155347,"authorities":["ROLE_USER","FACTOR_PASSWORD"]}
```

```bash
curl -s -D - http://localhost:9038/api/user/hello \
  -H "Authorization: Bearer <accessToken>"
```

```http
HTTP/1.1 200
Content-Type: text/plain;charset=UTF-8
Content-Length: 33

Hello, you're logged in as Aditya
```

ไม่ใส่ header ได้ 401 และ Basic Auth เดิมก็ใช้ไม่ได้อีก เพราะเอา `httpBasic` ออกแล้ว

```http
HTTP/1.1 401
WWW-Authenticate: Bearer resource_metadata="http://localhost:9038/.well-known/oauth-protected-resource"
Content-Length: 0
```

> ⚠️ **Boot 4 vs 3:** challenge นี้มีพารามิเตอร์ `resource_metadata` ตาม RFC 9728 `OAuth2ProtectedResourceMetadataFilter` อยู่ใน filter chain ของ resource server ตั้งแต่ Spring Security 7 Boot 3 ที่ใช้ Security 6 ตอบ `WWW-Authenticate: Bearer` โดยไม่มีพารามิเตอร์นี้เป็นค่า default
> ([OAuth2ProtectedResourceMetadataFilter](https://docs.spring.io/spring-security/reference/api/java/org/springframework/security/oauth2/server/resource/web/OAuth2ProtectedResourceMetadataFilter.html))

> 💡 **เพิ่มเติม:** แก้ claim `authorities` เป็น `ROLE_ADMIN` แล้วประกอบ token เดิมโดยไม่เซ็นใหม่ ได้ 401 พร้อม `error="invalid_token"` และ `Invalid signature` นี่คือการทดลองของขั้นตอนลายเซ็นบนกระดาน วิดีโอไม่ได้ยิง token ที่ถูกแก้ใน Postman

ผู้สอนปิดตอนว่าโค้ดส่วนใหญ่เป็น boilerplate ที่ซ้ำทุกแอป ของที่ต้องจำคือมี encoder สำหรับเซ็น, decoder สำหรับตรวจ, secret key, authentication converter, `SecurityFilterChain` ที่ปรับให้เป็น stateless, `AuthenticationManager` และ `PasswordEncoder` เมื่อชุดนี้ครบ การยืนยันตัวที่เหลือของแอปอยู่บนสถาปัตยกรรม stateless

---

## 16. สรุปท้ายตอน

### ประเด็นที่ต้องจำ

1. HTTP ไม่จำ request ก่อนหน้า Basic Auth จึงต้องส่งรหัสทุกครั้ง
2. Spring Security สร้าง session ให้เป็นค่า default (`IF_REQUIRED`) browser ส่ง cookie กลับมาเอง Postman ไม่ส่งให้ถ้าไม่ได้เปิดไว้
3. session id เป็นสตริงสุ่ม ของที่แปลได้คือ Authentication ที่ server เก็บไว้คู่กับ id นั้น
4. หลาย instance ทำให้ session ในเครื่องเดียวไม่พอ ต้องใช้ sticky session หรือที่เก็บส่วนกลาง
5. Stateless ไม่เก็บ Authentication สร้างใหม่จาก token ทุก request จึงไปเครื่องไหนก็ได้
6. Base64URL อ่านย้อนกลับได้โดยไม่ใช้กุญแจ hash ใครก็คำนวณใหม่ได้ ทั้งคู่กันการแก้ role ไม่ได้
7. JWT ที่ใช้ในตอนนี้เป็นลายเซ็น ไม่ใช่การเข้ารหัสลับ payload ยังอ่านได้ ห้ามใส่ความลับลงไป
8. Token มีสามส่วน `header.payload.signature` เดโมเซ็นด้วย HS256 และ secret key ดอกเดียว
9. Asymmetric ให้คนออก token ถือ private key และให้ service ที่เหลือตรวจด้วย public key
10. `JwtEncoder` เซ็นตอน login `JwtDecoder` ตรวจตอน request ถัดไปผ่าน `oauth2ResourceServer`
11. Login ของเราเองต้องสร้าง `UsernamePasswordAuthenticationToken` แล้วเรียก `AuthenticationManager` Spring ไม่ทำให้เมื่อเอา form login ออก
12. `SessionCreationPolicy.STATELESS` ทำให้ response ของ login ไม่ตั้ง `JSESSIONID`
13. อายุ token ในเดโมคือ 900 วินาที หมดแล้วต้อง login ใหม่

### คำถามทบทวน

1. ทำไม Postman ที่ยิง hello สำเร็จแล้วยิงรอบถัดไปแบบ No Auth ถึงได้ 401 ทั้งที่ Spring สร้าง session ไว้แล้ว
2. `IF_REQUIRED`, `NEVER` และ `STATELESS` ต่างกันอย่างไรเรื่องการสร้างและการใช้ session
3. ถ้ามีแอปสอง instance แล้วเก็บ session ไว้ใน memory ของเครื่องแรก request ถัดไปที่ไปเครื่องที่สองจะเกิดอะไร
4. ทำไมส่งเฉพาะ Base64URL ของ JSON แล้วให้ server เชื่อถือไม่ได้
5. client ที่รู้ว่า server ใช้ SHA-256 ปลอมส่วน hash ของ token ได้อย่างไร และลายเซ็นด้วย secret key ปิดช่องนี้ตรงไหน
6. JWE กับ JWT แบบมี signature ต่างกันตรงที่ client อ่าน payload ได้หรือไม่
7. ทำไม service ที่เหลือควรได้แต่ public key ถ้าใช้กุญแจแบบ asymmetric
8. claim ไหนในเดโมกลายเป็น principal ของ `Authentication` และ claim ไหนกลายเป็น authorities
9. ถ้าไม่เรียก `setAuthoritiesClaimName("authorities")` แล้ว `setAuthorityPrefix("")` สิทธิ์ใน token จะถูกอ่านเป็น role หรือไม่
10. ทำไม `GET /api/user/hello` หลังเปลี่ยนมาใช้ JWT ถึงไม่รับ Basic Auth อีก

### ตอนถัดไป

[#39 — OAuth 2.0 Masterclass | Google Login with Spring Boot, OIDC & PKCE](./39-oauth2-google-login.md) ผู้สอนบอกว่า dependency ของ resource server ที่เพิ่งใส่ จะถูกใช้ต่อเมื่อทำ OAuth ในวิดีโอหน้า
