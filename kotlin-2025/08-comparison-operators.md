# Comparison operators

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 22:58–27:19

## Concept

`=` คือการ assign `==` คือการเทียบว่าเท่ากันหรือไม่ ผลไม่ใช่ตัวเลข แต่เป็น `Boolean` คือ `true` หรือ `false` เท่านั้น เทียบได้ทั้งตัวเลขและค่าชนิดอื่น เขาเริ่มด้วยตัวเลข

`11 == 9` เป็น `false` ตัวอื่นที่เขาใช้

- `>` มากกว่า `11 > 9` เป็น `true` แต่ `11 > 11` เป็น `false`
- `>=` มากกว่าหรือเท่ากับ รวมค่าที่เท่ากันด้วย
- `<` และ `<=` ในทิศตรงข้าม

เก็บผลลงตัวแปรได้ type ที่ infer ได้คือ `Boolean`

```kotlin
val x = 11
val y = 11
val areNumbersTheSame = x == y
println(areNumbersTheSame) // true
```

เอาไปปนกับ arithmetic ได้ การเช็กเลขคู่คือเศษจากการหารด้วย 2 ต้องเป็น 0 เศษได้แค่ 0 หรือ 1

```kotlin
println(x % 2 == 0) // 11 เป็นเลขคี่ ได้ false
```

String template ดึงค่าตัวแปรเข้าข้อความด้วย `$` ถ้าเป็นนิพจน์ที่ยาวกว่าชื่อตัวแปร ให้หุบด้วย `${ }`

```kotlin
val x = 12
println("is $x an even number: ${x % 2 == 0}")
```

ถ้าเขียน `"is x an even number"` โดยไม่มี `$` จะได้ตัวอักษร x ไม่ใช่ค่า 12 เมื่อ `x` เป็น 12 ผลเป็นเลขคู่

## Example

```kotlin
fun main() {
    val x = 12
    val y = 11
    val areNumbersTheSame = x == y
    println(areNumbersTheSame)
    println("is $x an even number: ${x % 2 == 0}")
}
```

```java
public class Main {
    public static void main(String[] args) {
        final int x = 12;
        final int y = 11;
        final boolean areNumbersTheSame = x == y;
        System.out.println(areNumbersTheSame);
        System.out.println("is " + x + " an even number: " + (x % 2 == 0));
    }
}
```

Kotlin แทรกค่าเข้าสตริงด้วย `$` และ `${ }` ส่วน Java ในตัวอย่างนี้ต่อสตริงด้วย `+`

💡 เพิ่มเติม: ใน Java การเทียบ `String` ด้วย `==` เทียบตัวอ้างอิง ไม่ใช่ข้อความ Kotlin ใช้ `==` เรียกการเทียบค่าให้ บทนี้ Philipp เทียบแค่ตัวเลข
