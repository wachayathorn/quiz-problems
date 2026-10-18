# Spring Boot Course

Glossary กลางของโน้ตคอร์ส Spring Boot ทุกไฟล์ใช้ศัพท์และความหมายตามนี้ ศัพท์เทคนิคใช้ภาษาอังกฤษ ไม่แปลเป็นไทย

## Language

### Web foundations

**CRUD**:
งาน create, read, update และ delete ที่ backend ทำกับข้อมูล
_Avoid_: cred

**Controller**:
ชั้นที่รับ HTTP request แล้วส่ง HTTP response
_Avoid_: handler, API class

**Service**:
ชั้นที่ทำ business logic แล้วส่งต่องานเก็บข้อมูลให้ Repository
_Avoid_: manager

**Path Variable**:
ค่าที่ฝังอยู่ใน path เช่น `/students/1` อ่านด้วย `@PathVariable`
_Avoid_: path param

**Request Parameter**:
คู่ชื่อกับค่าใน query string อ่านด้วย `@RequestParam`
_Avoid_: param, search param

**Endpoint**:
คู่ของ HTTP method กับ path ที่ server เปิดรับ เช่น `GET /courses`
_Avoid_: route, URL (เมื่อหมายถึงทั้งคู่ method+path)

**Servlet**:
Java class ที่รับ `HttpServletRequest` และเขียน `HttpServletResponse` โดยรันอยู่ใน Servlet Container
_Avoid_: servlet class, handler

**Servlet Container**:
ซอฟต์แวร์ที่เปิด port, parse HTTP, จัดการ thread และเรียก Servlet ให้ เช่น Tomcat, Jetty
_Avoid_: web server, app server (ใช้ได้ในบทสนทนา แต่ในโน้ตให้เรียก Servlet Container)

**Embedded Server**:
Servlet Container ที่ถูกบรรจุมาใน `.jar` ของแอป และเริ่มทำงานพร้อมแอป
_Avoid_: built-in Tomcat

**WAR**:
แฟ้มสำหรับ deploy เว็บแอปขึ้น Servlet Container ที่ติดตั้งแยก มี class และ configuration ของแอป
_Avoid_: JAR (เมื่อหมายถึงแฟ้มเว็บนี้)

**Context path**:
Path ชั้นแรกที่ Servlet Container ใช้เลือกแอป เช่น `/crud-app`
_Avoid_: base URL

**Filter**:
Interface `jakarta.servlet.Filter` ที่ Servlet Container เรียกก่อนและหลัง Servlet งานอยู่ที่ `doFilter`
_Avoid_: middleware, interceptor

**`@Order`**:
ตัวเลขที่ใช้เรียง Filter ตัวเลขน้อยกว่าได้ทำก่อน
_Avoid_: priority

**Filter Chain**:
ลำดับของ Filter ที่ส่งงานต่อด้วย `chain.doFilter` โค้ดหลังการเรียกนั้นคือขากลับ
_Avoid_: pipeline

**OncePerRequestFilter**:
Base class ของ Spring ที่ให้ `doFilterInternal` รับชนิด HTTP และไม่รันซ้ำใน request เดียวกัน
_Avoid_: Spring Filter

**FilterRegistrationBean**:
Bean ของ Spring Boot สำหรับลงทะเบียน Filter พร้อมลำดับ ชื่อ และ URL pattern
_Avoid_: filter config

**DispatcherServlet**:
Servlet ตัวกลางของ Spring MVC ที่รับ request แล้วส่งต่อให้ controller
_Avoid_: front controller, main servlet

**HandlerInterceptor**:
Interface ของ Spring MVC ที่ห่อ handler ด้วย `preHandle`, `postHandle` และ `afterCompletion`
_Avoid_: filter, middleware

**Handler**:
เป้าหมายที่ DispatcherServlet เลือกไว้แล้ว อาจเป็น Controller method หรือ static resource
_Avoid_: controller (เมื่อหมายถึง handler ทุกชนิด)

