# Extension functions

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 1:42:03–1:49:07

## Concept

`input.toInt()` ไม่ได้ส่ง `input` เข้าวงเล็บ แต่เรียกฟังก์ชันบนค่านั้นผ่านจุด อ่านเป็น `input.reversed()` ใกล้ภาษาคนกว่า `reverse(input)` พฤติกรรมไม่ต่างจากฟังก์ชันปกติที่มีพารามิเตอร์ตัวหนึ่ง สิ่งที่ต่างคือการอ่าน

ประกาศโดยใส่ type นำหน้าชื่อ

```kotlin
fun String.reverse(): String {
    val finalString = buildString {
        for (i in this@reverse.lastIndex downTo 0) {
            append(this@reverse[i])
        }
    }
    return finalString
}
```

วงเล็บว่างได้ เพราะสตริงที่ถูกเรียกไม่ได้นับเป็น argument ในวงเล็บ ถ้าต้องการค่าเพิ่ม เช่น ตัวเลข ค่อยใส่ในวงเล็บ เรียกได้บน type เดียวเท่านั้น

`this` ใน extension คือตัวที่ถูกเรียกบนนั้น เช่น สตริงของ `input.reverse()` บางจุดละ `this` ได้ คอมไพเลอร์เติมให้ เช่น `lastIndex` แต่ตัวดำเนินการ index ต้องเขียน `this[i]` เพราะเขียน `[i]` ลอย ๆ ไม่รู้ว่าชี้ก้อนไหน

ในคลิปเขาสะดุดเพราะ `buildString` มี `this` ของตัวเอง ถ้าเขียน `lastIndex` เฉย ๆ ภายในบล็อกนั้น มันไม่ชี้สตริงต้นทาง เขาเลยใช้ `this@reverse` เพื่อเจาะจง `this` ของฟังก์ชัน `reverse` รายละเอียดเขาให้รอตอน lambda

เรียกใช้เป็น

```kotlin
val reversed = input.reverse()
```

สตริงและจำนวนเต็มมี extension จากภาษาให้ใช้อยู่แล้ว พิมพ์จุดหลังค่าจะเห็นรายการ เช่น กรองข้อความ จัดรูปแบบ แปลงเป็นตัวพิมพ์เล็ก แทนที่ตัวอักษร และ `reversed()` ของ standard library ฟังก์ชันชื่อเดียวกันที่เราเขียนเองจะบังตัวของไลบรารี พอลบของเราออก `reversed()` เดิมจึงโผล่กลับมา เขาใส่ของตัวเองกลับเพื่อใช้สอน loop ต่อ

## Example

```kotlin
fun main() {
    val input = readln()
    val reversed = input.reverse()
    if (input == reversed) {
        println("that's a palindrome")
    }
}

fun String.reverse(): String {
    val finalString = buildString {
        for (i in this@reverse.lastIndex downTo 0) {
            append(this@reverse[i])
        }
    }
    return finalString
}
```

```java
public class Main {
    public static void main(String[] args) {
        String input = "racecar";
        String reversed = Strings.reverse(input);
        if (input.equals(reversed)) {
            System.out.println("that's a palindrome");
        }
    }
}

class Strings {
    static String reverse(String value) {
        StringBuilder finalString = new StringBuilder();
        for (int i = value.length() - 1; i >= 0; i--) {
            finalString.append(value.charAt(i));
        }
        return finalString.toString();
    }
}
```

Kotlin เติมฟังก์ชันให้ type เดิมแล้วเรียกด้วยจุดได้ ส่วน Java ไม่มี extension จึงแยกเป็น method `static` ที่รับค่าเป็นพารามิเตอร์
