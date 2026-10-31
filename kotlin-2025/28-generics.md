# Generics

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:53:06–จบคลิป

## Concept

Generic คือพารามิเตอร์ชนิด type สิ่งที่เปลี่ยนระหว่างการเรียกไม่ใช่ค่า แต่เป็น type ที่ฟังก์ชันหรือ class ทำงานด้วย เขาเคยใช้ไปแล้วโดยยังไม่เรียกชื่อนี้ เช่น `filter`

`filter` บน `IntArray` ในคลิปผูกกับ `IntArray` โดยตรง และตัวบน `String` ก็ผูกกับ `String` แต่บน `List` ลายเซ็นเป็น extension ของ `Iterable<T>` ไม่ใช่ `Iterable<Int>` เรียกบนลิสต์จำนวนเต็มได้ ทั้งที่ `T` ไม่ได้เขียนว่าเป็น `Int` เพราะ `T` คือ type ที่ยังไม่ตายตัว ขอแค่วนได้ จะเป็นลิสต์ของจำนวนเต็ม สตริง หรือสี่เหลี่ยม ก็ใช้ `filter` ตัวเดียวกัน แล้วได้ลิสต์ชนิดเดิมกลับไป ไม่ต้องเขียน `filter` แยกทุก type

ฟังก์ชันของเราประกาศ type parameter ในมุมแหลมก่อนชื่อ ชื่อ `T` เป็นธรรมเนียมเมื่อมีตัวเดียว ใช้ชื่ออื่นเช่น `Type` ได้ แต่เขาให้ใช้ `T`

```kotlin
fun <T> List<T>.myFilter(predicate: (T) -> Boolean): List<T>
```

ลิสต์ที่ถูกเรียกเป็น `List<T>` lambda ได้สมาชิกชนิด `T` และผลที่คืนก็เป็น `List<T>` เรียกบน `List<String>` รอบนั้น `T` คือ `String` เรียกบน `List<Int>` รอบนั้น `T` คือ `Int`

ข้างในเขาสร้าง `mutableListOf<T>()` เดิน `this` แล้ว `add` เฉพาะตัวที่ `predicate` คืน `true` สุดท้าย `toList()` เพื่อให้ฝั่งเรียกได้ลิสต์ที่เติมตรง ๆ ไม่ได้

เขาลองสองลิสต์กับฟังก์ชันเดียว

- สตริง `"hello world"`, `"bye-bye"`, `"how is it going"` เก็บเฉพาะความยาวอย่างน้อย 10 ตัว `"bye-bye"` หลุด
- จำนวน `1, 2, 3, 4, 5` เก็บเฉพาะที่น้อยกว่า 3

ใน lambda ชี้ที่ชื่อพารามิเตอร์จะเห็น type ตามลิสต์ที่เรียก

Class ก็รับ type parameter ได้ ตัวอย่างที่เขาใช้บ่อยบน Android คือผลลัพธ์ที่มีได้สองทาง เช่น เรียกเซิร์ฟเวอร์แล้วสำเร็จพร้อมข้อมูล หรือล้มพร้อม error แต่ละ endpoint คืนข้อมูลคนละ type จึงไม่ฝัง type ตาย เขาใช้ชื่อ `D` กับ `E` เมื่อมีสองตัว เพื่อไม่ให้ปนกัน และชนชื่อ `Result` ของ Kotlin เอง เขาเลยเอา import ที่ชนออก

```kotlin
sealed interface Result<out D, out E> {
    data class Success<D>(val data: D) : Result<D, Nothing>
    data class Failure<E>(val error: E) : Result<Nothing, E>
}
```

`Success` ถือข้อมูลชนิด `D` ไม่มี error จึงใส่อีกฝั่งเป็น `Nothing` ซึ่งเป็น class พิเศษสำหรับกรณีที่ไม่มีค่า `Failure` ถือ error ชนิด `E` และไม่มีข้อมูลสำเร็จ เขาบอกว่าอันนี้เป็น sealed interface ได้เพราะไม่มี constructor ที่ต้องเรียก