**HandlerMethod**:
Object ที่เก็บ controller bean กับ method ที่จะถูกเรียก
_Avoid_: handler (คำนี้กว้างกว่า)

**InterceptorRegistry**:
ที่ลงทะเบียน interceptor ผ่าน `WebMvcConfigurer.addInterceptors`
_Avoid_: FilterRegistrationBean

**HandlerMapping**:
ส่วนที่จับคู่ endpoint กับ method ของ Controller
_Avoid_: router

**HandlerAdapter**:
ส่วนที่เรียก method ของ Controller และใส่ค่า argument
_Avoid_: invoker

**ViewResolver**:
ส่วนที่แปลงชื่อ view เป็น View จริง เช่นไฟล์ JSP
_Avoid_: template engine

**Model**:
Object ที่ถือข้อมูลจาก Controller ไปให้ View
_Avoid_: DTO

### Spring Core

**Dependency Injection**:
การส่ง object ที่ class หนึ่งต้องใช้เข้ามาจากข้างนอก แทนที่ class นั้นจะสร้างเอง
_Avoid_: wiring (ใช้เมื่อพูดถึงการเชื่อม Bean แต่ไม่ใช่ชื่อของแนวคิด)

**Inversion of Control**:
การให้ framework เป็นคนเรียกโค้ดของเราและสร้าง dependency แทนที่เราจะควบคุมเอง
_Avoid_: DI (DI เป็นเทคนิคหนึ่งของ IoC ไม่ใช่ชื่อเดียวกัน)

**IoC Container**:
ส่วนของ Spring ที่สร้าง Bean, inject dependency และดูแล lifecycle ชื่อใน API คือ `ApplicationContext`
_Avoid_: Spring container, BeanFactory (BeanFactory เป็น interface ระดับล่างกว่า)

**Bean**:
Object ที่อยู่ใน IoC Container ทุก Bean เป็น object แต่ไม่ใช่ทุก object ที่เป็น Bean
_Avoid_: component (component คือ class ที่ถูกประกาศให้เป็น Bean)

**Bean Definition**:
Metadata ของ Bean ที่ container เก็บไว้ก่อนสร้าง object จริง เช่น class, scope และ dependency
_Avoid_: bean config

**Component Scan**:
การที่ container หา class ที่มี `@Component` หรือ stereotype annotation ใน package ที่กำหนด แล้วลงทะเบียนเป็น Bean
_Avoid_: classpath scan

**Bean Scope**:
กฎว่า Bean หนึ่งตัวมี object กี่ชิ้น และอยู่ได้นานแค่ไหน ค่า default คือ singleton
_Avoid_: lifecycle (lifecycle คือลำดับขั้นของ Bean ไม่ใช่จำนวน instance)

**Circular Dependency**:
Bean สองตัวขึ้นไปที่ต้องการกันและกัน จน constructor injection สร้างไม่ได้
_Avoid_: cycle

### Maven

**POM**:
ไฟล์ `pom.xml` ที่บรรยายโปรเจกต์ Maven ได้แก่พิกัด GAV, dependency และวิธี build
_Avoid_: build file

**GAV**:
พิกัดของ artifact ใน Maven คือ `groupId`, `artifactId` และ `version`
_Avoid_: coordinates

**Transitive Dependency**:
Dependency ของ dependency ที่ Maven ดึงมาให้โดยไม่ต้องประกาศเอง
_Avoid_: indirect dependency

**Effective POM**:
POM ที่ Maven ใช้จริงหลังรวม Super POM, parent ทุกชั้น, POM ของโปรเจกต์ และ profile ที่เปิดอยู่
_Avoid_: parent POM

**provided**:
Maven scope ที่ใช้ตอน compile แต่ไม่แพ็กเข้า artifact เพราะปลายทางมีไลบรารีให้อยู่แล้ว
_Avoid_: optional

### Spring ecosystem

**Spring Core**:
ชั้นฐานของ Spring ที่มี IoC container, Dependency Injection และ Bean ทุก module ใช้ร่วมกัน
_Avoid_: Spring Framework (เมื่อหมายถึงแค่ core)

