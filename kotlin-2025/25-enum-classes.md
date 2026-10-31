# Enum classes

ที่มา: Philipp Lackner — Full 2025 Kotlin Crash Course For Beginners (`dzUc9vrsldM`) · 2:37:34–2:44:19

## Concept

Enum ใช้เมื่อมีชุดตัวเลือกที่เกี่ยวข้องกันและรู้ครบตั้งแต่ตอน compile ชื่อค่าเป็นตัวพิมพ์ใหญ่ทั้งหมด คั่นคำด้วย `_` เขาเรียกแบบนี้ว่า snake case แบบตัวใหญ่

```kotlin
enum class Country {
    GERMANY,
    FRANCE,
    USA
}
```

ฟังก์ชัน `greetMe` รับ `Country` แล้วคืนคำทักทายด้วย `when` ถ้าเงื่อนไขครบทุกค่าอยู่แล้ว ใน IntelliJ กด Alt+Enter เลือก add remaining branches จะได้สาขาของทุกประเทศให้เอง

- เยอรมนีได้ `guten tag`
- ฝรั่งเศสได้ `bonjour`
- สหรัฐได้ `hello`

เรียกเป็น `greetMe(Country.GERMANY)` สลับเป็น `Country.USA` จะได้ `hello`

ตัว enum ใส่ฟังก์ชันในตัวได้ แต่เขาบอกว่าไม่ค่อยได้ทำแบบ class ปกติ จุดที่คนสับสนกับ sealed คือทั้งคู่จำกัดจำนวนแบบได้ เขาเลยเทียบกับรูปทรง สี่เหลี่ยมไม่ใช่ค่าคงที่ เพราะมีได้หลายรูป แต่ละรูปกว้างสูงต่างกัน และขนาดมักยังไม่รู้ตอน compile ประเทศในตัวอย่างนี้คงที่ มีเยอรมนีประเทศเดียว ไม่มีความหมายที่จะสร้างเยอรมนีสองก้อน

ถ้าแต่ละค่ามีข้อมูลที่คงที่เหมือนกัน ให้ใส่ constructor แล้วเรียกตรงจุดประกาศค่า เพราะค่านี้รู้ตอน compile เขาใส่รหัสประเทศของเยอรมนีเป็น `"DE"`

```kotlin
enum class Country(val countryCode: String) {
    GERMANY("DE"),
    FRANCE("FR"),
    USA("US")
}
```

Enum ได้ `entries` มาให้ วนทุกประเทศแล้วพิมพ์ `countryCode` ได้ Sealed interface กับ sealed class ไม่มีรายการแบบนี้ พอเพิ่ม `BRAZIL` พร้อมรหัส `when` จะไม่ครบจนกว่าจะมีสาขาใหม่ และลูปที่เดิน `entries` จะพิมพ์รหัสใหม่ให้เองโดยไม่ต้องไปแก้ลูป

เขายกงาน Android ที่เมนูดรอปดาวน์แต่ละตัวเลือกเป็นค่าใน enum พอเพิ่มค่า หน้าจอที่วาดจาก `entries` จะยาวขึ้นเอง

ใช้ enum เมื่อจำนวนตัวเลือกจำกัด และข้อมูลใน constructor ของแต่ละตัวเป็นค่าคงที่ ใช้ sealed เมื่อแต่ละแบบยังสร้างได้อีกหลาย instance ที่ข้อมูลต่างกัน

## Example

```kotlin
enum class Country(val countryCode: String) {
    GERMANY("DE"),
    FRANCE("FR"),
    USA("US"),
    BRAZIL("BR")
}

fun greetMe(country: Country): String {
    return when (country) {
        Country.GERMANY -> "guten tag"
        Country.FRANCE -> "bonjour"
        Country.USA -> "hello"
        Country.BRAZIL -> "olá"
    }
}

fun main() {
    println(greetMe(Country.GERMANY))
    for (country in Country.entries) {
        println(country.countryCode)
    }
}
```

```java
enum Country {
    GERMANY("DE"),
    FRANCE("FR"),
    USA("US"),
    BRAZIL("BR");

    final String countryCode;

    Country(String countryCode) {
        this.countryCode = countryCode;
    }
}

public class Main {
    static String greetMe(Country country) {
        return switch (country) {
            case GERMANY -> "guten tag";
            case FRANCE -> "bonjour";
            case USA -> "hello";
            case BRAZIL -> "olá";
        };
    }

    public static void main(String[] args) {
        System.out.println(greetMe(Country.GERMANY));
        for (Country country : Country.values()) {
            System.out.println(country.countryCode);
        }
    }
}
```

Kotlin เขียน `enum class` แล้วอ่านทุกค่าจาก `entries` ส่วน Java เขียน `enum` แล้วอ่านทุกค่าจาก `values()`

💡 เพิ่มเติม: เสียงบรรยายพูดรหัส `"DE"` ชัด รหัส `FR`, `US`, `BR` และคำทักทายของบราซิลในตัวอย่างเติมให้ครบตามชื่อประเทศที่เขาไล่ เขาบอกเองว่าไม่รู้คำทักทายบราซิลแล้วให้ IDE เติมสาขาให้
