# Hello world program

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 09:20–11:32

## Concept

ไฟล์ Kotlin ใช้นามสกุล `.kt` โปรแกรมที่ง่ายที่สุดรับข้อความเข้าแล้วพิมพ์ข้อความออกทาง console ไม่มีหน้าจอกราฟิก

`fun main() { ... }` คือจุดที่โปรแกรมเริ่มเมื่อกดรัน ปีกกา `{ }` บอกว่าโค้ดชุดไหนอยู่ใน `main` ความหมายของฟังก์ชันเขาเก็บไว้บทหลัง

`println` ส่งข้อความไปที่ console ของ IDE ไม่ใช่เครื่องพิมพ์กระดาษ ข้อความที่พิมพ์ตรง ๆ ใส่เครื่องหมายคำพูด กดลูกศรเขียวที่ไฟล์ `Main.kt` แล้ว console จะขึ้น `Hello World!`

IDE อาจเปิด AI autocomplete ไว้ ถ้าเครื่องคุณไม่มี ให้เปิดใน settings

## Example

```kotlin
fun main() {
    println("Hello World!")
}
```

```java
public class Main {
    public static void main(String[] args) {
        System.out.println("Hello World!");
    }
}
```

Kotlin วาง `main` ไว้ระดับไฟล์ได้เลย ส่วน Java ต้องห่อไว้ใน class และเขียน `System.out.println`