**Spring MVC**:
Module สำหรับเว็บแบบ servlet ที่ทำงานบน Servlet API
_Avoid_: Spring Web (กำกวมกับ WebFlux)

**Cross-cutting concern**:
งานที่ตัดข้ามหลาย class เช่น log, security และการวัดเวลา
_Avoid_: shared logic, common code

**Scattering**:
Concern เดียวกันถูกเขียนซ้ำในหลาย class
_Avoid_: duplication

**Tangling**:
Method เดียวปนหลาย concern จน business logic มองไม่ชัด
_Avoid_: mixed code

**Aspect-Oriented Programming (AOP)**:
Module ของ Spring ที่แยก cross-cutting concern ออกจาก business logic โดยให้ framework หุ้ม Bean
_Avoid_: OOP

**Decorator**:
Object ที่หน้าตาเหมือน business object แล้วสอด logic ก่อนและหลังเรียกของจริง
_Avoid_: wrapper class

**Marker annotation**:
Annotation ที่ไม่มี member ใช้เป็นเครื่องหมายให้ aspect จับ method นั้น
_Avoid_: empty annotation

**Meta-annotation**:
Annotation ที่แปะบน annotation อื่น เช่น `@Target` และ `@Retention`
_Avoid_: annotation of annotation

**Annotation binding**:
รูปแบบ `@annotation(ชื่อparameter)` ที่ส่งตัว annotation เข้า advice เมื่อชื่อตรงกัน
_Avoid_: type matching (เมื่อหมายถึงแบบนี้)

**BeanPostProcessor**:
Interface ของ Spring ที่รับ Bean ทุกตัวก่อนและหลัง initialization จุดที่ AOP คืน proxy อยู่หลัง initialization
_Avoid_: post processor, Java interface

**Aspect**:
Class ที่รวม advice กับ pointcut เพื่อแทรก cross-cutting concern
_Avoid_: interceptor

**Advice**:
งานที่ aspect จะทำ และจังหวะที่จะทำ เช่นก่อนหรือหลัง method
_Avoid_: annotation (เมื่อหมายถึงทั้งก้อนของ advice)

**Pointcut designator**:
คำนำของ pointcut เช่น `execution`, `within` และ `bean`
_Avoid_: annotation (เมื่อหมายถึงคำนำนี้)

**Named pointcut**:
นิพจน์ที่เก็บใน method ว่างด้วย `@Pointcut` แล้วให้ advice อ้างชื่อ
_Avoid_: pointcut method

**CGLIB proxy**:
Proxy ที่เป็น subclass ของคลาสเป้าหมาย เป็นค่า default ของ Spring Boot
_Avoid_: CGI proxy

**JDK dynamic proxy**:
Proxy ที่ implement interface เดียวกับเป้าหมาย
_Avoid_: Java proxy

**Pointcut**:
นิพจน์ที่เลือก method ที่จะถูก intercept
_Avoid_: path, URL

**JoinPoint**:
จุดที่ advice ทำงาน และเป็นทางอ่านข้อมูลของ method เป้าหมาย เช่น argument
_Avoid_: method

**ProceedingJoinPoint**:
JoinPoint ของ `@Around` ที่เรียก method เป้าหมายผ่าน `proceed()` ได้
_Avoid_: JoinPoint (เมื่อหมายถึงตัวที่เรียก proceed ได้)

**Proxy**:
Object ที่ Spring สร้างตอน runtime เพื่อหุ้ม Bean แล้วสอด AOP
_Avoid_: decorator

**Authorization**:
การกำหนดว่า identity ที่ผ่านการตรวจแล้วทำอะไรได้
_Avoid_: authentication, login

**Principal**:
Identity ที่กำลังถูกใช้ใน `Authentication`
_Avoid_: user (เมื่อหมายถึง principal ทุกชนิด)

**GrantedAuthority**:
สิทธิ์หนึ่งรายการใน `Authentication` role คือ authority ที่ขึ้นต้นด้วย `ROLE_`
_Avoid_: permission (เมื่อหมายถึงชนิดใน API)

