# Exceptions and try/catch

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 59:48–1:05:30

## Concept

Exception คือทางทำให้โปรแกรมหยุดเมื่อเกิดสิ่งที่โค้ดจัดการตรง ๆ ไม่ได้ `toInt()` คืน `Int` ที่ไม่เป็น `null` จึงยังแครชได้ถ้าพิมพ์ `hello` ข้อความคือ `NumberFormatException` เพราะแปลงสตริงนั้นเป็นจำนวนเต็มไม่ได้ IDE อาจเตือนเป็นสีเหลืองที่ตัวแปร แต่ไม่ได้กันการพิมพ์ค่าผิด

Stack trace คือรายการฟังก์ชันที่ถูกเรียกจนมาถึงจุดแครช ลิงก์สีเทาเป็นฟังก์ชันภายในที่เราไม่ได้เขียน ลิงก์สีน้ำเงินชี้ไฟล์ของเรา คลิกแล้วไปบรรทัดที่พัง ซึ่งในตัวอย่างคือบรรทัด `toInt()`

Exception ที่ไม่ถูกจับเรียกว่า uncaught และจะแครช ถ้าพอรู้ว่าจุดนี้พังได้ ให้หุบด้วย `try` / `catch`

```kotlin
val inputAsInteger = try {
    input.toInt()
} catch (exception: NumberFormatException) {
    0
}
```

บล็อก `catch` รันเมื่อโค้ดใน `try` โยน `NumberFormatException` ค่าใน `catch` ต้องเป็น `Int` ด้วย เพราะตัวแปรสุดท้ายต้องเป็นจำนวนเต็ม เขาเอา `0` นี้ไปเข้า `when` เดิม แล้วตกสาขา `else` เพราะศูนย์ไม่ตรงสามสาขาแรก

ชี้ที่ `toInt()` แล้วอ่านเอกสาร จะเห็นว่ามันโยน exception นี้เมื่อสตริงไม่ใช่รูปของจำนวน

`finally` ต่อท้ายได้ โค้ดในนั้นรันแน่นอนแม้มี exception งานจริงใช้ปิด network stream หรือ file stream เขาบอกว่ายังไม่ใช่พื้นฐานของคลิปนี้ แล้วลบออก

`catch (exception: Exception)` จับได้ทุกชนิด ถ้าจับเฉพาะ `NumberFormatException` แล้วมีชนิดอื่นถูกโยน บล็อกนี้จะไม่รับ

โยนเองด้วย `throw`

```kotlin
throw Exception("dude, what did you pass there")
```

ถ้าโยนใน `main` แล้วไม่จับ โปรแกรมแครชทันที และ IDE รู้ว่าโค้ดหลัง `throw` รันไม่ถึง เขาบอกว่าการใช้จริงไม่ใช่การโยนทิ้งไว้กลาง `main` แต่มักโยนจากฟังก์ชันที่ลึกกว่า เมื่อรับพารามิเตอร์ที่ใช้ต่อไม่ได้ แล้วบอกผู้เรียกว่าค่าที่ส่งมาใช้ไม่ได้

## Example

```kotlin
fun main() {
    val input = readln()
    val inputAsInteger = try {
        input.toInt()
    } catch (exception: NumberFormatException) {
        0
    }
    println(inputAsInteger)
}
```

```java
public class Main {
    public static void main(String[] args) {
        String input = "hello";
        int inputAsInteger;
        try {
            inputAsInteger = Integer.parseInt(input);
        } catch (NumberFormatException exception) {
            inputAsInteger = 0;
        }
        System.out.println(inputAsInteger);
    }
}
```

โครง `try` / `catch` เหมือนกัน Kotlin เขียน type ของ exception หลัง colon และให้ทั้ง `try` เป็นค่าที่ assign ได้ ส่วน Java ประกาศตัวแปรก่อน แล้ว assign ในแต่ละบล็อก

💡 เพิ่มเติม: Java มี checked exception ที่ต้องประกาศหรือจับ Kotlin ไม่บังคับแบบนั้น Philipp ไม่ได้พูดเรื่องนี้ในคลิป
