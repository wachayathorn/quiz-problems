# Sealed interface and sealed class

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:35:24–2:37:34

## Concept

ตอน `when` ต้องคืนค่า คอมไพเลอร์บังคับให้มี `else` เพราะยังไม่รู้ว่า `Shape` มีกี่แบบ เขาเปลี่ยน interface เป็น `sealed interface` แล้ว `else` ไม่จำเป็นอีกต่อไป

เหตุผลคือในโปรเจกต์นี้ shape มีได้แค่ `Circle` กับ `Rectangle` Sealed interface ถูก implement ได้เฉพาะในโมดูลปัจจุบัน เขาอธิบายโมดูลแบบย่อว่าเป็นขอบเขตโค้ดของเรา คนอื่นที่เอาโค้ดเราไปใช้เป็นไลบรารีจะเพิ่ม shape เองไม่ได้ ไลบรารีในความหมายของเขาคือชุดฟังก์ชันสำเร็จที่ไม่ได้เขียนในโปรเจกต์นี้ แต่หยิบมาใช้ เหมือนฟังก์ชันแปลง list ของ standard library คอมไพเลอร์จึงรู้ครบว่ามีสาขาอะไรบ้าง `when` ที่คืนค่าเลยไม่ต้องมี `else`

`sealed class` ใช้เมื่อลูกทุกตัวต้องรับ constructor จากแม่ `Shape` ไม่เข้าเคสนี้ เพราะ `area` กับ `circumference` มาจากสูตรในลูก ไม่มีค่าที่ทุกรูปต้องรับจากข้างนอก ถ้าไม่มี constructor ให้ใช้ sealed interface ถ้าลูกทุกตัวต้องรับค่าจากแม่ ให้ใช้ sealed class

## Example

```kotlin
sealed interface Shape {
    val area: Float
}

data class Rectangle(val width: Float, val height: Float) : Shape {
    override val area = width * height
}

data class Circle(val radius: Float) : Shape {
    override val area = radius * radius * kotlin.math.PI.toFloat()
}

fun label(shape: Shape): String {
    return when (shape) {
        is Circle -> "yo, that's a circle"
        is Rectangle -> "that's a rectangle"
    }
}

fun main() {
    println(label(Rectangle(5f, 7f)))
    println(label(Circle(5f)))
}
```

```java
sealed interface Shape permits Rectangle, Circle {
    float area();
}

record Rectangle(float width, float height) implements Shape {
    public float area() {
        return width * height;
    }
}

record Circle(float radius) implements Shape {
    public float area() {
        return radius * radius * (float) Math.PI;
    }
}

public class Main {
    static String label(Shape shape) {
        return switch (shape) {
            case Circle circle -> "yo, that's a circle";
            case Rectangle rectangle -> "that's a rectangle";
        };
    }

    public static void main(String[] args) {
        System.out.println(label(new Rectangle(5f, 7f)));
        System.out.println(label(new Circle(5f)));
    }
}
```

Kotlin ประกาศ `sealed interface` แล้วลูกในโมดูลเดียวกันคือชุดที่อนุญาต ส่วน Java ต้องเขียนรายชื่อใน `permits` ทั้งคู่ทำให้ `switch` / `when` ที่คืนค่าครบทุกแบบโดยไม่มีสาขาอื่น