**SecurityContext**:
Container ที่เก็บ `Authentication` ปัจจุบัน
_Avoid_: ApplicationContext

**SecurityContextHolder**:
ทางอ่าน SecurityContext จากเธรดปัจจุบัน
_Avoid_: session

**Authentication**:
การตรวจว่าคนที่อ้างตัวตนเป็นเจ้าของหลักฐานนั้นจริง
_Avoid_: login object

**PasswordEncoder**:
ตัวที่ `encode` ตอนเก็บรหัส และ `matches` ตอนเทียบ
_Avoid_: encrypt

**BCrypt**:
Hash ที่ทำงานช้าและฝัง salt ไว้ในสตริงผลลัพธ์
_Avoid_: SHA (เมื่อหมายถึงการเก็บรหัสผ่าน)

**Salt**:
ค่าสุ่มต่อรหัสที่ต่อก่อน hash เพื่อไม่ให้ hash ไปตรงตารางที่คำนวณล่วงหน้า
_Avoid_: pepper

**UserDetails**:
สัญญาที่ Spring Security ใช้ดึง username, password, authorities และสถานะบัญชี
_Avoid_: User (เมื่อหมายถึงตัวนี้)

**UserDetailsService**:
Interface ที่มี `loadUserByUsername` สำหรับหา `UserDetails` จาก username
_Avoid_: user service

**DaoAuthenticationProvider**:
`AuthenticationProvider` ที่โหลด `UserDetails` แล้วเทียบรหัสด้วย `PasswordEncoder.matches`
_Avoid_: DaoAuthenticationFilter

**SecurityFilterChain**:
เส้น filter หนึ่งชุดพร้อม matcher ว่า path ไหนเข้าเส้นนี้
_Avoid_: filter list

**DelegatingFilterProxy**:
Filter ตัวเดียวที่ Servlet Container รู้จัก แล้วส่งต่อให้ `FilterChainProxy`
_Avoid_: security filter

**AuthenticationManager**:
Interface ที่รับ `Authentication` แล้วคืน `Authentication` ของที่ใช้จริงคือ `ProviderManager`
_Avoid_: authenticator

**OAuth 2.0**:
โปรโตคอลให้ client ขอสิทธิ์เข้าถึง resource ของผู้ใช้โดยไม่รับรหัสผ่านของเจ้าของ resource
_Avoid_: login (เมื่อหมายถึงทั้งโปรโตคอล)

**Authorization Server**:
ระบบที่ให้ผู้ใช้ล็อกอิน กด consent แล้วออก token
_Avoid_: authentication server

**OpenID Connect**:
ชั้นยืนยันตัวตนที่วางบน OAuth 2.0 ตอบว่าใครล็อกอิน
_Avoid_: OAuth (เมื่อหมายถึงการยืนยันตัวตน)

**JWT**:
Token สามส่วน `header.payload.signature` ที่เซ็นแล้วให้ server สร้าง Authentication ใหม่ได้โดยไม่เก็บ session
_Avoid_: encryption, session id

**Access token**:
JWT ที่ login คืนให้ client ส่งกลับใน `Authorization: Bearer`
_Avoid_: session id, API key

**Claim**:
คู่ชื่อกับค่าใน payload ของ JWT เช่น `sub`, `iss`, `iat` และ `exp`
_Avoid_: field, attribute

**Stateful**:
Server เก็บ Authentication คู่กับ session id
_Avoid_: session login

**Stateless**:
ไม่เก็บ Authentication แล้วสร้างใหม่จาก token ทุก request
_Avoid_: ไม่มี login

**SessionCreationPolicy**:
นโยบายว่าจะสร้างหรือใช้ HTTP session หรือไม่ ค่า default คือ `IF_REQUIRED`
_Avoid_: cookie policy

**ID Token**:
JWT ที่ client ใช้รู้จักผู้ใช้
_Avoid_: bearer token (เมื่อหมายถึง token สำหรับเรียก API)

