# if conditions

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 47:27–55:48

## Concept

`if` รันก้อนโค้ดเมื่อเงื่อนไขในวงเล็บเป็น `true` เท่านั้น `!=` คือไม่เท่ากับ ซึ่งเขาบอกว่าหลุดไปจากบท comparison

```kotlin
if (inputAsInteger != null) {
    val isEven = inputAsInteger % 2 == 0
}
```

ข้างในบล็อกนี้คอมไพเลอร์รู้ว่าค่าไม่เป็น `null` จึงใช้ modulo ได้ นี่เป็นส่วนของ null safety ถ้าพิมพ์ `hello` เงื่อนไขเป็นเท็จ บล็อกไม่ถูกรัน และไม่มีข้อความออก

`else` รันเมื่อเงื่อนไขเป็นเท็จ เช่น บอกให้ใส่ตัวเลขที่ถูกต้อง

ซ้อน `if` ได้ และอย่าเขียน `if (isEven == true)` เพราะ `isEven` เป็น `Boolean` อยู่แล้ว การเทียบ `true == true` ฟุ่มเฟือย ส่งตัวแปรเข้าไปตรง ๆ

`!` ด้านหน้ากลับค่า `Boolean` ใช้เมื่ออยากพิมพ์เฉพาะกรณีเลขคี่ โดยไม่ต้องเปิด `if` ว่างไว้เพื่อไปหา `else` ถ้ามี `else` อยู่แล้ว เขาแนะนำให้เงื่อนไขฝั่ง `if` ไม่ถูกกลับค่า เพราะเงื่อนไขจริงในงานมักยาวและผูกด้วยตัวดำเนินการตรรกะ อ่านแบบไม่กลับค่าง่ายกว่า

`if` ใน Kotlin เป็น expression คือ assign ผลให้ตัวแปรได้ ทั้งสองสาขาต้องให้ค่าชนิดเดียวกัน ถ้าไม่มี `else` คอมไพเลอร์ไม่รู้ว่าเงื่อนไขเท็จแล้วจะใส่อะไร ถ้าสาขาเท็จควรเป็น `null` ให้เขียน `else null` ได้ type จะกลายเป็น `String?` ส่ง `null` เข้า `println` ได้ มันจะพิมพ์คำว่า null

`else if` แทรกได้หลายก้อน ลำดับตรวจจากบนลงล่าง ตัวอย่างของเขา

- หารด้วย 2 ลงตัว พิมพ์ว่าเลขคู่
- น้อยกว่า 10 พิมพ์ว่าคี่และน้อยกว่า 10
- นอกนั้น พิมพ์ว่าคี่และอย่างน้อย 11

ใส่ `5` ได้สาขาที่สอง ใส่ `15` ได้สาขาที่สาม ใส่เลขคู่ได้สาขาแรก

## Example

```kotlin
fun main() {
    println("please enter a number")
    val inputAsInteger = readln().toIntOrNull()
    val output = if (inputAsInteger == null) {
        "dude, enter a valid number"
    } else if (inputAsInteger % 2 == 0) {
        "the number is even"
    } else if (inputAsInteger < 10) {
        "the number is odd and less than 10"
    } else {
        "the number is odd and at least 11"
    }
    println(output)
}
```

```java
public class Main {
    public static void main(String[] args) {
        Integer inputAsInteger = 15;
        final String output;
        if (inputAsInteger == null) {
            output = "dude, enter a valid number";
        } else if (inputAsInteger % 2 == 0) {
            output = "the number is even";
        } else if (inputAsInteger < 10) {
            output = "the number is odd and less than 10";
        } else {
            output = "the number is odd and at least 11";
        }
        System.out.println(output);
    }
}
```

Kotlin ให้ `if` / `else` ทั้งก้อนเป็นค่าที่ assign ได้ ส่วน Java ใช้ `if` เป็นคำสั่ง แล้ว assign ในแต่ละสาขาเอง กรณีสองสาขาสั้น ๆ ของ Java ใช้ ternary `เงื่อนไข ? ค่าเมื่อจริง : ค่าเมื่อเท็จ`
