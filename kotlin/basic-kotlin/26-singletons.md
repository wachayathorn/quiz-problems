# Singletons

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:44:19–2:49:29

## Concept

Singleton คือ instance ของ class ที่มีได้ก้อนเดียวตลอดการรันของโปรแกรม สี่เหลี่ยมตัวเดียวใน `main` ยังไม่นับ เพราะไม่มีอะไรกันไม่ให้สร้างเพิ่ม สร้างของที่มีได้ก้อนเดียวด้วย `object`

`object` เหมือน class แต่ไม่มี constructor เพราะ constructor มีไว้รวมค่าที่ทำให้แต่ละ instance ต่างกัน ถ้ามีได้ก้อนเดียว ก็ไม่มีอะไรให้ส่งเข้า constructor

เขาทำสี่เหลี่ยมจัตุรัสขนาดตายตัวเป็น `Shape` ด้านยาว 4 พื้นที่กับเส้นรอบรูปจึงเป็น `16f` ทั้งคู่ กด Command+I เพื่อให้ IDE เติม `override` อ้างถึงตัวนี้ได้เลยโดยไม่เรียก constructor

```kotlin
println(FixedSizeSquare.area)
printShapes(FixedSizeSquare)
```

`when` ใน `printShapes` จะขอสาขาของ object นี้เพิ่ม เขาพิมพ์ว่าเป็น fixed size square

มี `data object` ด้วย ข้อดีเหนือ `object` ธรรมดาไม่ใหญ่แล้ว เพราะก๊อปอีกก้อนไม่ได้ พิมพ์แล้วอ่านง่ายขึ้น และมีประโยชน์ขั้นสูงที่เขาไม่ลงในคลิป เขาบอกว่าในบางจังหวะที่โชคร้าย `object` ธรรมดาอาจมีสอง instance แล้วการเทียบจะเป็นแบบ reference เหมือน class ปกติ `data object` ไม่เป็นแบบนั้น ถ้าจะตามต่อให้ดู reflection ซึ่งเกินขอบเขต crash course นี้

เขาเองมองว่าสี่เหลี่ยมขนาดตายตัวไม่ใช่ตัวอย่างดี เพราะสร้าง `Rectangle` ที่กว้างเท่าสูงก็ได้ รูปทรงที่ต้องมีรัศมีหรือด้านเพื่อคำนวณพื้นที่ มักควรเป็น class ไม่ใช่ object ใช้ `object` เมื่อต้องสืบทอดหรือรับ interface แต่ไม่ต้องการ constructor และไม่มี state ที่ทำให้แต่ละก้อนต่างกัน

อีกงานที่ใช้บ่อยคือเป็นที่รวมฟังก์ชันที่เกี่ยวกับกัน เช่น `DateUtil` มี `formatDate` รับค่าวันที่แล้วคืนสตริงตามแพตเทิร์น เรียกเป็น `DateUtil.formatDate(...)` แล้ว autocomplete จะโชว์ฟังก์ชันวันที่ทั้งหมดจากชื่อ object ไม่ต้องจำชื่อฟังก์ชันลอย ๆ

## Example

```kotlin
sealed interface Shape

object FixedSizeSquare : Shape

object DateUtil {
    fun formatDate(date: java.time.LocalDate): String {
        return date.toString()
    }
}

fun main() {
    println(FixedSizeSquare)
    println(DateUtil.formatDate(java.time.LocalDate.of(2025, 1, 2)))
}
```

```java
sealed interface Shape permits FixedSizeSquare {}

enum FixedSizeSquare implements Shape {
    INSTANCE
}

class DateUtil {
    private DateUtil() {}

    static String formatDate(java.time.LocalDate date) {
        return date.toString();
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println(FixedSizeSquare.INSTANCE);
        System.out.println(DateUtil.formatDate(java.time.LocalDate.of(2025, 1, 2)));
    }
}
```

Kotlin ใช้คำว่า `object` แล้วอ้างชื่อนั้นเป็นก้อนเดียวที่มีอยู่ ส่วน Java ไม่มีคำนี้ ตัวอย่างจึงกันไม่ให้สร้าง `DateUtil` เพิ่มด้วย constructor แบบ `private` และใช้ `enum` ที่มีค่าเดียวเมื่อต้องมี instance จริง

💡 เพิ่มเติม: ตัวฟังก์ชัน `formatDate` ในคลิปถูกเปิดไว้เป็นโครง เขายังไม่เขียนรูปแบบวันที่ข้างใน ตัวอย่างนี้คืน `toString()` ของ `LocalDate` เพื่อให้เรียกดูได้