**Authorization Code**:
รหัสครั้งเดียวที่ browser ได้รับ แล้วเซิร์ฟเวอร์เอาไปแลก token ทางหลังบ้าน
_Avoid_: token

**PKCE**:
คู่ code verifier กับ code challenge ที่กันคนดัก authorization code ไปแลก token
_Avoid_: Bean Scope

**Unit Test**:
เทส method หรือ class เดียวโดยไม่เปิดส่วนอื่นของแอป
_Avoid_: integration (เมื่อยัง mock dependency อยู่)

**Integration Test**:
เทสหลาย component พร้อมกัน เช่น Controller ถึง Repository หรือ Repository กับฐานข้อมูล
_Avoid_: end-to-end

**Slice Test**:
การโหลด application context แค่ชั้นเดียวผ่าน `@WebMvcTest` หรือ `@DataJpaTest`
_Avoid_: layer test

**Test Double**:
ตัวแทนของ dependency ตอนเทส มี dummy, stub, fake และ mock
_Avoid_: mock (เมื่อหมายถึงทั้งสี่แบบ)

**Arrange-Act-Assert**:
โครงของ test case คือจัด input, เรียก method แล้วเทียบผล
_Avoid_: setup-run-check

**MockMvc**:
ตัวยิง HTTP เข้า Spring MVC โดยไม่เปิด Servlet Container
_Avoid_: embedded Tomcat

**Spring Boot**:
Automation layer แบบ opinionated ที่วางอยู่บน module ต่าง ๆ ของ Spring ทำหน้าที่ตั้งค่าให้อัตโนมัติ และ override ได้
_Avoid_: Spring (เมื่อหมายถึง Boot โดยเฉพาะ)

**Auto-configuration**:
การที่ Spring Boot สร้าง Bean จาก class ที่ติด `@AutoConfiguration` เมื่อเงื่อนไขบน classpath ผ่าน โดยผู้ใช้ไม่ได้ประกาศ Bean นั้นเอง
_Avoid_: automatic config

**Profile**:
ชื่อ environment ที่ Spring ใช้เลือกไฟล์ config และ Bean เช่น `dev` หรือ `prod`
_Avoid_: profiling

**`spring.profiles.active`**:
Property ที่บอกว่า profile ไหนกำลังทำงาน
_Avoid_: active env

**Profile-specific properties**:
ไฟล์ `application-{profile}.properties` หรือ `.yml` ที่ทับ key เดียวกันในไฟล์หลักเมื่อ profile นั้นทำงาน
_Avoid_: env file

**YAML**:
ไฟล์ config แบบต้นไม้ที่ Spring ทำให้แบนเป็น key แบบ `application.properties`
_Avoid_: yml (เมื่อหมายถึงฟอร์แมตทั้งกลุ่ม)

**`@Profile`**:
Annotation ที่ให้สร้าง Bean เฉพาะเมื่อ profile ที่ระบุทำงาน
_Avoid_: `@ActiveProfiles`

**application.properties**:
ไฟล์ `key=value` ที่ Spring Boot อ่านตอน startup เพื่อส่งค่าจากข้างนอกเข้าแอป
_Avoid_: config file, `@Configuration`

**Relaxed binding**:
การจับ `retry-count`, `retryCount`, `retry_count` และ environment variable `RETRY_COUNT` เข้า field เดียวกัน
_Avoid_: canonical form

**ApplicationRunner**:
Interface ที่ Spring เรียก `run(ApplicationArguments)` หลัง application context พร้อม
_Avoid_: main

**CommandLineRunner**:
Interface จังหวะเดียวกับ ApplicationRunner แต่รับ `run(String... args)`
_Avoid_: main

**`@SpringBootApplication`**:
Annotation บน class หลักที่รวม `@SpringBootConfiguration`, `@EnableAutoConfiguration` และ `@ComponentScan`
_Avoid_: `@SpringBoot` (ไม่มี annotation ชื่อนี้)

### Data access

