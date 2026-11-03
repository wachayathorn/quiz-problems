# Nullability and null-safe operators

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 36:41–47:27

## Concept

`null` แปลว่าค่านั้นไม่มีอยู่ ไม่ใช่ `0` และไม่ใช่เลขสุ่ม ทุก type มีรูปที่ลงท้ายด้วย `?` ได้ เช่น `Int?` แปลว่าอาจเป็น `Int` หรืออาจเป็น `null`

`toIntOrNull()` คืนจำนวนเต็มเมื่อแปลงได้ และคืน `null` เมื่อแปลงไม่ได้ พิมพ์ `45` ได้ `45` พิมพ์ `hello world` ได้ `null` โดยไม่แครชที่บรรทัดแปลง

เอา `Int?` ไป `% 2` ตรง ๆ ไม่ได้ เพราะถ้าค่าไม่มีอยู่ ก็ไม่มีผลของ modulo ให้คำนวณ เขาสอนตัวดำเนินการสามตัว

**`!!`** ยืนยันว่าค่าไม่เป็น `null` ถ้าเป็น `null` จะแครชด้วย `NullPointerException` เขาบอกว่าไม่ได้อะไรเพิ่ม และไม่ควรใช้ เพราะอยากได้ null safety ไม่ใช่แครชสุ่ม

**`?:` (Elvis operator)** เขาให้หันหัวดู จะเห็นหน้า Elvis ฝั่งซ้ายถูกใช้เมื่อไม่เป็น `null` ถ้าเป็น `null` จะใช้ค่าทางขวาแทน

```kotlin
val inputAsInteger = input.toIntOrNull() ?: 0
```

พิมพ์ `hello world` แล้วเช็กเลขคู่ จะได้ `true` เพราะตกไปที่ `0`

**`?.`** เรียกฟังก์ชันต่อเมื่อค่าข้างหน้าไม่เป็น `null` ถ้าเป็น `null` ทั้งนิพจน์เป็น `null` และไม่เรียกฟังก์ชันนั้น

```kotlin
val incremented = input.toIntOrNull()?.inc()
```

`inc()` คือการบวกหนึ่งแบบฟังก์ชัน คู่กับ `++` พิมพ์ `5` ได้ `6` พิมพ์ `hello` ได้ `null`

`rem` คือฟังก์ชันของ `%` ใช้คู่กับ `?.` ได้ เพราะหลัง `?.` ต้องเป็นฟังก์ชัน

```kotlin
val isEven = input.toIntOrNull()?.rem(2)?.equals(0)
```

เลขคู่ได้ `true` เลขคี่ได้ `false` ข้อความที่ไม่ใช่ตัวเลขได้ `null` เพราะเทียบ `null` กับ `0` ไม่ถูกทำเมื่อขั้นก่อนเป็น `null` แล้ว ถ้าเทียบ `null == 0` ตรง ๆ ผลเป็น `false` เพราะ `null` ไม่เท่ากับศูนย์

คอมไพเลอร์ห้ามเรียกฟังก์ชันบนค่าที่อาจเป็น `null` ตั้งแต่ตอน compile เขาเรียกกลไกนี้ว่า null safety ภาษาส่วนใหญ่ไม่มีแบบนี้ Java เขียนเรียกบน `null` ได้ แต่พอรันจะแครช ซึ่งตรงกับพฤติกรรมของ `!!`

คอมเมนต์ทำได้ด้วยการตัดโค้ดออกจากโปรแกรม เขาใช้ตอนเอาบรรทัดเช็กเลขคู่ออกชั่วคราว

## Example

```kotlin
fun main() {
    println("please enter a number")
    val input = readln()
    val inputAsInteger = input.toIntOrNull() ?: 0
    val incremented = input.toIntOrNull()?.inc()
    val isEven = input.toIntOrNull()?.rem(2)?.equals(0)
    println(inputAsInteger)
    println(incremented)
    println(isEven)
}
```

```java
public class Main {
    public static void main(String[] args) {
        String input = "hello world";
        Integer parsed;
        try {
            parsed = Integer.parseInt(input);
        } catch (NumberFormatException exception) {
            parsed = null;
        }
        int inputAsInteger = parsed == null ? 0 : parsed;
        Integer incremented = parsed == null ? null : parsed + 1;
        Boolean isEven = parsed == null ? null : parsed % 2 == 0;
        System.out.println(inputAsInteger);
        System.out.println(incremented);
        System.out.println(isEven);
    }
}
```

Kotlin ติด `?` ที่ type แล้วใช้ `?:` กับ `?.` ให้คอมไพเลอร์กันการเรียกบน `null` ส่วน Java ใช้ type ที่เป็น object แล้วตรวจ `null` เอง ถ้าลืมตรวจจะได้ `NullPointerException` ตอนรัน
