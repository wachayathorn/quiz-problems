# Lambda functions

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 1:52:25–2:07:33

## Concept

`buildString { }` เป็นฟังก์ชันที่มีปีกกาตามหลัง เพราะรับ lambda เป็นพารามิเตอร์ Lambda คือฟังก์ชันก้อนหนึ่งที่ส่งให้ฟังก์ชันอื่นได้ และเรียกซ้ำได้

`filter` ใช้เมื่อกฎการคัดเปลี่ยนตามจุดเรียก บางทีคัดเลขคู่ บางทีคัดเฉพาะตัวอักษร กฎนั้นไม่ถูกฝังตายใน `filter` แต่ให้ผู้เรียกส่งมา พารามิเตอร์นี้ในลายเซ็นชื่อ `predicate` type ประมาณ `Char -> Boolean` คือรับตัวอักษรหนึ่งตัวแล้วคืน `Boolean`

ถ้าพารามิเตอร์เป็น lambda วงเล็บของฟังก์ชันตัดออกได้ เหลือปีกกา ซึ่งคือตัว lambda ก้อนในปีกกาถูกเรียกทีละตัวอักษรของสตริง คืน `true` แปลว่าเก็บไว้ คืน `false` แปลว่าตัดออก พารามิเตอร์ตัวเดียวอ้างด้วย `it`

```kotlin
val lettersOnly = input.filter { it.isLetter() }
```

พิมพ์ `hello123` ได้ `hello` ใช้กับ array ของจำนวนได้เช่นกัน type ของ `it` จะเป็น `Int`

```kotlin
val favoriteNumbers = intArrayOf(1, 2, 3, 69)
val evenNumbers = favoriteNumbers.filter { it % 2 == 0 }
```

ได้แค่ `2`

เก็บ lambda ลงตัวแปรได้ แต่ต้องเขียน type เอง เพราะ infer ไม่รู้ว่าเข้าและออกเป็นอะไร

```kotlin
val lambda: (Char) -> Boolean = { it.isLetter() }
val lettersOnly = input.filter(lambda)
```

ก้อนปีกกาจึงเป็นค่าที่ส่งเข้า `filter` จริง ไม่ใช่ไวยากรณ์ลอย ๆ ถ้ามีพารามิเตอร์เดียวจะตั้งชื่อเองก็ได้ หลังชื่อใส่ลูกศรแล้วตัวถัดไปคือตัวฟังก์ชัน ตั้งชื่อแล้วใช้ `it` ไม่ได้

```kotlin
input.filter { currentCharacter -> currentCharacter.isLetter() }
```

`it` ใช้ได้เมื่อมีพารามิเตอร์ตัวเดียว ถ้ามีหลายตัว เช่น `Char` กับ `Int` ต้องตั้งชื่อเอง

ฟังก์ชันของเราเองรับ lambda ได้ เขาเขียน extension ชื่อ `myFilter` เพื่อไม่ชนของไลบรารี

```kotlin
fun String.myFilter(predicate: (Char) -> Boolean): String {
    return buildString {
        for (character in this@myFilter) {
            if (predicate(character)) {
                append(character)
            }
        }
    }
}
```

`predicate(character)` คือการเรียก lambda ที่ผู้ใช้ส่งมา คืน `true` จึง `append` สร้างสตริงใหม่เฉพาะตัวที่ผ่าน แล้วคืนสตริงนั้น

`map` เป็นอีกตัวที่ใช้บ่อย เดินทุกสมาชิกแล้วแปลงเป็นค่าใหม่

```kotlin
favoriteNumbers.map { it * it }
```

ได้ลิสต์ของเลขยกกำลังสอง เขาบอกว่าชื่อตัวแปร `evenNumbers` หลัง `map` ไม่ตรงความหมายแล้ว แต่ใช้โชว์การแปลงต่อ

Extension lambda ใส่ type ทางซ้ายของจุดแทนการรับเป็นพารามิเตอร์ในวงเล็บ

```kotlin
predicate: Char.() -> Boolean
```

เรียกเป็น `character.predicate()` และข้างใน lambda อ้างตัวอักษรปัจจุบันได้โดยไม่เขียน `it` เพราะ `this` คือตัวอักษรนั้น เขียน `this.isLetter()` หรือ `isLetter()` ก็ได้

`buildString` ใช้แบบนี้อยู่แล้ว lambda ของมันคืน `Unit` ซึ่งเขาแปลว่าไม่มีค่าส่งกลับ ฟังก์ชันทำอะไรบางอย่างแต่ไม่มี output ชัด และถูกเรียกบน `StringBuilder` จึงมี `append` ให้ใช้ `this` ข้างในจึงเป็น `StringBuilder` ไม่ใช่สตริงต้นทางของ `reverse` นั่นคือเหตุที่บทก่อนต้องเขียน `this@reverse` เมื่ออยู่ใน scope ของ `buildString` แต่ต้องการ `this` ของฟังก์ชันนอก

เขาปิดบทฟังก์ชันว่าฟังก์ชันปกติ, extension และ lambda ครอบคลุมประมาณ 90% ของที่ใช้ ชนิดพิเศษที่เหลือจะเจอเองตอนเขียนของจริง

## Example

```kotlin
fun main() {
    val input = readln()
    val lettersOnly = input.myFilter { it.isLetter() }
    println(lettersOnly)

    val favoriteNumbers = intArrayOf(1, 2, 3, 69)
    println(favoriteNumbers.filter { it % 2 == 0 })
    println(favoriteNumbers.map { it * it })
}

fun String.myFilter(predicate: (Char) -> Boolean): String {
    return buildString {
        for (character in this@myFilter) {
            if (predicate(character)) {
                append(character)
            }
        }
    }
}
```

```java
import java.util.Arrays;
import java.util.function.Predicate;

public class Main {
    public static void main(String[] args) {
        String input = "hello123";
        System.out.println(myFilter(input, Character::isLetter));

        int[] favoriteNumbers = {1, 2, 3, 69};
        System.out.println(Arrays.stream(favoriteNumbers).filter(n -> n % 2 == 0).boxed().toList());
        System.out.println(Arrays.stream(favoriteNumbers).map(n -> n * n).boxed().toList());
    }

    static String myFilter(String value, Predicate<Character> predicate) {
        StringBuilder result = new StringBuilder();
        for (char character : value.toCharArray()) {
            if (predicate.test(character)) {
                result.append(character);
            }
        }
        return result.toString();
    }
}
```

Kotlin ส่งก้อนโค้ดด้วยปีกกา และอ้างพารามิเตอร์ตัวเดียวด้วย `it` ส่วน Java ส่งด้วย `->` หรือ method reference และเรียกผ่าน `Predicate.test`