**Object-Relational Impedance Mismatch**:
เรื่องเดียวกันต้องเขียนเป็น object ใน Java แต่เป็นตาราง แถว และคอลัมน์ใน SQL
_Avoid_: object-relational model

**Embeddable**:
Object ที่ฝังคอลัมน์เข้าตารางของเจ้าของ ด้วย `@Embeddable` และ `@Embedded`
_Avoid_: inner entity

**ElementCollection**:
Collection ที่ Hibernate แยกไปอีกตารางพร้อม foreign key กลับมาที่เจ้าของ โดยที่สมาชิกยังไม่ใช่ Entity
_Avoid_: one-to-many

**AttributeConverter**:
ตัวแปลงค่าคนละชนิดระหว่างฟิลด์ Java กับคอลัมน์ เช่น boolean กับ Yes/No
_Avoid_: mapper

**ORM**:
Object-Relational Mapping การจับ object หนึ่งตัวกับหนึ่งแถวในตาราง
_Avoid_: object-relational model

**JdbcTemplate**:
Class ของ Spring ที่เปิด connection, รัน SQL และปิด resource ให้
_Avoid_: Spring Data JDBC

**DataSource**:
Interface `javax.sql.DataSource` ที่หน้าที่คือให้ `Connection` ผ่าน `getConnection()`
_Avoid_: connection string

**Connection pool**:
ชุด connection ที่เปิดค้างไว้ให้ยืม การ `close` แปลว่าคืนเข้า pool
_Avoid_: connection cache

**HikariCP**:
Connection pool ที่ Spring Boot ใช้เป็นค่า default ชนิดจริงของ Bean คือ `HikariDataSource`
_Avoid_: Java connection pool

**RowMapper**:
Interface ที่บอกวิธีแปลงแถวหนึ่งเป็น object ผ่าน `mapRow`
_Avoid_: ORM

**DataAccessException**:
Unchecked exception ที่ Spring แปลมาจาก `SQLException`
_Avoid_: SQLException (เมื่อพูดถึงชั้นที่แอปจับ)

**JDBC**:
API มาตรฐานใน `java.sql` สำหรับคุยกับฐานข้อมูลเชิงสัมพันธ์
_Avoid_: database library (เมื่อหมายถึงตัว API นี้)

**JDBC Driver**:
ไลบรารีของผู้ผลิตฐานข้อมูลที่ implement interface ของ JDBC
_Avoid_: connector (เมื่อหมายถึงตัว interface)

**Connection**:
ช่องทางหนึ่งชุดที่เปิดไปยังฐานข้อมูล ผ่าน `DriverManager.getConnection`
_Avoid_: session (เมื่อหมายถึง object นี้)

**DriverManager**:
คลาสใน `java.sql` ที่เลือก driver จาก URL แล้วคืน `Connection`
_Avoid_: connection pool

**PreparedStatement**:
คำสั่งที่แยก template SQL กับค่าที่ใส่ผ่าน `?`
_Avoid_: class (มันเป็น interface)

**ResultSet**:
เคอร์เซอร์บนแถวที่ `SELECT` ส่งกลับ ต้องเรียก `next()` ก่อนอ่าน
_Avoid_: list (เมื่อหมายถึงตัวนี้)

**JPA**:
Specification (Jakarta Persistence) ที่นิยามวิธี map object ไปยัง relational database
_Avoid_: Spring JPA

**Persistence Context**:
ที่เก็บ Entity ที่ Hibernate กำลังติดตามในหนึ่ง transaction ตายเมื่อ transaction จบ
_Avoid_: cache (เมื่อหมายถึงทั้งก้อน)

**EntityManager**:
API ของ JPA สำหรับคุยกับ Persistence Context
_Avoid_: Session (เมื่อหมายถึงตัวที่แอปเรียก)

**First-level cache**:
Persistence Context ในบทบาทที่ไม่ยิง SELECT ซ้ำใน transaction เดียวกัน
_Avoid_: L2 cache, application cache

