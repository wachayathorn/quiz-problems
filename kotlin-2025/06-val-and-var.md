# val and var

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 11:32–18:20

## Concept

ตัวแปรเก็บค่าได้เหมือนตัวแปรในวิชาคณิตศาสตร์ `=` ใน Kotlin เป็น assignment operator คือยัดค่าเข้าตัวแปร ไม่ใช่การเปรียบเทียบ

```kotlin
val x = 5
println(x)
```

พิมพ์ `x` โดยไม่ใส่คำพูด จะได้ค่า `5` ถ้าใส่ `"x"` จะได้ตัวอักษร x

ทุกตัวแปรมี type จำนวนเต็มคือ `Int` เขียนเองได้ด้วย colon

```kotlin
val x: Int = 5
```

ถ้า IDE ทำให้ `Int` เป็นสีเทา แปลว่าซ้ำซ้อน ใส่หรือไม่ใส่ผลรันเหมือนกัน เพราะมี type inference คอมไพเลอร์ดูค่าที่ assign แล้วเดา type ให้ เลข `5` เป็นจำนวนเต็ม จึงเป็น `Int`

ทศนิยมมีสองแบบ

- `Float` กว้าง 32 bit ต้องเติม `f` ท้ายค่า เช่น `5.123f` ถ้าเขียนแค่ `5` IDE จะฟ้องว่า integer literal ไม่ตรงกับ `Float` พิมพ์ออกมาเป็น `5.0`
- `Double` กว้าง 64 bit แม่นกว่า เพราะมี bit ให้เก็บมากกว่า ประกาศด้วยจุดทศนิยมโดยไม่เติม `f` เช่น `5.123`

เขาแนะนำ `Float` เมื่อไม่ต้องการความแม่นยำสูง เพราะกิน memory น้อยกว่า และใช้ `Double` เมื่อต้องการความแม่น

`Boolean` มีได้แค่ `true` หรือ `false` ใช้ตัดสินใจว่าจะรันโค้ดบางก้อนหรือไม่

`String` คือข้อความ อาจเป็นย่อหน้า คำเดียว หรือตัวเลขที่ไม่ได้ตีความเป็นจำนวน ประกาศด้วยเครื่องหมายคำพูด

ห้า type นี้เขาเรียกว่า primitive data type ของบทนี้ และบอกว่ายังมีอีกหนึ่งแบบที่จะมาทีหลัง

`val` มาจาก value เปลี่ยนค่าทีหลังไม่ได้ `var` มาจาก variable เปลี่ยนได้

```kotlin
var x = "hello world"
x = "bye-bye world"
println(x)
```

ได้ `bye-bye world` ซึ่งเป็นค่าล่าสุดตอนเรียก `println` IDE อาจทำให้บรรทัดแรกเป็นสีเทา เพราะถูกทับทันที แต่โค้ดยังถูกต้อง

## Example

```kotlin
fun main() {
    val whole: Int = 5
    val decimal = 5.123
    val rough = 5.123f
    val ok = true
    var message = "hello world"
    message = "bye-bye world"
    println(whole)
    println(decimal)
    println(rough)
    println(ok)
    println(message)
}
```

```java
public class Main {
    public static void main(String[] args) {
        final int whole = 5;
        final double decimal = 5.123;
        final float rough = 5.123f;
        final boolean ok = true;
        String message = "hello world";
        message = "bye-bye world";
        System.out.println(whole);
        System.out.println(decimal);
        System.out.println(rough);
        System.out.println(ok);
        System.out.println(message);
    }
}
```

Kotlin ใส่ `val` หรือ `var` นำหน้า แล้วค่อยระบุ type หลัง colon ส่วน Java ใส่ type นำหน้า และใช้ `final` เมื่อไม่ต้องการให้เปลี่ยนค่า
