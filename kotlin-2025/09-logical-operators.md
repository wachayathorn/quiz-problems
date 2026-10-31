# Logical operators

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 27:19–32:40

## Concept

`&&` (and) และ `||` (or) ใช้ผูกหลายเงื่อนไข `Boolean` เข้าด้วยกัน โซ่ได้มากกว่าสองเงื่อนไข

`&&` เป็นจริงทั้งก้อนก็ต่อเมื่อทุกเงื่อนไขเป็นจริง ถ้ามีสักข้อเป็นเท็จ ทั้งก้อนเป็นเท็จ

```kotlin
val x = 12
val y = 14
val areBothEven = x % 2 == 0 && y % 2 == 0
```

ถ้า `y` เป็น 11 ผลเป็น `false` เพราะมีแค่ `x` ที่เป็นเลขคู่

`||` เขียนด้วยขีดตั้งสองอัน เป็นจริงเมื่อมีอย่างน้อยหนึ่งเงื่อนไขเป็นจริง ไม่ใช่ either-or ถ้าทั้งคู่เป็นจริง ผลยังเป็น `true` เขาเรียก either-or ว่า XOR และบอกว่าในงานจริง `||` ถูกใช้บ่อยกว่ามาก ถ้าทั้ง `x` และ `y` เป็น 11 ผลเป็น `false`

ลำดับความสำคัญมีเหมือนเลขคณิต `*` และ `/` มาก่อน `+`

```kotlin
val z = 3 + 3 * 3 // 12 ไม่ใช่ 18
```

`&&` มาก่อน `||` คอมไพเลอร์ใส่วงเล็บให้เงื่อนไขที่ผูกด้วย `&&` ก่อน แล้วค่อยเอาผลไป `||` กับเงื่อนไขแรก

`||` ยังตัดการประเมินที่เหลือได้ ถ้าข้อแรกเป็นจริงอยู่แล้ว ข้อหลังไม่เปลี่ยนผล จึงไม่ถูกดู ตัวอย่างของเขา `x` เป็น 12 ซึ่งเป็นเลขคู่ แล้วตามด้วย `&&` ที่ `x + y == 25` ซึ่งเป็นเท็จ พอกินด้วย `||` ทั้งก้อนยังเป็น `true`

ค่าที่กำหนดตายตัวตอน compile เขาเรียกว่า hardcoded IDE จะเตือนเมื่อเห็นว่าเงื่อนไขเป็นจริงเสมอเพราะรู้ค่าอยู่แล้ว ในโปรแกรมจริงค่าพวกนี้มักมาทีหลัง ไม่รู้ตอนเขียน

## Example

```kotlin
fun main() {
    val x = 12
    val y = 14
    val areBothEven = x % 2 == 0 && y % 2 == 0
    val atLeastOneEven = x % 2 == 0 || y % 2 == 0
    val mixed = x % 2 == 0 || y % 2 == 0 && x + y == 25
    println(areBothEven)
    println(atLeastOneEven)
    println(mixed)
}
```

```java
public class Main {
    public static void main(String[] args) {
        final int x = 12;
        final int y = 14;
        final boolean areBothEven = x % 2 == 0 && y % 2 == 0;
        final boolean atLeastOneEven = x % 2 == 0 || y % 2 == 0;
        final boolean mixed = x % 2 == 0 || y % 2 == 0 && x + y == 25;
        System.out.println(areBothEven);
        System.out.println(atLeastOneEven);
        System.out.println(mixed);
    }
}
```

`&&` กับ `||` เขียนและลำดับความสำคัญเหมือนกัน ทั้งคู่หยุดประเมินข้อถัดไปเมื่อรู้ผลแล้ว
