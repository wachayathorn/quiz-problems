# Abstract classes

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:26:14–2:35:24

## Concept

Abstract class กับ interface ถูกสับสนบ่อย เขาลองเปลี่ยน `Shape` เป็น class

```kotlin
abstract class Shape {
    abstract val area: Float
    abstract val circumference: Float
}
```

`abstract` บนค่าหรือฟังก์ชันแปลว่าตรงนี้ยังไม่ลงมือทำ แต่ class ที่บอกว่าเป็น `Shape` ต้อง `override` ให้ ลูกเรียก constructor ของแม่ด้วย `()` แม้ไม่มี argument เหมือนเรียกฟังก์ชันที่ไม่มีพารามิเตอร์ รันแล้วรวมพื้นที่ได้เหมือนเดิม

ความต่างที่เขาเน้นมีสามจุด

Abstract class เป็นพิมพ์เขียวจริง ไม่ได้มีไว้ประกาศสัญญาอย่างเดียว และสร้าง instance ตรง ๆ ไม่ได้ `Shape()` จะพัง สร้างได้แค่ class ลูก เขาเรียกความสัมพันธ์นี้ว่า inheritance ลูกได้สิ่งที่แม่กำหนดให้สืบทอด ในตัวอย่างคือ `area` กับ `circumference`

Abstract class มีพฤติกรรมธรรมดาและ state ได้ State ในความหมายที่เขาใช้คือค่าที่เปลี่ยนตามเวลา ซึ่งก็คือ `var`

```kotlin
var counter = 0
fun increment() {
    counter++
}
```

`Rectangle` ไม่ได้เขียน `increment` เอง แต่เรียก `rect1.increment()` ได้ เพราะสืบทอดมา สี่เหลี่ยมคนละรูปมี `counter` คนละตัว ไม่ปนกัน Interface ใส่ค่าเริ่มต้นแบบ `counter` ที่เปลี่ยนในตัว interface ไม่ได้ ให้ implementation เริ่มต้นของฟังก์ชันได้ แต่ตัวแปรที่มีค่าและถูกแก้ในตัว interface ไม่มี เขาบอกว่าส่วนใหญ่ควรใช้ interface ถ้าต้องการทั้ง state และการสืบทอด ค่อยใช้ abstract class

เอาคำว่า `abstract` ออกแล้วต้องลงมือทำให้ค่าในแม่เอง และ class ปกติกับ data class สืบทอดไม่ได้ ถ้าอยากให้สืบทอดได้โดยยังสร้าง instance ของแม่ได้ ให้ใส่ `open`

```kotlin
open class Shape {
    open val area: Float = 0f
    open val circumference: Float = 0f
}
```

`open` ไม่ได้ห้ามสร้าง `Shape()` ลูกไม่ override ก็ได้ และจะได้ค่าของแม่ ซึ่งในตัวอย่างคือ 0 เขามองว่าไม่เข้ากับรูปทรง เพราะ shape ลอย ๆ ไม่มีสูตรพื้นที่ เขาเลยย้อนกลับเป็น interface และเติมฟังก์ชันใน interface ให้ดูสั้น ๆ ว่าทำได้ เช่น `helloWorld` ที่รับ counter แล้วลูกทุกตัวต้องลงมือทำ จากนั้นเขาลบฟังก์ชันนี้ออก

เช็กชนิดจริงของ shape ใช้ `when` ได้สะดวก

```kotlin
fun printShapes(vararg shapes: Shape) {
    for (shape in shapes) {
        when (shape) {
            is Circle -> println("yo, that's a circle")
            is Rectangle -> println("that's a rectangle")
        }
    }
}
```

`is` ใช้ใน `if` ได้เหมือนกันเมื่ออยากเช็กชนิดเดียว ถ้าตัวแปรถูกประกาศเป็น `Circle` อยู่แล้ว ไปเทียบ `is Rectangle` จะฟ้อง incompatible types เพราะคอมไพเลอร์รู้ว่าเป็นวงกลมแน่นอน การเช็กจึงมีประโยชน์เมื่อรับมาเป็น `Shape` และยังไม่รู้ว่าเป็นรูปไหน

ถ้า `when` ถูกใช้เป็น expression เพื่อเอาสตริงออกมา มันต้อง exhaustive คือทุกกรณีที่เป็นไปได้มีสาขา เขาเจอว่าต้องมี `else` และสาขาต้องให้ค่า ไม่ใช่แค่ `println` ถ้าไม่ครบ คอมไพเลอร์ไม่รู้ว่าจะ assign อะไร ใช้ `when` แค่รันคำสั่งแล้วไม่เอาค่า ไม่ครบได้ เพราะกรณีที่ไม่มีสาขาคือไม่ทำอะไร

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

fun printShapes(vararg shapes: Shape) {
    for (shape in shapes) {
        when (shape) {
            is Circle -> println("yo, that's a circle")
            is Rectangle -> println("that's a rectangle")
            else -> println("some other shape")
        }
    }
}

fun main() {
    printShapes(Rectangle(5f, 7f), Circle(5f))
}
```

```java
interface Shape {
    float area();
    float circumference();
}

record Rectangle(float width, float height) implements Shape {
    public float area() { return width * height; }
    public float circumference() { return 2 * width + 2 * height; }
}

record Circle(float radius) implements Shape {
    public float area() { return radius * radius * (float) Math.PI; }
    public float circumference() { return 2 * radius * (float) Math.PI; }
}

public class Main {
    static void printShapes(Shape... shapes) {
        for (Shape shape : shapes) {
            if (shape instanceof Circle) {
                System.out.println("yo, that's a circle");
            } else if (shape instanceof Rectangle) {
                System.out.println("that's a rectangle");
            } else {
                System.out.println("some other shape");
            }
        }
    }

    public static void main(String[] args) {
        printShapes(new Rectangle(5f, 7f), new Circle(5f));
    }
}
```

Kotlin ใช้ `abstract class` เมื่อแม่ต้องมี state และใช้ `is` ใน `when` เพื่อแยกชนิด ส่วน Java ใช้ `abstract class` กับ field ได้เช่นกัน และแยกชนิดด้วย `instanceof` เขาเลือก interface สำหรับรูปทรงเพราะไม่มี state

💡 เพิ่มเติม: class ใน Kotlin ปิดการสืบทอดไว้ก่อน ต้องใส่ `open` หรือ `abstract` เอง class ใน Java สืบทอดได้เลยจนกว่าจะใส่ `final`
