# Loops

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 1:16:01–1:31:10

## Concept

โจทย์ในคลิปคือให้ผู้ใช้บอกว่าจะกรอกกี่จำนวน แล้วกรอกทีละตัว จากนั้นพิมพ์ผลรวม ถ้าจำนวนเป็น 5 ก็ต้องอ่านห้าครั้ง ถ้าเป็น 7 ก็อ่านเจ็ดครั้ง งานซ้ำแบบนี้ใช้ loop

**while** หน้าตาเหมือน `if` แต่รันซ้ำตราบใดที่เงื่อนไขยังเป็น `true` ตัวนับตั้งชื่อ `i` เป็นธรรมเนียม เริ่มที่ 0

```kotlin
var i = 0
while (i < amountOfNumbers) {
    i++
}
```

ถ้า `amountOfNumbers` เป็น 3 ลำดับคือ ตรวจ `0 < 3` แล้วรัน เพิ่มเป็น 1, ตรวจ `1 < 3` แล้วรัน เพิ่มเป็น 2, ตรวจ `2 < 3` แล้วรัน เพิ่มเป็น 3, ตรวจ `3 < 3` ไม่ผ่าน ออกจาก loop ถ้าเป็น 100 ก้อนข้างในจะรัน 100 ครั้ง

ในก้อนเขาอ่านเลขด้วย `readln().toIntOrNull() ?: 0` แล้ว `sum += number` โดย `var sum = 0` ถ้าใส่ข้อความที่ไม่ใช่ตัวเลข จะบวก 0 แต่ยังนับเป็นหนึ่งครั้ง ข้อความช่วยพิมพ์เป็น `please enter number #${i + 1}` เพราะ `i` เริ่มที่ 0

`continue` กระโดดไปตรวจเงื่อนไขรอบถัดไปทันที และไม่รันโค้ดที่เหลือในรอบนี้ เขาเอาไปแทนค่า 0 เมื่อแปลงไม่ได้ และยังไม่ `i++` ดังนั้นจะถามเลขหมายเลขเดิมซ้ำจนได้ตัวเลขจริง

`break` ออกจาก loop ทั้งหมด เขาบอกว่าโจทย์นี้ไม่ใช้ เพราะยังต้องรับเลขที่เหลือ

พิมพ์ array ตรง ๆ ได้ข้อความอ่านไม่รู้เรื่อง เขาใช้ `contentToString()` เพื่อดูสมาชิก และเตือนว่า `numbers += number` สร้าง array ใหม่ทุกรอบ

**Mutable list** เพิ่มสมาชิกได้โดยไม่สร้างก้อนใหม่ทั้งก้อน type ใส่ในมุมแหลม `mutableListOf<Int>()` เป็น `val` ได้ เพราะที่เปลี่ยนคือข้างใน ไม่ใช่ตัวแปรชี้ไปก้อนใหม่

```kotlin
val numbers = mutableListOf<Int>()
numbers.add(number)
println(numbers)
```

**for** ใช้เมื่อมีชุดค่า และอยากให้แต่ละรอบได้หนึ่งค่า ตัวนับแบบเดิมเขียนสั้นลงได้

```kotlin
for (i in 0 until amountOfNumbers) {
}
```

`until` ไม่รวมปลาย ถ้าใช้ `0..amountOfNumbers` ปลายถูกรวม จะรันเกินหนึ่งรอบ อีกทางคือ `1..amountOfNumbers` เขาเริ่มที่ 0 เพราะ index ของ list และ array เริ่มที่ 0 `for` ไม่มีเงื่อนไขให้ตรวจซ้ำ แต่ให้ค่า `i` ไล่ตามช่วง และนับให้เอง จึงไม่ต้องเขียน `i++`

วนสมาชิกของ list ได้ตรง ๆ สิ่งที่วนได้เขาเรียกว่า iterable

```kotlin
for (number in numbers) {
    println(number)
}
```

รอบแรกได้สมาชิกตัวแรก รอบถัดไปได้ตัวถัดไป `String` ก็เป็น iterable เพราะข้างในเป็น array ของตัวอักษร จึงเข้า index ได้เหมือน array

กลับสตริงทำโดยวนจากท้ายมาหน้า แล้วต่อตัวอักษรด้วย `buildString` เขาบอกว่ายังไม่ต้องสนใจรายละเอียดของ `buildString` ตอนนี้

```kotlin
val input = readln()
val finalString = buildString {
    for (i in input.lastIndex downTo 0) {
        append(input[i])
    }
}
println(finalString)
```

พิมพ์ `hello world` แล้วได้สตริงกลับด้าน

## Example

```kotlin
fun main() {
    println("how many numbers will you enter?")
    val amountOfNumbers = readln().toIntOrNull() ?: 0
    val numbers = mutableListOf<Int>()
    for (i in 0 until amountOfNumbers) {
        println("please enter number #${i + 1}")
        val number = readln().toIntOrNull()
        if (number == null) {
            continue
        }
        numbers.add(number)
    }
    var sum = 0
    for (number in numbers) {
        sum += number
    }
    println("the total sum is $sum")
    println(numbers)
}
```

```java
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println("how many numbers will you enter?");
        int amountOfNumbers = Integer.parseInt(scanner.nextLine());
        List<Integer> numbers = new ArrayList<>();
        for (int i = 0; i < amountOfNumbers; i++) {
            System.out.println("please enter number #" + (i + 1));
            int number = Integer.parseInt(scanner.nextLine());
            numbers.add(number);
        }
        int sum = 0;
        for (int number : numbers) {
            sum += number;
        }
        System.out.println("the total sum is " + sum);
        System.out.println(numbers);
    }
}
```

Kotlin เขียนช่วงด้วย `0 until n` และวน list ด้วย `for (number in numbers)` ส่วน Java เขียน `for (int i = 0; i < n; i++)` และวนด้วย `for (int number : numbers)` list ที่เพิ่มสมาชิกทีหลังใน Java ตัวอย่างนี้คือ `ArrayList`
