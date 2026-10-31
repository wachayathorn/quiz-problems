# Arithmetic operators

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 18:20–22:58

## Concept

ตัวดำเนินการคำนวณมีบวก ลบ คูณ หาร และ modulo

```kotlin
val x = 5
val y = 9
println(x + y) // 14
println(x - y) // -4
```

หารจำนวนเต็มด้วยจำนวนเต็มได้จำนวนเต็มเสมอ ส่วนทศนิยมถูกตัดทิ้ง เขาใช้คำว่าปัดลง `5 / 9` ได้ `0`, `9 / 9` ได้ `1`, `10 / 9` ได้ `1` การคูณใช้ `*` เช่น `10 * 9` ได้ `90`

Modulo ใช้ `%` คือเศษที่เหลือหลังหาร `11 % 9` ได้ `2` เพราะ 9 ลงตัวใน 11 ได้ครั้งเดียว เหลือ 2

ถ้าตัวแปรเป็น `var` มีทางลัด

- `x++` คือ `x = x + 1`
- `x--` คือ `x = x - 1`
- `x += 5` คือ `x = x + 5`
- มี `-=`, `/=`, `*=` ในความหมายเดียวกัน

`val` ใช้ `++` ไม่ได้ เพราะเป็นการคำนวณแล้ว assign กลับ IDE จะฟ้องว่า `Val cannot be reassigned`

`x + 5` คนละอย่างกับ `x += 5` บรรทัดแรกคำนวณแล้วทิ้งผล บรรทัดหลังเขียนผลกลับเข้า `x` ตัวอย่างของเขา `x` เริ่มที่ 11 แล้ว `x += 5` จากนั้นพิมพ์ `x + 9` ได้ 25 ถ้าเขียนแค่ `x + 5` โดยไม่ assign ผลบวกกับ 9 ยังเป็น 20

## Example

```kotlin
fun main() {
    var x = 11
    val y = 9
    println(x / y)
    println(x % y)
    x += 5
    println(x + y)
}
```

```java
public class Main {
    public static void main(String[] args) {
        int x = 11;
        final int y = 9;
        System.out.println(x / y);
        System.out.println(x % y);
        x += 5;
        System.out.println(x + y);
    }
}
```

เครื่องหมาย `+`, `-`, `*`, `/`, `%`, `++` และ `+=` เขียนเหมือนกัน ทั้งสองภาษาตัดเศษทิ้งเมื่อหารจำนวนเต็มด้วยจำนวนเต็ม