```kotlin
fun makeNetworkCall(): Result<Int, String> {
    return Result.Failure("something went wrong")
}
```

อีกครั้งหนึ่งของฟังก์ชันนี้อาจคืน `Rectangle` หรือลิสต์ข้อมูลก็ได้ แล้วแต่ endpoint ตอนแรกการคืน `Failure` ฟ้องว่าได้ `Result.Failure<String>` แต่จุดคืนต้องการ `Result<Int, String>` ทั้งที่ `Failure` เป็นลูกของ `Result` เขาใส่ `out` หน้า type parameter แล้วผ่าน `out` ทำให้ไม่จำกัดแค่ type นั้นเป๊ะ แต่รับลูกของ type นั้นได้ด้วย

`map` ใช้ generic สองตัว `T` คือชนิดสมาชิกของ iterable ที่ถูกเรียก `R` คือชนิดที่แปลงไป เขาชี้ว่า lambda ที่คืน `toDouble()` ทำให้ `R` ถูก infer เป็น `Double` ลิสต์ผลจึงเป็น `List<Double>` ทั้งที่ต้นทางเป็นจำนวนเต็ม

เขาปิดคลิปว่า generic กับ lambda ไม่จำเป็นต้องเข้าใจทุกมุมในรอบเดียว การเรียนโปรแกรมใช้เวลา สิ่งที่ควรทำต่อคือเอาไปฝึก playlist เก่าของเขามีการบ้านท้ายคลิป หรือไปแก้โจทย์สั้น ๆ ที่ codewars.com พอคล่องกับสิ่งที่คลิปนี้ครอบคลุมแล้ว ค่อยเลือกทาง เช่น แอป Android ตัวแรก, backend ด้วย Ktor หรือ Spring Boot บน Kotlin ซึ่งเขาบอกว่าจะทำคลิปต่อ

## Example

```kotlin
fun <T> List<T>.myFilter(predicate: (T) -> Boolean): List<T> {
    val result = mutableListOf<T>()
    for (element in this) {
        if (predicate(element)) {
            result.add(element)
        }
    }
    return result.toList()
}

sealed interface Result<out D, out E> {
    data class Success<D>(val data: D) : Result<D, Nothing>
    data class Failure<E>(val error: E) : Result<Nothing, E>
}

fun makeNetworkCall(): Result<Int, String> {
    return Result.Failure("something went wrong")
}

fun main() {
    val strings = listOf("hello world", "bye-bye", "how is it going")
    val integers = listOf(1, 2, 3, 4, 5)
    println(strings.myFilter { current -> current.length >= 10 })
    println(integers.myFilter { current -> current < 3 })
    println(integers.map { it.toDouble() })
    println(makeNetworkCall())
}
```

```java
import java.util.ArrayList;
import java.util.List;
import java.util.function.Predicate;

public class Main {
    static <T> List<T> myFilter(List<T> values, Predicate<T> predicate) {
        List<T> result = new ArrayList<>();
        for (T element : values) {
            if (predicate.test(element)) {
                result.add(element);
            }
        }
        return List.copyOf(result);
    }

    public static void main(String[] args) {
        List<String> strings = List.of("hello world", "bye-bye", "how is it going");
        List<Integer> integers = List.of(1, 2, 3, 4, 5);
        System.out.println(myFilter(strings, current -> current.length() >= 10));
        System.out.println(myFilter(integers, current -> current < 3));
        System.out.println(integers.stream().map(Integer::doubleValue).toList());
    }
}
```

Kotlin ประกาศ `<T>` หน้าชื่อฟังก์ชันแล้วเรียกเป็น extension บน `List<T>` ส่วน Java ประกาศ `<T>` หน้า type ที่คืน แล้วส่งลิสต์เข้าเป็นพารามิเตอร์

💡 เพิ่มเติม: `out` ในคลิปทำให้คืนลูกของ `Result` ตรงจุดที่ขอ `Result<Int, String>` ได้ Java ไม่ได้ใส่ความยืดนี้ที่ตัว class เอง ต้องเขียน `? extends` ที่จุดใช้แทน และไม่มี type ชื่อ `Nothing`
