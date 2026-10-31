# Creating the project

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 05:15–09:20

## Concept

Build system รับซอร์สที่คนอ่านออก แล้วแพ็กเป็นไฟล์ที่รันได้ใน IDE JDK (Java Development Kit) คือชุดเครื่องมือ Java ที่ Kotlin พึ่ง เพราะ Kotlin ทำงานบน Java IntelliJ ติด JDK มาให้แล้ว เวอร์ชันในคลิปไม่สำคัญ เขาเปิดไว้ที่ 21 แต่ต่ำกว่านั้นก็ตามคลิปได้

ไฟล์แรกที่เปิดคือ `build.gradle.kts` เขาไม่ลงรายละเอียด Gradle ในคลิป จุดที่ควรรู้คือสตริงเวอร์ชัน Kotlin อยู่ในไฟล์นี้ ถ้าต้องเปลี่ยนภาษาเป็นเวอร์ชันใหม่ที่ JetBrains ปล่อย (ฟีเจอร์ใหม่ แก้บั๊ก) ให้แก้ที่นี่

โฟลเดอร์ที่สนใจมีแค่ `src/main/kotlin` ซึ่งเป็นที่วางไฟล์ `.kt` ส่วน `src/main/resources` ไว้รูปหรือไฟล์ภายนอก ซึ่งคลิปนี้ไม่ใช้

## Example

บทนี้เป็นขั้นตอนสร้างโปรเจกต์ ไม่มี syntax ให้เทียบกับ Java

1. กดสร้างโปรเจกต์ใหม่ หน้าตาจะต่างกันถ้าเพิ่งติดตั้ง
2. เลือกเทมเพลต Kotlin (Ktor เขาบอกว่าอยู่ใน Ultimate, Spring Boot ก็มีในรายการ แต่คลิปนี้ไม่ใช้)
3. ตั้งชื่อโปรเจกต์ เช่น `kotlin crash course`
4. ที่เก็บโปรเจกต์ใช้ค่าเริ่มต้น
5. Build system เลือก Gradle
6. JDK เลือกตัวที่ IntelliJ ติดมาให้
7. Gradle DSL เลือก Kotlin
8. ไม่ต้องใส่ sample code แล้วกด Create
9. ปิด `build.gradle.kts` ได้
10. คลิกขวาที่ `src/main/kotlin` แล้วสร้าง Kotlin Class/File ชื่อ `Main` เลือกชนิด File
