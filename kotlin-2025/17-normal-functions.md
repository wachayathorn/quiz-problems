# Normal functions

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 1:31:10–1:42:03

## Concept

โปรแกรมทั้งก้อนรับ input แล้วให้ output ฟังก์ชันคือสิ่งเดียวกันในระดับที่เล็กลง ฟังก์ชันที่เจอมาแล้วมี `main`, `readln` และ `toIntOrNull` จุดร่วมคือมีวงเล็บท้าย บางครั้งว่าง บางครั้งมีค่าข้างใน

ถ้าโค้ดกลับสตริงถูกก็อปไปหลายจุด ไฟล์จะบวม และบั๊กต้องแก้ทุกจุด นิยามครั้งเดียวแล้วเรียกหนึ่งบรรทัดได้ทุกที่ที่ต้องการ อีกประโยชน์คืออ่านชื่อฟังก์ชันก็รู้ว่าก้อนนั้นทำอะไร โดยไม่ต้องไล่ loop ข้างใน

ฟังก์ชันประกาศนอก `main` เพราะ `main` เองก็เป็นฟังก์ชันที่ถูกรันตอนเปิดโปรแกรม รูปแบบใช้ `fun` เหมือนกัน

```kotlin
fun reverse(stringToReverse: String): String {
    val finalString = buildString {
        for (i in stringToReverse.lastIndex downTo 0) {
            append(stringToReverse[i])
        }
    }
    return finalString
}
```

พารามิเตอร์อยู่ในวงเล็บ ต้องเขียน type เอง เพราะยังไม่มีค่าให้ infer ชื่อเช่น `stringToReverse` คือตัวแปรที่ใช้ข้างในฟังก์ชัน ฟังก์ชันอยู่คนละ scope กับ `main` จึงไม่เห็นตัวแปร `input` ของ `main` เห็นเฉพาะสิ่งที่ถูกส่งเข้ามา หลายพารามิเตอร์คั่นด้วยจุลภาค

`return` ส่งค่ากลับจุดที่เรียก และจบการทำงานของฟังก์ชันนั้น type ที่คืนเขียนหลังวงเล็บด้วย `: String` ถ้าประกาศ type นี้แล้วไม่มี `return` โปรแกรมคอมไพล์ไม่ผ่าน เรียกแล้วไม่ใช้ค่าที่คืน จะไม่เห็นข้อความอะไรบน console จนกว่าจะเอาไป `println` หรือทำอย่างอื่น

```kotlin
val reversed = reverse(input)
println(reversed)
println(reverse(reversed))
```

`readln()` ก็คืน `String` แบบเดียวกัน คือข้อความที่ผู้ใช้พิมพ์ เอาสตริงที่กลับด้านแล้วไปเทียบกับต้นฉบับ ได้ตัวตรวจ palindrome เขาลองกับ `racecar` ซึ่งอ่านจากท้ายมาหน้าได้คำเดิม เก็บผล `reverse` ไว้ในตัวแปรแล้วใช้ซ้ำ ดีกว่าเรียกฟังก์ชันสองครั้งเพราะแต่ละครั้งวนสตริงใหม่

ค่าเริ่มต้นใส่ที่พารามิเตอร์ได้

```kotlin
fun reverse(stringToReverse: String = "hello world"): String {
```

ถ้าเรียก `reverse()` โดยไม่ส่งค่า จะใช้ `"hello world"`

Named argument เขียนชื่อพารามิเตอร์ตรงจุดเรียก เพื่อให้อ่านรู้ว่าค่านี้ไปบทบาทไหน เมื่อฟังก์ชันมีหลายพารามิเตอร์

```kotlin
reverse(stringToReverse = input)
```

พฤติกรรมไม่เปลี่ยน มีแค่ความอ่านง่าย เขาแนะนำให้ใช้เมื่อดูจากค่าอย่างเดียวแล้วไม่ชัด และบอกว่าใช้กับฟังก์ชันที่เขียนมาจาก Java ไม่ได้ เพราะเป็นฟีเจอร์ของ Kotlin

## Example

```kotlin
fun main() {
    println("enter a string")
    val input = readln()
    val reversed = reverse(stringToReverse = input)
    if (input == reversed) {
        println("that's a palindrome")
    }
    println(reversed)
}

fun reverse(stringToReverse: String = "hello world"): String {
    val finalString = buildString {
        for (i in stringToReverse.lastIndex downTo 0) {
            append(stringToReverse[i])
        }
    }
    return finalString
}
```

```java
public class Main {
    public static void main(String[] args) {
        String input = "racecar";
        String reversed = reverse(input);
        if (input.equals(reversed)) {
            System.out.println("that's a palindrome");
        }
        System.out.println(reversed);
    }

    static String reverse(String stringToReverse) {
        StringBuilder finalString = new StringBuilder();
        for (int i = stringToReverse.length() - 1; i >= 0; i--) {
            finalString.append(stringToReverse.charAt(i));
        }
        return finalString.toString();
    }
}
```

Kotlin ประกาศด้วย `fun` วางไว้นอก class ได้ มีค่าเริ่มต้นของพารามิเตอร์ และเรียกแบบระบุชื่อได้ ส่วน Java ในตัวอย่างนี้เป็น `static` method ใน class ไม่มีค่าเริ่มต้นให้พารามิเตอร์ และเทียบ `String` ด้วย `equals`

💡 เพิ่มเติม: `input == reversed` ใน Kotlin เทียบข้อความให้แล้ว ใน Java `==` ระหว่าง `String` เทียบว่าเป็นวัตถุก้อนเดียวกันหรือไม่
