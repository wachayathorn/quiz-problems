# Function overloading

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 1:49:07–1:52:25

## Concept

Overload คือมีฟังก์ชันชื่อเดียวกันหลายตัว แต่คนละชุดพารามิเตอร์ ตอนนี้ `reverse` เรียกบน `String` ได้ แต่เรียกบน `Int` ไม่ได้ จึงเขียนอีกตัวที่รับ `Int` แล้วคืน `Int`

```kotlin
fun Int.reverse(): Int {
    return this.toString().reverse().toInt()
}
```

จำนวนเต็มแปลงเป็นสตริงได้เสมอ กลับทางข้ามไม่ได้เพราะไม่ใช่ทุกสตริงจะเป็นจำนวนเต็ม จากนั้นเรียก `reverse` ของ `String` ที่มีอยู่ แล้วแปลงกลับเป็น `Int`

เขาลอง `45` แล้วบวกกับตัวที่กลับด้าน ได้ `99` เพราะ `45 + 54`

แนวนี้ใช้กับฟังก์ชันปกติได้เช่นกัน ถ้าพารามิเตอร์เป็น `String` ตัวหนึ่ง และอีกตัวเป็นจำนวน ภายในตัวหลังจะอ้างชื่อพารามิเตอร์ของตัวเอง ไม่ใช่ `this`

สิ่งที่ทำไม่ได้คือสองฟังก์ชันที่มีชุดพารามิเตอร์และค่าที่คืนเหมือนกันทุกอย่าง เช่น `reverse` สองตัวที่รับ `String` แล้วคืน `String` ตอนเรียกด้วยสตริง คอมไพเลอร์ไม่รู้ว่าจะเข้าตัวไหน แต่คนละ type ชัดเจน เรียกด้วย `Int` เข้าตัวหนึ่ง เรียกด้วย `String` เข้าอีกตัว

## Example

```kotlin
fun String.reverse(): String {
    val finalString = buildString {
        for (i in this@reverse.lastIndex downTo 0) {
            append(this@reverse[i])
        }
    }
    return finalString
}

fun Int.reverse(): Int {
    return this.toString().reverse().toInt()
}

fun main() {
    val input = readln().toInt()
    println(input + input.reverse())
}
```

```java
public class Main {
    static String reverse(String value) {
        StringBuilder finalString = new StringBuilder();
        for (int i = value.length() - 1; i >= 0; i--) {
            finalString.append(value.charAt(i));
        }
        return finalString.toString();
    }

    static int reverse(int value) {
        return Integer.parseInt(reverse(Integer.toString(value)));
    }

    public static void main(String[] args) {
        int input = 45;
        System.out.println(input + reverse(input));
    }
}
```

ทั้งคู่แยกฟังก์ชันชื่อ `reverse` ตาม type ที่รับ Kotlin เขียนเป็น extension บน `String` กับ `Int` ส่วน Java เขียนเป็น method ใน class เดียวกันโดยดูจาก type ของพารามิเตอร์