**Dirty checking**:
การเทียบ snapshot ตอนโหลดกับค่าใน heap แล้วสร้าง UPDATE ตอน flush
_Avoid_: dirty read

**Transaction**:
บล็อกของงานฐานข้อมูลที่ commit ทั้งก้อนหรือ rollback ทั้งก้อน
_Avoid_: session, request

**ACID**:
Atomicity, Consistency, Isolation และ Durability ของ transaction ใน relational database
_Avoid_: การแปลเป็นคำไทย

**`@Transactional`**:
Annotation ของ Spring ที่ให้ method รันใน transaction ผ่าน proxy
_Avoid_: `jakarta.transaction.Transactional` (เมื่อพูดถึงตัวที่มีกฎ rollback ของ Spring)

**Commit**:
ปิด transaction แล้วให้ผลทนต่อความล้มเหลวหลังจากนี้
_Avoid_: flush, save

**Rollback**:
ยกเลิกงานที่ยังไม่ commit ใน transaction นั้น
_Avoid_: reverse query

**Propagation**:
กฎว่า method ที่ถูกเรียกจะ join, suspend, สร้างใหม่ หรือปฏิเสธ transaction ของคนเรียก
_Avoid_: transaction level

**Isolation**:
ระดับที่กำหนดว่า transaction ที่รันพร้อมกันจะเห็นข้อมูลของอีกก้อนได้แค่ไหน
_Avoid_: lock level

**Physical transaction**:
Transaction จริงในฐานข้อมูล ก้อนที่ `REQUIRED` ใช้ร่วมกัน และก้อนที่ `REQUIRES_NEW` แยกออกไป
_Avoid_: connection

**Dirty read**:
การอ่านค่าที่อีก transaction ยังไม่ commit
_Avoid_: dirty data, dirty checking

**Non-repeatable read**:
อ่านแถวเดิมซ้ำแล้วได้คนละค่า เพราะมีคน commit ทับระหว่างทาง
_Avoid_: repeat read

**Phantom read**:
จำนวนแถวของ query เดิมเปลี่ยน เพราะมี transaction อื่น insert แล้ว commit
_Avoid_: phantom row

**Flush**:
การทำให้สถานะใน Persistence Context ตรงกับฐานข้อมูล
_Avoid_: commit

**Hibernate**:
Implementation ของ JPA ที่ Spring Boot ใช้เป็น default
_Avoid_: ORM (เมื่อหมายถึง Hibernate โดยเฉพาะ)

**Cascade**:
การให้ operation ของ Entity แม่กระจายไปยัง association ที่ระบุ เช่น `PERSIST` หรือ `REMOVE`
_Avoid_: chain delete

**FetchType**:
การกำหนดว่า association จะถูกโหลดพร้อมแถวหลัก (`EAGER`) หรือเมื่อมีการอ่านฟิลด์ (`LAZY`)
_Avoid_: FetchMode

**Lazy Loading**:
การคืน proxy มาก่อน แล้วค่อยส่ง query เมื่อมีการอ่านฟิลด์ของ association
_Avoid_: deferred load

**N+1 Problem**:
Query หลักหนึ่งครั้ง ตามด้วย query ของ association อีก N ครั้งตอนวนลิสต์
_Avoid_: loop query

**EntityGraph**:
การสั่งให้ query หนึ่งตัว initialize association ที่ระบุ แม้ mapping จะเป็น lazy
_Avoid_: fetch join

**Owning side**:
ฝั่งของความสัมพันธ์ที่ถือ foreign key Hibernate เขียนค่าความสัมพันธ์จากฝั่งนี้
_Avoid_: parent, owner object

**Inverse side**:
ฝั่งที่ใส่ `mappedBy` ไม่ได้ถือคอลัมน์ความสัมพันธ์ในตารางตัวเอง
_Avoid_: child

**`mappedBy`**:
ชื่อฟิลด์ Java ของ owning side ไม่ใช่ชื่อตารางหรือชื่อคอลัมน์
_Avoid_: mapped table

