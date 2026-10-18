# Spring Boot Full Course 2026 — สรุปรายตอน

สรุปจาก playlist [Spring Boot Full Course 2026](https://www.youtube.com/playlist?list=PLEYgx5hMdopw) ของ Coder Army (Spring 7 / Spring Boot 4) หนึ่งไฟล์ต่อหนึ่งวิดีโอ ใช้อ่านควบคู่กับการดูวิดีโอ ทุกหัวข้อมี timestamp ที่ลิงก์ไปยังช่วงนั้นในคลิป

**วิธีอ่าน**

- เนื้อหาหลักมาจากซับไตเติลภาษาฮินดีที่ YouTube สร้างอัตโนมัติ
- `💡 เพิ่มเติม` = ผู้สรุปเติมเอง ไม่ได้อยู่ในวิดีโอ
- `⚠️ แก้จากวิดีโอ` = จุดที่วิดีโอพูดคลาดเคลื่อน
- `⚠️ Boot 4 vs 3` = จุดที่ Spring Boot 3.x เขียนต่างจากในคลิป
- โค้ดตัวอย่างใช้ Spring Boot 4.1 + Java 21
- ศัพท์หลักและนิยามอยู่ใน [CONTEXT.md](./CONTEXT.md)

## รายการตอน

| # | ตอน | สถานะ |
|---|---|---|
| 1 | [Spring Boot & Spring Framework: ภาพรวมก่อนเริ่ม](./01-spring-ecosystem-introduction.md) · [▶](https://youtu.be/Q04sTEwC7Kg) | ✅ |
| 2 | [สร้างแอป Spring Boot ตัวแรก](./02-first-spring-boot-application.md) · [▶](https://youtu.be/PVreyET_Q6I) | ✅ |
| 3 | [Apache Maven: POM, Dependencies และ Lifecycle](./03-apache-maven.md) · [▶](https://youtu.be/N2EXGMJVwUU) | ✅ |
| 4 | [Dependency Injection และ IoC](./04-dependency-injection-and-ioc.md) · [▶](https://youtu.be/dT10m_5POfs) | ✅ |
| 5 | [IoC Container: Bean, `@Component`, `@Autowired`, `@Bean`](./05-spring-ioc-container-and-beans.md) · [▶](https://youtu.be/Az9tIgCdaRU) | ✅ |
| 6 | [Circular Dependency, Bean Scope, Lazy/Eager](./06-circular-dependency-bean-scope-lazy-eager.md) · [▶](https://youtu.be/ass8cR52Cf8) | ✅ |
| 7 | [Bean Lifecycle](./07-bean-lifecycle.md) · [▶](https://youtu.be/WXg3uyZPqYI) | ✅ |
| 8 | [Spring XML Configuration](./08-spring-xml-configuration.md) · [▶](https://youtu.be/HqBh9hV7rik) | ✅ |
| 9 | [`@SpringBootApplication` และ Auto-configuration](./09-spring-boot-core-auto-configuration.md) · [▶](https://youtu.be/a06KMY209AA) | ✅ |
| 10 | [application.properties, `@Value` และ Runner](./10-application-properties-value-runners.md) · [▶](https://youtu.be/4naCY3vCSfk) | ✅ |
| 11 | [CRUD: สถาปัตยกรรมและ Create API](./11-crud-create-api-architecture.md) · [▶](https://youtu.be/6kCCIdLLGRs) | ✅ |
| 12 | [CRUD ลงฐานข้อมูลด้วย Spring Data JPA](./12-crud-mysql-spring-jpa.md) · [▶](https://youtu.be/uRr8_0x_HYw) | ✅ |
| 13 | [CRUD กับ Soft Delete](./13-crud-soft-delete-custom-jpa.md) · [▶](https://youtu.be/1lUBVZkDANg) | ✅ |
| 14 | [Servlet, Tomcat และ WAR](./14-servlets-tomcat-war.md) · [▶](https://youtu.be/TSHrq65q5dw) | ✅ |
| 15 | [Spring MVC: DispatcherServlet, HandlerMapping และ JSP](./15-spring-mvc-architecture.md) · [▶](https://youtu.be/M7_XxF6Zczw) | ✅ |
| 16 | [DTO และ Validation](./16-dtos-and-validations.md) · [▶](https://youtu.be/vXA78rtCeTY) | ✅ |
| 17 | [Exception Handling](./17-exception-handling.md) · [▶](https://youtu.be/1au7XK_b1sI) | ✅ |
| 18 | [Profiles และ YAML](./18-profiles-and-yaml.md) · [▶](https://youtu.be/K3H-rqhndEM) | ✅ |
| 19 | [Filter: FilterChain, log และ auth](./19-filters-filterchain.md) · [▶](https://youtu.be/ma-macdL9DI) | ✅ |
| 20 | [Filter: แก้ request, response, header และ body](./20-filters-modify-request-response.md) · [▶](https://youtu.be/fXcDynZ8C-g) | ✅ |
| 21 | [Interceptor: preHandle, postHandle, afterCompletion](./21-interceptors.md) · [▶](https://youtu.be/XE4xsgbArCQ) | ✅ |
| 22 | [Spring AOP: cross-cutting concern และ proxy](./22-spring-aop.md) · [▶](https://youtu.be/s4B-WbUWEFs) | ✅ |
| 23 | [AOP Advice: `@Before`, `@After`, `@Around`](./23-aop-advice-types.md) · [▶](https://youtu.be/hMsALeV_uG0) | ✅ |
| 24 | [AOP pointcut และชนิดของ proxy](./24-aop-pointcuts-and-proxies.md) · [▶](https://youtu.be/Kf2V2QYikVY) | ✅ |
| 25 | [Custom annotation ใน Spring AOP](./25-custom-annotations-aop.md) · [▶](https://youtu.be/TT-b9ob-lN8) | ✅ |
| 26 | [JDBC จากศูนย์](./26-jdbc-from-scratch.md) · [▶](https://youtu.be/IbCHISQOMUA) | ✅ |
| 27 | [Spring JDBC: JdbcTemplate, DataSource, HikariCP](./27-spring-jdbc.md) · [▶](https://youtu.be/OPjnbYbcb2E) | ✅ |
| 28 | [Hibernate fundamentals: CRUD และ entity mapping](./28-hibernate-fundamentals.md) · [▶](https://youtu.be/XVQ27D8IIWU) | ✅ |
| 29 | [Hibernate internals: Persistence Context และ cache](./29-hibernate-internals.md) · [▶](https://youtu.be/0Ko1Ux1otqE) | ✅ |
| 30 | [JPA Relationships](./30-jpa-relationships.md) · [▶](https://youtu.be/cMwR23y85qk) | ✅ |
| 31 | [Cascade, lazy loading และ N+1](./31-jpa-cascading-lazy-n-plus-one.md) · [▶](https://youtu.be/mB4NMgaZE8U) | ✅ |
| 32 | [Spring Data JPA: JpaRepository, JPQL, pagination](./32-spring-data-jpa.md) · [▶](https://youtu.be/xugZno16FgQ) | ✅ |
| 33 | [Transaction: ACID, `@Transactional`, commit และ rollback](./33-spring-transactions.md) · [▶](https://youtu.be/mpxdC8fKy1Y) | ✅ |
| 34 | [`@Transactional`: propagation และ isolation](./34-transactional-propagation-isolation.md) · [▶](https://youtu.be/E_Sb48BFoF0) | ✅ |
| 35 | [Spring Security: authentication และ authorization](./35-spring-security.md) · [▶](https://youtu.be/KVJuG95JSpM) | ✅ |
| 36 | [Spring Security: ฐานข้อมูลและรหัสผ่าน](./36-spring-security-db-authentication.md) · [▶](https://youtu.be/iF8KMT51QtY) | ✅ |
| 37 | [Spring Security: login flow](./37-spring-security-login-flow.md) · [▶](https://youtu.be/0LLpt3OZ03g) | ✅ |
| 38 | [JWT authentication](./38-jwt-authentication.md) · [▶](https://youtu.be/pUWp1EuVnGs) | ✅ |
| 39 | [OAuth 2.0: Google Login, OIDC และ PKCE](./39-oauth2-google-login.md) · [▶](https://youtu.be/t9kRN4wMVk0) | ✅ |
| 40 | [Testing: JUnit, Mockito และ MockMvc](./40-testing-junit-mockito-mockmvc.md) · [▶](https://youtu.be/8MAjlB7g6uI) | ✅ |
| 41 | Production-Grade Project · Microservices, Kafka & Redis · [▶](https://youtu.be/-Wv5uCu2PIg) | ❌ members-only เข้าถึงไม่ได้ |
| 42 | Production-Style API Gateway · GateKeeper Project · [▶](https://youtu.be/ZxCcYHL96x8) | ❌ members-only เข้าถึงไม่ได้ |
