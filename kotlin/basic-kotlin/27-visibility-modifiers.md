# Visibility modifiers

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:49:29–2:53:06

## Concept

สมาชิกของ class ที่ไม่เขียนตัวกำหนดการมองเห็น จะเป็น `public` อ่านได้จากข้างนอก class เช่น `rect1.diagonal` จาก `main` ได้ ก็เพราะค่านี้เป็น public

`private` จำกัดให้อ่านได้แค่ในตัว class นั้น พอใส่ให้ `diagonal` แล้วพิมพ์ `rect1.` จะไม่เห็นชื่อนี้โผล่ข้างนอก เขาแนะนำว่าอย่าเปิดกว้างกว่าที่ใช้ ถ้า field มีไว้ใน class อย่างเดียว ให้เป็น `private` แล้วเปิดเฉพาะสิ่งที่ฟังก์ชันหรือไฟล์อื่นต้องใช้

`protected` ใช้เมื่อมี class แม่ เป็นได้ทั้ง `open class` และ `abstract class` ค่าที่ประกาศ `protected` ในแม่ อ่านได้ในตัวแม่และในลูกทุกตัว แต่อ่านจากข้างนอกไม่ได้ สร้าง shape แล้วจะแตะ `counter` จากนอกสายนี้ไม่ได้ แต่ในตัว class ลูกอ้าง `counter` ได้

`internal` มองเห็นได้ในโมดูลปัจจุบัน เขาบอกว่ามือใหม่ยังไม่ต้องใช้ มีประโยชน์เมื่อแจกโค้ดเป็นไลบรารีให้คนอื่น ถ้าทำโปรเจกต์เดียวและไม่ส่งออกไปโมดูลอื่น ก็ยังไม่จำเป็น เขาเลยลบตัวอย่างนี้ออก

ก่อนจบบท class เขาบอกว่ายังมีรายละเอียดที่คลิปไม่ได้ทำให้ดู สองอย่างที่เอ่ยชื่อคือ operator function ซึ่งกำหนดได้ว่า `rect1 + rect2` หมายความว่าอะไร และ constructor ได้มากกว่าหนึ่งตัว สิ่งที่ทำให้ดูแล้วคือฐานที่ใช้เริ่มเขียน Kotlin จริงได้

## Example

```kotlin
abstract class Shape {
    protected var counter = 0
}

data class Rectangle(val width: Float, val height: Float) : Shape() {
    private val diagonal = kotlin.math.sqrt(width * width + height * height)

    fun grow() {
        counter++
    }
}

fun main() {
    val rect1 = Rectangle(5f, 7f)
    println(rect1.width)
    rect1.grow()
}
```

```java
abstract class Shape {
    protected int counter = 0;
}

class Rectangle extends Shape {
    public final float width;
    public final float height;
    private final float diagonal;

    Rectangle(float width, float height) {
        this.width = width;
        this.height = height;
        this.diagonal = (float) Math.sqrt(width * width + height * height);
    }

    public void grow() {
        counter++;
    }
}

public class Main {
    public static void main(String[] args) {
        Rectangle rect1 = new Rectangle(5f, 7f);
        System.out.println(rect1.width);
        rect1.grow();
    }
}
```

`private` กับ `protected` จำกัดขอบเขตแบบเดียวกัน Kotlin ไม่เขียนตัวกำหนดแล้วสมาชิกเป็น `public` ส่วน Java ไม่เขียนแล้วสมาชิกใน class ถูกมองเห็นได้แค่ใน package เดียวกัน

💡 เพิ่มเติม: `internal` ของ Kotlin จำกัดที่โมดูล Java ไม่มีคำนี้ตรงกัน package-private จำกัดที่ package
