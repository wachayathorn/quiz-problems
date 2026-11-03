# when expression

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 55:48–59:48

## Concept

`else if` ยาวแล้วอ่านยาก เพราะต้องกระโดดเข้าวงเล็บทีละก้อน `when` เป็นแนวคิดเดียวกับ `if` แต่เป็น syntactic sugar ให้อ่านเงื่อนไขหลายทางในบรรทัดเดียว แต่ละสาขาใช้ `->`

แบบแรกใส่เงื่อนไขเต็มในแต่ละสาขา ใช้เมื่อแต่ละข้อคนละแบบ เช่น modulo กับน้อยกว่า 10 สาขา `else` รับกรณีที่เหลือ

```kotlin
val output = when {
    inputAsInteger % 2 == 0 -> "the number is even"
    inputAsInteger < 10 -> "the number is odd and less than 10"
    else -> "the number is odd and at least 11"
}
```

ถ้าสาขาต้องทำหลายขั้น ไม่ใช่แค่ให้ค่าเดียว ให้เปิดปีกกาหลัง `->`

แบบที่สองส่งค่าเข้าวงเล็บของ `when` แล้วเทียบกับค่านอลูกศรตรง ๆ ใช้เมื่อทุกสาขาดูตัวเดียวกัน

```kotlin
val output = when (inputAsInteger) {
    null -> "dude, enter a valid number"
    3 -> "the number is 3"
    5 -> "the number is 5"
    in 10..20 -> "the number is between 10 and 20"
    else -> "something else"
}
```

`in` เช็กว่าอยู่ในช่วง `10..20` สาขาถูกตรวจจากบนลงล่าง สาขาแรกที่จริงจะถูกใช้ ถ้าเปลี่ยนช่วงเป็นตั้งแต่ 4 แล้วใส่ `5` จะได้ข้อความว่าเป็น 5 ทั้งที่ 5 ก็อยู่ในช่วงนั้นด้วย เพราะสาขา `5` อยู่ก่อน

`when` ไม่จำเป็นต้อง assign ให้ตัวแปร ใช้รันโค้ดในสาขาก็ได้ แต่ถ้าต้องได้คนละค่าตามหลายเงื่อนไข การ assign จาก `when` อ่านง่ายกว่าโซ่ `else if`

## Example

```kotlin
fun main() {
    val inputAsInteger = readln().toIntOrNull()
    val output = when (inputAsInteger) {
        null -> "dude, enter a valid number"
        3 -> "the number is 3"
        5 -> "the number is 5"
        in 10..20 -> "the number is between 10 and 20"
        else -> "something else"
    }
    println(output)
}
```

```java
public class Main {
    public static void main(String[] args) {
        Integer inputAsInteger = 15;
        final String output = switch (inputAsInteger) {
            case null -> "dude, enter a valid number";
            case 3 -> "the number is 3";
            case 5 -> "the number is 5";
            case Integer n when n >= 10 && n <= 20 -> "the number is between 10 and 20";
            default -> "something else";
        };
        System.out.println(output);
    }
}
```

Kotlin ใช้ `when` ได้ทั้งแบบไม่มีตัวให้เทียบ และแบบ `when (ค่า)` พร้อมช่วง `in 10..20` ส่วน Java 21 ใช้ `switch` เป็น expression ด้วยลูกศร `->` และ `default` แทน `else`
