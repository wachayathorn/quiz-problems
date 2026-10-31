# Arrays

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 1:05:30–1:16:01

## Concept

ก่อนหน้านี้ตัวแปรหนึ่งตัวเก็บค่าเดียว Array เป็น primitive อีกแบบ เป็นกล่องขนาดคงที่สำหรับค่าชนิดเดียวกันหลายค่า

```kotlin
val favoriteNumbers = intArrayOf(1, 2, 3, 69)
```

`intArrayOf` รับ argument กี่ตัวก็ได้ของ type `Int` ชี้ที่ตัวแปรจะเห็น type เป็น `IntArray`

ดึงสมาชิกด้วย index ในวงเล็บเหลี่ยม การนับเริ่มที่ 0 ดังนั้น index 0 คือ 1, index 1 คือ 2, index 3 คือ 69 ถ้าขอ index 4 ซึ่งเป็นตัวที่ห้า ทั้งที่มีแค่สี่ตัว จะได้ `ArrayIndexOutOfBoundsException`

`getOrNull(4)` คืน `null` แทนการแครชเมื่อ index ไม่มี

เมื่อ index มาจากผู้ใช้ เขาอ่านด้วย `toIntOrNull()` แล้วเข้าบล็อกเมื่อไม่เป็น `null` ใส่ค่าในสตริงได้ด้วย `"your number is ${favoriteNumbers[inputAsInteger]}"` พิมพ์จุดหลังชื่อ array แล้วเลือกวงเล็บเหลี่ยม IDE จะจัดเป็น placeholder ให้

เช็กแค่ `inputAsInteger < favoriteNumbers.size` ยังไม่พอ `size` คือจำนวนสมาชิก ซึ่งเป็น 4 แต่ index สูงสุดที่ใช้ได้คือ 3 และค่าติดลบเช่น `-1` เป็น `Int` ที่ถูกต้องแต่ไม่มีช่องนั้น เขาเลยใช้ช่วง

```kotlin
if (inputAsInteger in 0..favoriteNumbers.lastIndex) {
    println("your number is ${favoriteNumbers[inputAsInteger]}")
} else {
    println("that index doesn't exist")
}
```

`..` รวมปลายทั้งสองข้าง `lastIndex` คือ index สูงสุดที่ยังถูกต้อง อย่าใช้ `size` เป็นปลายช่วงเพราะ 4 จะถูกรวมทั้งที่ไม่มี index 4 เขาบอกให้ตั้งใจลองพังโค้ดตอนพัฒนา จะได้รู้ว่าต้องกันอะไรก่อนผู้ใช้จริงเจอ

ขนาดถูกตรึงตอนสร้าง มี factory อื่นเช่น `floatArrayOf`, `doubleArrayOf`, `arrayOf` สำหรับ type ทั่วไปรวม `String`, `booleanArrayOf`, `byteArrayOf`

`favoriteNumbers + 5` ดูเหมือนแทรกสมาชิก แต่ array แรกยังขนาดเท่าเดิม สิ่งที่เกิดคือสร้าง array ก้อนใหม่ที่จุห้าตัวได้ เขาจึงเรียก array ว่า immutable ในแง่จำนวนสมาชิก

สมาชิกเดิมแก้ได้โดยไม่สร้างก้อนใหม่

```kotlin
favoriteNumbers[2] = 69
```

รันแล้วขอ index 2 จะได้ 69 เขาใช้จุดนี้พาไปสู่คำถามว่า ถ้าต้องเก็บเลขที่ผู้ใช้พิมพ์เรื่อย ๆ โดยยังไม่รู้ว่ากี่ตัว เช่น จนกว่าผลรวมจะถึง 20 จะเพิ่มช่อง array ง่าย ๆ ไม่ได้ และไม่อยากเขียน `readln()` สิบบรรทัด คำตอบอยู่ที่ loop ในบทถัดไป ซึ่งจะมีโครงสร้างที่ยืดหยุ่นกว่า array ในบางงาน array ก็ยังเหมาะกว่าในบางงาน

## Example

```kotlin
fun main() {
    val favoriteNumbers = intArrayOf(1, 2, 3, 69)
    favoriteNumbers[2] = 69
    val inputAsInteger = readln().toIntOrNull()
    if (inputAsInteger != null && inputAsInteger in 0..favoriteNumbers.lastIndex) {
        println("your number is ${favoriteNumbers[inputAsInteger]}")
    } else {
        println("that index doesn't exist")
    }
}
```

```java
public class Main {
    public static void main(String[] args) {
        int[] favoriteNumbers = {1, 2, 3, 69};
        favoriteNumbers[2] = 69;
        int inputAsInteger = 3;
        if (inputAsInteger >= 0 && inputAsInteger <= favoriteNumbers.length - 1) {
            System.out.println("your number is " + favoriteNumbers[inputAsInteger]);
        } else {
            System.out.println("that index doesn't exist");
        }
    }
}
```

ทั้งคู่ใช้วงเล็บเหลี่ยมและเริ่มนับที่ 0 Kotlin สร้างด้วย `intArrayOf` แล้วเช็กช่วงด้วย `in 0..lastIndex` ส่วน Java สร้างด้วย `{ }` แล้วเทียบขอบเองด้วย `length - 1`
