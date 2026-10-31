# Console input

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 32:40–36:41

## Concept

`readln()` อ่านหนึ่งบรรทัดจาก console แล้วคืน `String`

```kotlin
println("please enter a number")
val input = readln()
println("you've entered $input")
```

โปรแกรมจะหยุดรอจนผู้ใช้พิมพ์แล้วกด Enter เช่น พิมพ์ `50` จะได้ข้อความ `you've entered 50`

เอา `String` ไป `% 2` ไม่ได้ IDE หาฟังก์ชัน modulo ของ type นี้ไม่เจอ เพราะข้อความอาจเป็น `hello world` ซึ่งไม่มีเศษจากการหาร ต้องแปลงเป็นจำนวนเต็มก่อน

```kotlin
val inputAsInteger = input.toInt()
val isEven = inputAsInteger % 2 == 0
println("is even: $isEven")
```

`45` ได้ `false`, เลขคู่ได้ `true` ถ้าพิมพ์ `hello world` โปรแกรมล่มด้วย `NumberFormatException` เพราะ `toInt()` แปลงได้เฉพาะข้อความที่เป็นจำนวนเต็ม ถ้าไม่บอกว่าให้ทำอะไรเมื่อค่าไม่ใช่ตัวเลข พฤติกรรมเริ่มต้นคือแครช เขายกเรื่องนี้ไปต่อที่ nullability

## Example

```kotlin
fun main() {
    println("please enter a number")
    val input = readln()
    val inputAsInteger = input.toInt()
    val isEven = inputAsInteger % 2 == 0
    println("is even: $isEven")
}
```

```java
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println("please enter a number");
        String input = scanner.nextLine();
        int inputAsInteger = Integer.parseInt(input);
        boolean isEven = inputAsInteger % 2 == 0;
        System.out.println("is even: " + isEven);
    }
}
```

Kotlin อ่านบรรทัดด้วย `readln()` และแปลงด้วย `toInt()` ส่วน Java ในตัวอย่างนี้ใช้ `Scanner` แล้วเรียก `Integer.parseInt` ทั้งคู่โยน `NumberFormatException` เมื่อข้อความไม่ใช่จำนวนเต็ม
