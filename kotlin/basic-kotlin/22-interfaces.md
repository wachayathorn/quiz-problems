# Interfaces

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:18:28–2:26:14

## Concept

ฟังก์ชันรวมพื้นที่ของหลายรูปส่งลิสต์ของรูปเข้ามาได้ หรือใช้ `vararg` ซึ่งเป็น argument จำนวนไม่고정แบบที่เขาใช้ตอนสร้าง array ถ้า type เป็น `Rectangle` จะรวมวงกลมไม่ได้ ถ้าเป็น `Circle` จะรวมสี่เหลี่ยมไม่ได้

Interface คือสัญญาว่า class ที่ขอเป็น shape ต้องมีอะไรบ้าง

```kotlin
interface Shape {
    val area: Float
    val circumference: Float
}
```

เขาวาง `area` กับ `circumference` เพราะทุกรูปมี สิ่งที่ไม่มีทุกรูปเช่นเส้นทแยงหรือเส้นผ่านศูนย์กลางไม่ใส่ในสัญญา ตัว interface ยังไม่บอกวิธีคำนวณ เพราะแต่ละรูปคนละสูตร ถ้า class บอกว่าเป็น `Shape` แต่ไม่มีสองค่านี้ แปลว่าผิดสัญญา

ประกาศว่ารับสัญญาหลัง constructor ด้วย colon

```kotlin
data class Rectangle(
    val width: Float,
    val height: Float
) : Shape
```

IDE จะฟ้องว่ายังไม่มี `circumference` ใน IntelliJ กด Command+I แล้วเลือก field ที่ต้องทำให้ จะได้โครงให้ `override` แปลว่าค่านี้ทับของระดับบน ซึ่งตรงนี้คือ `Shape`

ทุก field ใน class มี getter กับ setter Getter คือฟังก์ชันที่ถูกเรียกตอนอ่านค่า เช่นตอน `println` ผู้เขียนไม่ได้เรียกเอง

```kotlin
override val circumference: Float
    get() = 2 * width + 2 * height
```

เขียน `val circumference = 2 * width + 2 * height` ก็ได้ แต่แบบนี้คำนวณครั้งเดียวตอนสร้าง object แบบ `get()` คำนวณทุกครั้งที่ถูกอ่าน เขาเลือกแบบคำนวณครั้งเดียวเพราะกว้างกับสูงของตัวอย่างไม่เปลี่ยน ถ้าค่าต้นทางเปลี่ยนระหว่างการอ่านหลายครั้ง ค่อยใช้ getter

Setter ถูกเรียกตอน assign ค่าใหม่ และไม่มีใน `val` เพราะ `val` เปลี่ยนค่าไม่ได้

`area` ที่มีอยู่แล้วต้องเติม `override` ด้วย ไม่งั้นยังไม่ถือว่าทำตามสัญญา

วงกลมคำนวณเส้นรอบรูปเป็น `2 * radius * PI` แล้วเจอว่า type เป็น `Double` ไม่ใช่ `Float` เพราะ `PI` เป็น `Double` เพื่อเก็บทศนิยมให้ละเอียด เขาแปลงด้วย `PI.toFloat()`

เมื่อทั้งสอง class เป็น `Shape` ฟังก์ชันรับ `vararg shapes: Shape` ได้ ไม่ต้องสนว่าเป็นสี่เหลี่ยมหรือวงกลม เขาใช้ `sumOf` ซึ่งรับ lambda แล้วบวกค่าที่ lambda คืนของแต่ละตัว `area` เป็น `Float` แต่ overload ของ `sumOf` ที่เขาเจอรับ `Double` จึงแปลงด้วย `toDouble()` แล้วให้ฟังก์ชันคืน `Double`

```kotlin
fun sumAreas(vararg shapes: Shape): Double {
    return shapes.sumOf { shape -> shape.area.toDouble() }
}
```

ส่ง `rect1` กับวงกลมเข้าไปได้ เพราะฟังก์ชันรู้แค่ว่าทุก shape มี `area`

เขาส่งท้ายว่าของที่คล้าย interface คือ abstract class เก็บไว้บทถัดไป

## Example

```kotlin
interface Shape {
    val area: Float
    val circumference: Float
}

data class Rectangle(val width: Float, val height: Float) : Shape {
    override val area = width * height
    override val circumference = 2 * width + 2 * height
}

data class Circle(val radius: Float) : Shape {
    override val area = radius * radius * kotlin.math.PI.toFloat()
    override val circumference = 2 * radius * kotlin.math.PI.toFloat()
}

fun sumAreas(vararg shapes: Shape): Double {
    return shapes.sumOf { shape -> shape.area.toDouble() }
}

fun main() {
    val rect1 = Rectangle(5f, 7f)
    val circle = Circle(5f)
    println(sumAreas(rect1, circle))
}
```

```java
interface Shape {
    float area();
    float circumference();
}

record Rectangle(float width, float height) implements Shape {
    public float area() {
        return width * height;
    }

    public float circumference() {
        return 2 * width + 2 * height;
    }
}

record Circle(float radius) implements Shape {
    public float area() {
        return radius * radius * (float) Math.PI;
    }

    public float circumference() {
        return 2 * radius * (float) Math.PI;
    }
}

public class Main {
    static double sumAreas(Shape... shapes) {
        double sum = 0;
        for (Shape shape : shapes) {
            sum += shape.area();
        }
        return sum;
    }

    public static void main(String[] args) {
        System.out.println(sumAreas(new Rectangle(5f, 7f), new Circle(5f)));
    }
}
```

ทั้งคู่ให้ class ผูกกับสัญญา แล้วฟังก์ชันรับตัวสัญญาได้โดยไม่สน class ลูก Kotlin ใช้ `val` ใน interface และ `vararg` ส่วน Java ใช้ method ใน interface และ `Shape...`
