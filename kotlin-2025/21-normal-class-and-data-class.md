# Normal class and data class

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:07:33–2:18:28

## Concept

Kotlin ใช้ object-oriented programming แบบเดียวกับ Java และภาษาอื่นอีกหลายภาษา Class มีสองงาน คือมัดข้อมูลที่เกี่ยวกับกัน และให้พฤติกรรมที่ทำงานบนก้อนนั้น อีกมุมหนึ่ง class คือพิมพ์เขียว

```kotlin
class Rectangle(val width: Float, val height: Float)
```

`Rectangle(5f, 7f)` ไม่ใช่การใส่เลขหรือสตริง แต่เป็นการสร้าง instance คือรูปสี่เหลี่ยมรูปหนึ่งที่มีความกว้างและความสูงของตัวเอง สร้างได้กี่ instance ก็ได้ ทุกตัวมี `width` กับ `height` ตามพิมพ์เขียวเดียวกัน อ่านค่าด้วย `rect1.width`

ค่าในวงเล็บของ class คือ constructor สิ่งที่ต้องมีจึงจะสร้าง instance ได้ สี่เหลี่ยมที่ไม่มีสูงไม่มีจริง

ในตัว class ใส่พฤติกรรมเพิ่มได้ เขาคำนวณเส้นทแยงด้วยทฤษฎีพีทาโกรัส `sqrt(width * width + height * height)` ฟังก์ชัน `sqrt` ใช้เลขทศนิยม เขาเลยเปลี่ยน `width` และ `height` เป็น `Float` แต่ยังส่งจำนวนเต็มได้โดยเติม `f` เช่น `5f` กับ `7f` พื้นที่คือ `width * height` ซึ่งตัวอย่างนี้ได้ `35`

instance ส่งเข้าฟังก์ชันอื่นได้ เขาเขียน `maxArea` ที่รับสี่เหลี่ยมสองรูป เอา `area` ของแต่ละรูป แล้วคืนค่าที่ใหญ่กว่าด้วย `maxOf` จากนั้นลบฟังก์ชันนี้ออกเพราะต้องการไปที่การเทียบ

สร้างอีกรูปด้วยกว้าง `5f` สูง `7f` แล้วพิมพ์ `rect1 == rect2` ได้ `false` ทั้งที่ขนาดเท่ากัน เพราะ class ปกติเทียบแบบ reference คือเป็น instance เดียวกันหรือไม่ ไม่ได้เทียบข้อมูลข้างใน เขายกตัวอย่างรถแดงสองคัน คนละคันแม้สีเดียวกัน ถ้าเขียน `rect2 = rect1` ผลเป็น `true` เพราะทั้งสองชื่อชี้ instance เดียว

เมื่อ class เอาไว้ถือข้อมูลเป็นหลัก และฟังก์ชันแค่คำนวณจากข้อมูลนั้น ให้ใช้ `data class` การเทียบจะดู field คือ `width` เท่ากันและ `height` เท่ากันหรือไม่ เปลี่ยนสูงเป็น `8f` แล้วการเทียบเป็นเท็จ

`println(rect1)` ของ data class พิมพ์ค่าให้อ่านออก ถ้าเป็น class ปกติจะได้ที่อยู่ในหน่วยความจำ โดยไม่บอกข้อมูลข้างใน

`copy` มีเฉพาะ data class สร้าง instance ใหม่ที่ก๊อปค่ามา แล้วเปลี่ยนเฉพาะ field ที่ระบุ

```kotlin
val rect2 = rect1.copy(height = 10f)
```

กว้างยังเป็น `5f` จากต้นทาง สูงกลายเป็น `10f` เอา `data` ออกจะเรียก `copy` ไม่ได้

เขาสร้าง `Circle` ต่อ ข้อมูลขั้นต่ำที่ทำให้วงกลมหนึ่งวงต่างจากอีกวงคือ `radius` ไม่มีความกว้างความสูง พื้นที่คือ `radius * radius * PI` และเส้นผ่านศูนย์กลางคือ `2 * radius` ใน Kotlin มีค่าคงที่ `PI`

## Example

```kotlin
fun main() {
    val rect1 = Rectangle(width = 5f, height = 7f)
    val rect2 = rect1.copy(height = 10f)
    println("${rect1.width} ${rect1.height}")
    println(rect1.area)
    println(rect1.diagonal)
    println(rect1 == Rectangle(5f, 7f))
    println(rect1)
    println(rect2)
}

data class Rectangle(val width: Float, val height: Float) {
    val diagonal = kotlin.math.sqrt(width * width + height * height)
    val area = width * height
}
```

```java
public class Main {
    public static void main(String[] args) {
        Rectangle rect1 = new Rectangle(5f, 7f);
        Rectangle rect2 = new Rectangle(rect1.width(), 10f);
        System.out.println(rect1.width() + " " + rect1.height());
        System.out.println(rect1.area());
        System.out.println(rect1.diagonal());
        System.out.println(rect1.equals(new Rectangle(5f, 7f)));
        System.out.println(rect1);
        System.out.println(rect2);
    }
}

record Rectangle(float width, float height) {
    float diagonal() {
        return (float) Math.sqrt(width * width + height * height);
    }

    float area() {
        return width * height;
    }
}
```

Kotlin สร้าง object โดยไม่เขียน `new` และ `data class` ได้ `equals`, ข้อความตอนพิมพ์ และ `copy` ให้ ส่วน Java 16 ขึ้นไปใช้ `record` เพื่อได้ `equals` กับข้อความตอนพิมพ์ในทางเดียวกัน การก๊อปแล้วเปลี่ยนบาง field ในตัวอย่างนี้สร้าง `record` ใหม่เอง