**Bidirectional**:
ความสัมพันธ์ที่มีฟิลด์ทั้งสองฝั่ง ฝั่งเดียวเรียกว่า Unidirectional
_Avoid_: two-way reference

**Spring Data**:
ร่มของรูปแบบ Repository ที่ใช้ข้ามฐานข้อมูลได้หลายแบบ
_Avoid_: Spring Data JPA (เมื่อหมายถึงทั้งร่ม)

**JpaRepository**:
Interface ที่รวม CRUD แบบคืน `List`, paging และ method ของ JPA อย่าง `flush`
_Avoid_: Repository (เมื่อหมายถึง marker ที่ว่าง)

**JPQL**:
ภาษา query ที่อ้างชื่อ Entity กับฟิลด์ Java
_Avoid_: SQL, HQL

**Pageable**:
คำขอหนึ่งหน้า มีเลขหน้าเริ่มที่ 0 กับขนาดหน้า
_Avoid_: offset

**Page**:
หนึ่งหน้าพร้อมจำนวนแถวทั้งหมด จึงมี query นับแยก
_Avoid_: Slice

**Spring Data JPA**:
Module ของ Spring ที่สร้าง Repository ครอบ JPA ไว้อีกชั้น
_Avoid_: Spring JPA, JPA (เมื่อหมายถึงตัว Repository)

**DTO**:
Class สำหรับย้ายข้อมูลข้ามชั้น ไม่ได้ map ไปเป็นตาราง
_Avoid_: model, entity (เมื่อหมายถึงก้อนที่ client ส่งหรือรับ)

**Request DTO**:
DTO ที่ถือฟิลด์ที่ client ส่งเข้ามา
_Avoid_: payload, request body class

**Response DTO**:
DTO ที่ถือฟิลด์ที่ส่งกลับให้ client
_Avoid_: response model

**Mapper**:
ส่วนที่แปลงระหว่าง DTO กับ Entity
_Avoid_: converter

**Bean Validation**:
การตรวจค่าด้วย annotation บนฟิลด์ของ object
_Avoid_: validator (เมื่อหมายถึงทั้งกลไก)

**`@RestControllerAdvice`**:
Class ตัวช่วยของ Controller ทุกตัวสำหรับ logic ร่วม โดยเฉพาะจับ exception แล้วตอบ JSON
_Avoid_: filter, interceptor

**`@ExceptionHandler`**:
Method ที่ Spring เรียกเมื่อเกิด exception ชนิดที่ระบุ โดยเลือกชนิดที่เฉพาะที่สุด
_Avoid_: catch

**Custom exception**:
Exception ที่สร้างเองเพื่อบอก business case แล้วให้ handler แปลงเป็น status code
_Avoid_: error code

**Exception response**:
Body รูปแบบเดียวทุกครั้งที่ error มี timestamp, status, error, message และ path
_Avoid_: error string

**Whitelabel Error Page**:
หน้า HTML สำรองเมื่อ client ขอ HTML และไม่มี view สำหรับ `/error`
_Avoid_: หน้า error ของ Tomcat

**`@Valid`**:
Annotation ที่สั่งให้ตรวจ Bean Validation ที่ argument นั้น
_Avoid_: `@Validated`

**Entity**:
Class ที่ `@Entity` map ไปเป็นตาราง ฟิลด์เป็นคอลัมน์ และ `@Id` เป็น primary key
_Avoid_: model, table class

**Repository**:
Interface ชั้นที่คุยกับฐานข้อมูล Spring Data สร้างตัวทำงานให้ตอน runtime
_Avoid_: DAO

**Soft Delete**:
การทำเครื่องหมายว่าแถวถูกลบ โดยไม่ลบแถวนั้นออกจากตาราง
_Avoid_: logical delete, flag delete

**Hard Delete**:
การเอาแถวออกจากตารางจริง
_Avoid_: physical delete

**Derived Query Method**:
Method ใน Repository ที่ Spring Data สร้าง query จากชื่อ method
_Avoid_: magic method, Spring JPA method
