# faker

A Go library for generating random fake data. Create dummy data easily for testing and other purposes.

This library is a Go port of the PHP [Faker](https://github.com/FakerPHP/Faker), with modifications and changes.

> **日本語の説明は英語の説明の後に記載されています。** [→ 日本語版へジャンプ](#faker日本語)
> *Japanese documentation follows the English documentation below.*

---

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Creating a faker Instance](#creating-a-faker-instance)
- [Logging](#logging)
- [Methods](#methods)
  - [Rand](#rand---primitive-type-fakes)
  - [Global](#global)
  - [Locale](#locale)
  - [Global/Locale](#globallocale)

---

## Requirements

```
Go >= 1.24
```

## Installation

```
go get github.com/ensoria/faker/pkg/faker
```

## Creating a faker Instance

```go
// By default, the locale is set to en_US
f := faker.Create()

// To create with a Japanese locale, use CreateWithLocale()
j := ja_JP.New()
jp := faker.CreateWithLocale(j)
```

## Logging

faker does not output any logs by default.
If you want to output warning logs when calling methods that don't exist for the current locale or when using invalid arguments, set a logger with `SetLogger`.

```go
import (
	"log"
	"os"

	fakerlog "github.com/ensoria/faker/pkg/faker/common/log"
)

// Example: output logs to stderr
fakerlog.SetLogger(log.New(os.Stderr, "[faker] ", log.LstdFlags))

// Disable logging (default)
fakerlog.SetLogger(nil)
```

## Methods


### Rand - Primitive Type Fakes

#### Bool

```go
f := faker.Create()

// 50% chance of true/false
fake := f.Rand.Bool.Evenly() // example: true

// 80% chance of true, 20% chance of false
fake := f.Rand.Bool.WeightedToTrue(0.8) // example: true
```

#### Num

```go
// Returns a random Int between 1 and 10 (inclusive)
f.Rand.Num.IntBt(1, 10) // example: 5

f.Rand.Num.Int32Bt(1, 10) // example: 4

f.Rand.Num.Int64Bt(1, 10) // example: 3

f.Rand.Num.Float32Bt(1.0, 10.0) // example: +2.056392e+000

f.Rand.Num.Float64Bt(1.0, 10.0) // example: +3.627652e+000

// Aliases for rand.Rand methods
f.Rand.Num.Int()
f.Rand.Num.Intn(10)
```

#### Str

```go
// Random character including alphabets, digits, and special characters
f.Rand.Str.Char() // example: "y"
// Single alphabet letter
f.Rand.Str.Letter() // example: "a"
// Single digit (0-9)
f.Rand.Str.Digit() // example: "1"
// Single non-zero digit (1-9)
f.Rand.Str.NonZeroDigit() // example: "5"
// Single special character
f.Rand.Str.SpecialChar() // example: "$"
// Random ASCII string of specified length (alphanumeric + special characters)
f.Rand.Str.RandomASCII(10) // example: "hS3Y.wnGtu"

f.Rand.Str.AlphaRange(5, 10) // example: "VLkwXtKTJ"
f.Rand.Str.AlphaFixedLength(10) // example: "PQRpBVWHow"
// # → digit (0-9), ? → letter, % → non-zero digit (1-9),
// ! → special character, * → letter or digit
f.Rand.Str.AlphaDigitsLike("###-???-***-!!!-%%%") // example: "391-lwe-11u-$&(-457"

```

#### Time

```go
past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)

f.Rand.Time.PastFuture() // example: 2022-01-08 12:24:06.622832978 +0900 JST

f.Rand.Time.PastFrom(past30Years) // example: 1995-05-25 20:49:02.288568665 +0900 JST m=-912941844.443213042

f.Rand.Time.Past() // example: 2002-10-09 19:18:57.246421312 +0900 JST m=-680185687.650174271

f.Rand.Time.FutureTo(future30Years) // example: 2045-03-13 10:18:07.04601675 +0900 JST m=+658635502.851367918

f.Rand.Time.Future() // example: 2041-07-25 21:11:12.119502619 +0900 JST m=+544021834.272691954

f.Rand.Time.TimeRange(past30Years, future30Years) // example: 2016-02-24 08:21:32.32422701 +0900 JST m=-258075587.937158531

f.Rand.Time.Duration() // example: 309679h21m56.609248762s

f.Rand.Time.DurationMilliSec() // example: 875.892572ms

f.Rand.Time.DurationMin() // example: 46m29.429733821s

f.Rand.Time.DurationHour() // example: 8h19m52.645864323s

f.Rand.Time.DurationTo(1 * time.Second) // example: 798.391093ms

f.Rand.Time.DurationRange(1*time.Second, 2*time.Second) // example: 1.818209421s


```

#### Slice

```go

f.Rand.Slice.IntElem([]int{1, 2, 3}) // example: 2

f.Rand.Slice.StrElem([]string{"foo", "bar", "bazz"}) // example: "foo"

```

#### Map

```go
simpleValues := map[any]any{
  "key1": "value1",
  "key2": "value2",
  "key3": "value3",
  "key4": "value4",
}

f.Rand.Map.KeyValue(simpleValues) // example: key4, value4

sliceValues := map[any][]any{
  1: {"value11", "value12"},
  2: {"value21", "value22"},
}

f.Rand.Map.KeySliceValue(sliceValues) // example: 1, [value11 value12]

```


## Global

Data generated regardless of locale.


### Barcode

```go
f.Barcode.EAN8() // example: "58594605"

f.Barcode.EAN13() // example: 5945059001019

f.Barcode.ISBN10() // example: 4509472889

f.Barcode.ISBN13() // example: 9787672549372

```

### Color

```go
f.Color.SafeName() // example: "silver"

f.Color.Name() // example: "Gainsboro"

f.Color.Hex() // example: "#0e457a"

f.Color.SafeHex() // example: "#ff3300"

f.Color.RGBAsNum() // example: 15, 247, 177

f.Color.RGBAsStr() // example: "161,181,228"

f.Color.RGBAsArr() // example: [98 35 65]

f.Color.RGBCSS() // example: "rgb(223,67,224)"

f.Color.RGBACSS() // example: "rgba(66,112,144,0.3)"

f.Color.HSLAsNum() // example: 153, 97, 56

f.Color.HSLAsStr() // example: "149,85,59"

f.Color.HSLAsArr() // example: [31 69 46]
```

### File

```go
f.File.MIMEType() // example: "application/widget"

f.File.Extension() // example: "tga"

destDir := "./tmp"
content := "Hello, World!"
returnFullPath := false
f.File.WriteWithText(destDir, content, "txt", returnFullPath)

srcFilePath := "./file/sample.txt"
f.File.CopyFrom(destDir, srcFilePath, "txt", returnFullPath)
```

### Image

```go
// Binary image data
binary, err := f.Image.Binary(100, 100, image.JPG)

// Base64-encoded image string
bs64Str, err := f.Image.Base64(100, 100, image.JPG)

// Returns an `image.Image` object usable directly in Go
obj, err := f.Image.Object(100, 100, image.JPG)

```

### Internet

```go
f.Internet.UserName() // example: ayla.prosacco

f.Internet.DomainWord() // example: klein

f.Internet.TLD() // example: com

f.Internet.DomainName() // example: gutkowski.info

f.Internet.Email() // example: charity.ziemann@mertz.net

f.Internet.Password(false) // example: XDb34186c6z76np12
f.Internet.Password(true) // example: 836a52'68?'S#12+5_31

f.Internet.IPv4() // example: 190.238.68.2

f.Internet.IPv6() // example: 112a:792e:884c:99e5:d7a0:2b2c:df2b:9c48

f.Internet.LocalIPv4() // example: 172.27.204.249

f.Internet.MACAddress() // example: 01:ED:77:9F:1C:E1

f.Internet.Slug(6, false) // example: voluptate-non-mollitia-sit-deleniti-vel

f.Internet.URL() // example: https://www.tremblay.info/deleniti-suscipit-corrupti-neque-ea-cum-odit

```

### Lorem

```go
f.Lorem.Word() // example: "qui"

// Returns a slice of strings with the specified count
f.Lorem.WordSliceFixedLength(5) // example: [consectetur quia reprehenderit est consectetur]

// Returns a random-length slice of strings up to the specified count
f.Lorem.WordSlice(5) // example: [nisi porro]

// Returns random words as a string up to the specified count
f.Lorem.Words(5) // example: "sint officia eveniet ut sint"

f.Lorem.SentenceFixedLength(5) // example: "Aut corrupti ullam delectus exercitationem."

f.Lorem.Sentence(5) // example: "Vel."

f.Lorem.SentenceSliceFixedLength(5, 5) // example: ["Quis." "Vitae et quisquam." "Earum." "Omnis." "Perferendis eius fugit voluptas qui."]

f.Lorem.SentenceSlice(5, 5) // example: ["Facilis." "Consequatur sed saepe necessitatibus et."]

f.Lorem.Sentences(5, 5) // example: "Voluptatem sed omnis vel repudiandae. Quo et. Sit optio ipsa beatae. Veritatis iusto."

f.Lorem.ParagraphSliceFixedLength(5, 5) // example:
// ["Et incidunt quia necessitatibus." "Porro ut ipsa nulla quos et dignissimos in." "Voluptas illo consectetur." "Illo." "Doloremque."" Eos placeat." "Nam nostrum sed necessitatibus voluptas provident est quibusdam saepe reprehenderit ut illum quae consequatur excepturi corporis illo voluptatum sint omnis magni qui adipisci voluptatem." "Sed vel veritatis dolores et voluptatum molestiae sequi aut." "Fugiat est ducimus et eos eligendi." "Omnis molestias dolorem animi sapiente voluptatem soluta nostrum qui reprehenderit." "Enim dolor aliquam mollitia beatae omnis autem sunt perspiciatis corrupti molestiae sunt sed qui id facilis laudantium ut eveniet." "Exercitationem quibusdam corporis alias porro vel." "Aut dolorum magni." "Voluptatem libero ipsa." "Eaque."]

f.Lorem.ParagraphSlice(5, 5) // example:
// ["Neque impedit inventore qui repellendus dolores nulla minima nulla ratione similique illum non asperiores error iusto." "Voluptatem soluta adipisci qui odio magnam fuga consequatur pariatur veniam aut quis ipsam quibusdam voluptatibus et sapiente." "Molestiae quaerat consectetur pariatur possimus." "Nisi et eum quia suscipit itaque magnam architecto porro ut earum vel possimus at commodi aliquid possimus est magni sit et molestias odit animi velit eos numquam animi voluptatum voluptatem." "Deserunt error id consequatur." "A quis." "Eius labore molestiae ut omnis." "Et nam." "Qui nobis aut. ""Consectetur eum aut non dolorem enim voluptas vitae." "Alias ullam voluptas est voluptatem dolore reprehenderit." "Voluptate qui possimus animi ut voluptatem quo asperiores." "Veritatis voluptatibus."]

f.Lorem.Paragraphs(5, 5) // example:
// "Veniam sed enim quidem blanditiis excepturi dicta molestias numquam enim. Nisi ipsum reiciendis vel voluptatum dolorum eum deleniti voluptas eum sed rem nulla."

```

### Medical

```go
f.Medical.BloodType() // example: "O"

f.Medical.BloodRhFactor() // example: "+"

f.Medical.BloodGroup() // example: "AB-"

```

### Payment

```go
f.Payment.CreditCardType() // example: "Visa"

f.Payment.CreditCardNumber("") // example: "4024007142538"

f.Payment.CreditCardNumber("Visa") // example: "4532463395885642"

f.Payment.CreditCardNumberFormatted("Visa", "-") // example: "4597-9650-4195-3089"

f.Payment.CreditCardExpirationDate(true) // example: `time.Time`: 2026-05-31 17:44:02.496343249 +0900 JST m=+7520088.956891708

f.Payment.CreditCardExpirationDateString(true, "") // example: 04/27

f.Payment.CreditCardDetailsResult(true, f.Person.Name()) // example: `payment.CreditCardDetails`: &{MasterCard 2590375307974525 John Doe 05/28}

f.Payment.IBAN("", "") // example: "DO125R9578724012889910398673"

f.Payment.IBAN("DE", "") // example: "DE78970818524005997887"

f.Payment.SWIFTBICNumber() // example: "UBAIYK48"

```


### User Agent

```go
f.UserAgent.RandomUserAgent() // example: "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/5342 (KHTML, like Gecko) Chrome/37.0.849.0 Mobile Safari/5342"

f.UserAgent.Chrome() // example: "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Mobile Safari/5312"

f.UserAgent.MSEdge() // example: "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/90.0.4664.61 Safari/535.2 Edg/90.01016.28"

f.UserAgent.Firefox() // example: "Mozilla/5.0 (X11; Linux i686; rv:7.0) Gecko/20101231 Firefox/36.0"

f.UserAgent.Safari() // example: "Mozilla/5.0 (Windows; U; Windows NT 5.1) AppleWebKit/534.12.5 (KHTML, like Gecko) Version/4.0 Safari/534.12.5"

f.UserAgent.Opera() // example: "Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00"

f.UserAgent.InternetExplorer() // example: "Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)"

f.UserAgent.WindowsPlatformToken() // example: "Windows NT 6.1"

f.UserAgent.MacPlatformToken() // example: "Macintosh; Intel Mac OS X 10_7_3"

f.UserAgent.IOSMobileToken() // example: "iPhone; CPU iPhone OS 14_1 like Mac OS X"

f.UserAgent.LinuxPlatformToken() // example: "X11; Linux x86_64"

f.UserAgent.MacProcessor() // example: "Intel"

f.UserAgent.LinuxProcessor() // example: "x86_64"

```

## Locale

Data varies depending on the locale passed when creating the instance.

### Address

```go
f.Address.CitySuffix() // example: "land"
// example ja_JP: "村"

f.Address.CityPrefix() // example: "Port"

f.Address.CityName() // example: "Crona"
// example ja_JP: "日高"

f.Address.City() // example: "Fayside"
// example ja_JP: "天童市"

f.Address.StreetSuffix() // example: "Parkways"

f.Address.StreetName() // example: "Mraz"

f.Address.Street() // example: "Considine Island"

f.Address.BuildingNumber() // example: "3114"

f.Address.SecondaryAddress() // example: "Suite 760"
// example ja_JP: "レジデンス加納8K号"

f.Address.StreetAddress() // example: "618 Lynch Apt. 42"

f.Address.Postcode() // example: "75602"
// example ja_JP: "6858030"

f.Address.StateAbbr() // example: "CA"

f.Address.State() // example: "Colorado"

f.Address.Address() // example:
// "0730 Gleason Apt. 34\n
// South West, NM 08956"
// example ja_JP: "334-7397  茨城県大仙町東区東夷川3-7-0"

f.Address.Country() // example: "United States Minor Outlying Islands"
// example ja_JP: "サウジアラビア"

f.Address.Prefecture() // example: "佐賀県"

f.Address.WardSuffix() // example: "区"

f.Address.WardName() // example: "東"

f.Address.Ward() // example: "北区"

f.Address.AreaName() // example: "北小路"

f.Address.AreaNumber() // example: "2-8-04"

f.Address.Area() // example: "谷口町7丁目1番地610"

f.Address.BuildingName() // example: "笹田"

f.Address.RoomNumber() // example: "804"

f.Address.Latitude() // example: 54.617171

f.Address.Longitude() // example: 74.851822

f.Address.LocalCoordinates() // example: 82.130718, -121.140770

```

### Company

```go
f.Company.CompanyName() // example: "Reichel"
// example ja_JP: "山口"
f.Company.CompanyPrefix() // example: ""
// example ja_JP: "株式会社"

f.Company.CompanySuffix() // example: "Group"
// example ja_JP: "有限会社"

f.Company.Name() // example: "Osinski-Schinner"
// example ja_JP: "株式会社 小林"
f.Company.JobTitleName() // example: "Telecommunications Equipment Installer"

f.Company.JobTitle() // example: "Airframe Mechanic"

f.Company.EINPrefix() // example: "54"

f.Company.EIN() // example: "52-9635645"

```

### Person


```go
f.Person.FirstNameMale() // example: "Xzavier"
// example ja_JP: "直樹"

f.Person.FirstNameFemale() // example: "Stephany"
// example ja_JP: "真綾"

f.Person.FirstName() // example: "Ernesto"
// example ja_JP: "学"

f.Person.LastName() // example: "Hauck"
// example ja_JP: "桐山"

f.Person.TitleMale() // example: "Prof."

f.Person.TitleFemale() // example: "Miss"

f.Person.Title() // example: "Dr."

f.Person.Suffix() // example: "Jr."

f.Person.MaleName() // example: "Charles Flatley"
// example ja_JP: "山口 康弘"

f.Person.FemaleName() // example: "Esther Rosenbaum"
// example ja_JP: "青山 あすか"

f.Person.Name() // example: "Gregory Braun III"
// example ja_JP: "佐藤 智也"

f.Person.SSN() // example: "733-20-4400"

f.Person.FirstKanaNameMale() // example: ""
// example ja_JP: "タクマ"

f.Person.FirstKanaNameFemale() // example: ""
// example ja_JP: "ミキ"

f.Person.FirstKanaName() // example: ""
// example ja_JP: "ツバサ"

f.Person.LastKanaName() // example: ""
// example ja_JP: "スギヤマ"

f.Person.MaleKanaName() // example: ""
// example ja_JP: "コンドウ ケンイチ"

f.Person.FemaleKanaName() // example: ""
// example ja_JP: "サトウ ミキ"

f.Person.KanaName() // example: ""
// example ja_JP: "イダカ サトミ"

```

## Global/Locale

### Phone Number

```go
f.PhoneNumber.PhoneNumber() // example: "201-886-0269"
// example ja_JP: "090-1234-5678"

f.PhoneNumber.E164PhoneNumber() // example: "+27113456789"

f.PhoneNumber.IMEI() // example: "354809024498147"

```

---

# faker（日本語）

fakerは、ランダムなダミーデータを作成するためのライブラリです。
テストのためなどに、適当なダミーデータを簡単に作成することができます。

このライブラリはPHPの[Faker](https://github.com/FakerPHP/Faker)を、修正と変更を加え、Goへ移植したものになります。

---

## 目次

- [要件](#要件)
- [インストール](#インストール)
- [fakerのインスタンス作成](#fakerのインスタンス作成)
- [ログ出力](#ログ出力)
- [メソッド](#メソッド)
  - [Rand](#rand-プリミティブ型のフェイク)
  - [Global](#global-1)
  - [Locale](#locale-1)
  - [Global/Locale](#globallocale-1)

---

## 要件

```
Go >= 1.24
```

## インストール

```
go get github.com/ensoria/faker/pkg/faker
```

## fakerのインスタンス作成

```go
// デフォルトでは、localeがen_USで作成されます
f := faker.Create()

// 日本語のロケールで作成する場合は、CreateWithLocale()を使います。
j := ja_JP.New()
jp := faker.CreateWithLocale(j)
```

## ログ出力

fakerはデフォルトではログを出力しません。
ロケールに存在しないメソッドの呼び出しや、不正な引数の使用時に警告ログを出力したい場合は、`SetLogger`でロガーを設定してください。

```go
import (
	"log"
	"os"

	fakerlog "github.com/ensoria/faker/pkg/faker/common/log"
)

// 標準エラー出力にログを出力する例
fakerlog.SetLogger(log.New(os.Stderr, "[faker] ", log.LstdFlags))

// ログ出力を無効にする（デフォルト）
fakerlog.SetLogger(nil)
```

## メソッド


### Rand プリミティブ型のフェイク

#### Bool

```go
f := faker.Create()

// true/falseが50%ずつの確率
fake := f.Rand.Bool.Evenly() // example: true

// trueが80%, falseが20の確率で返ります
fake := f.Rand.Bool.WeightedToTrue(0.8) // example: true
```

#### Num

```go
// 1から10までのIntを返します。引数に渡した数字が含まれた、ランダムなIntです。
// 例えば、ここでは、1と10は、ランダムな値に含まれます。
f.Rand.Num.IntBt(1, 10) // example: 5

f.Rand.Num.Int32Bt(1, 10) // example: 4

f.Rand.Num.Int64Bt(1, 10) // example: 3

f.Rand.Num.Float32Bt(1.0, 10.0) // example: +2.056392e+000

f.Rand.Num.Float64Bt(1.0, 10.0) // example: +3.627652e+000

// rand.Randのメソッドを使いたい場合は、エイリアスが用意されています
f.Rand.Num.Int()
f.Rand.Num.Intn(10)
```

#### Str

```go
// アルファベット、数字、特殊文字を含むランダムな文字
f.Rand.Str.Char() // example: "y"
// アルファベット1文字
f.Rand.Str.Letter() // example: "a"
// 数字1文字 (0-9)
f.Rand.Str.Digit() // example: "1"
// 0を除く数字1文字 (1-9)
f.Rand.Str.NonZeroDigit() // example: "5"
// 特殊文字1文字
f.Rand.Str.SpecialChar() // example: "$"
// 英数字+特殊文字を含むランダムな文字列（指定した長さ）
f.Rand.Str.RandomASCII(10) // example: "hS3Y.wnGtu"

f.Rand.Str.AlphaRange(5, 10) // example: "VLkwXtKTJ"
f.Rand.Str.AlphaFixedLength(10) // example: "PQRpBVWHow"
// #は数字(0-9)に、?はアルファベットに、%は0を除く数字(1-9)に、
// !は特殊文字に、*は英数のどちらかに置き換わります
f.Rand.Str.AlphaDigitsLike("###-???-***-!!!-%%%") // example: "391-lwe-11u-$&(-457"

```

#### Time

```go
past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)

f.Rand.Time.PastFuture() // example: 2022-01-08 12:24:06.622832978 +0900 JST

f.Rand.Time.PastFrom(past30Years) // example: 1995-05-25 20:49:02.288568665 +0900 JST m=-912941844.443213042

f.Rand.Time.Past() // example: 2002-10-09 19:18:57.246421312 +0900 JST m=-680185687.650174271

f.Rand.Time.FutureTo(future30Years) // example: 2045-03-13 10:18:07.04601675 +0900 JST m=+658635502.851367918

f.Rand.Time.Future() // example: 2041-07-25 21:11:12.119502619 +0900 JST m=+544021834.272691954

f.Rand.Time.TimeRange(past30Years, future30Years) // example: 2016-02-24 08:21:32.32422701 +0900 JST m=-258075587.937158531

f.Rand.Time.Duration() // example: 309679h21m56.609248762s

f.Rand.Time.DurationMilliSec() // example: 875.892572ms

f.Rand.Time.DurationMin() // example: 46m29.429733821s

f.Rand.Time.DurationHour() // example: 8h19m52.645864323s

f.Rand.Time.DurationTo(1 * time.Second) // example: 798.391093ms

f.Rand.Time.DurationRange(1*time.Second, 2*time.Second) // example: 1.818209421s


```

#### Slice

```go

f.Rand.Slice.IntElem([]int{1, 2, 3}) // example: 2

f.Rand.Slice.StrElem([]string{"foo", "bar", "bazz"}) // example: "foo"

```

#### Map

```go
simpleValues := map[any]any{
  "key1": "value1",
  "key2": "value2",
  "key3": "value3",
  "key4": "value4",
}

f.Rand.Map.KeyValue(simpleValues) // example: key4, value4

sliceValues := map[any][]any{
  1: {"value11", "value12"},
  2: {"value21", "value22"},
}

f.Rand.Map.KeySliceValue(sliceValues) // example: 1, [value11 value12]

```


## Global

ロケールに関係なく、同じデータが作成されます


### Barcode

```go
f.Barcode.EAN8() // example: "58594605"

f.Barcode.EAN13() // example: 5945059001019

f.Barcode.ISBN10() // example: 4509472889

f.Barcode.ISBN13() // example: 9787672549372

```

### Color

```go
f.Color.SafeName() // example: "silver"

f.Color.Name() // example: "Gainsboro"

f.Color.Hex() // example: "#0e457a"

f.Color.SafeHex() // example: "#ff3300"

f.Color.RGBAsNum() // example: 15, 247, 177

f.Color.RGBAsStr() // example: "161,181,228"

f.Color.RGBAsArr() // example: [98 35 65]

f.Color.RGBCSS() // example: "rgb(223,67,224)"

f.Color.RGBACSS() // example: "rgba(66,112,144,0.3)"

f.Color.HSLAsNum() // example: 153, 97, 56

f.Color.HSLAsStr() // example: "149,85,59"

f.Color.HSLAsArr() // example: [31 69 46]
```

### File

```go
f.File.MIMEType() // example: "application/widget"

f.File.Extension() // example: "tga"

destDir := "./tmp"
content := "Hello, World!"
returnFullPath := false
f.File.WriteWithText(destDir, content, "txt", returnFullPath)

srcFilePath := "./file/sample.txt"
f.File.CopyFrom(destDir, srcFilePath, "txt", returnFullPath)
```

### Image

```go
// バイナリのイメージデータ
binary, err := f.Image.Binary(100, 100, image.JPG)

// base64にエンコードされたイメージのstring
bs64Str, err := f.Image.Base64(100, 100, image.JPG)

// Goでそのまま扱える`image.Image`のオブジェクトを返します
obj, err := f.Image.Object(100, 100, image.JPG)

```

### Internet

```go
f.Internet.UserName() // example: ayla.prosacco

f.Internet.DomainWord() // example: klein

f.Internet.TLD() // example: com

f.Internet.DomainName() // example: gutkowski.info

f.Internet.Email() // example: charity.ziemann@mertz.net

f.Internet.Password(false) // example: XDb34186c6z76np12
f.Internet.Password(true) // example: 836a52'68?'S#12+5_31

f.Internet.IPv4() // example: 190.238.68.2

f.Internet.IPv6() // example: 112a:792e:884c:99e5:d7a0:2b2c:df2b:9c48

f.Internet.LocalIPv4() // example: 172.27.204.249

f.Internet.MACAddress() // example: 01:ED:77:9F:1C:E1

f.Internet.Slug(6, false) // example: voluptate-non-mollitia-sit-deleniti-vel

f.Internet.URL() // example: https://www.tremblay.info/deleniti-suscipit-corrupti-neque-ea-cum-odit

```

### Lorem

```go
f.Lorem.Word() // example: "qui"
// 指定した個数の文字列の配列を返す

f.Lorem.WordSliceFixedLength(5) // example: [consectetur quia reprehenderit est consectetur]

// 指定した文字数を上限としてランダムな個数の文字列のスライスを返す
f.Lorem.WordSlice(5) // example: [nisi porro]

// 指定した文字数を上限としてランダムな個数の文字列を返す
f.Lorem.Words(5) // example: "sint officia eveniet ut sint"

f.Lorem.SentenceFixedLength(5) // example: "Aut corrupti ullam delectus exercitationem."

f.Lorem.Sentence(5) // example: "Vel."

f.Lorem.SentenceSliceFixedLength(5, 5) // example: ["Quis." "Vitae et quisquam." "Earum." "Omnis." "Perferendis eius fugit voluptas qui."]

f.Lorem.SentenceSlice(5, 5) // example: ["Facilis." "Consequatur sed saepe necessitatibus et."]

f.Lorem.Sentences(5, 5) // example: "Voluptatem sed omnis vel repudiandae. Quo et. Sit optio ipsa beatae. Veritatis iusto."

f.Lorem.ParagraphSliceFixedLength(5, 5) // example:
// ["Et incidunt quia necessitatibus." "Porro ut ipsa nulla quos et dignissimos in." "Voluptas illo consectetur." "Illo." "Doloremque."" Eos placeat." "Nam nostrum sed necessitatibus voluptas provident est quibusdam saepe reprehenderit ut illum quae consequatur excepturi corporis illo voluptatum sint omnis magni qui adipisci voluptatem." "Sed vel veritatis dolores et voluptatum molestiae sequi aut." "Fugiat est ducimus et eos eligendi." "Omnis molestias dolorem animi sapiente voluptatem soluta nostrum qui reprehenderit." "Enim dolor aliquam mollitia beatae omnis autem sunt perspiciatis corrupti molestiae sunt sed qui id facilis laudantium ut eveniet." "Exercitationem quibusdam corporis alias porro vel." "Aut dolorum magni." "Voluptatem libero ipsa." "Eaque."]

f.Lorem.ParagraphSlice(5, 5) // example:
// ["Neque impedit inventore qui repellendus dolores nulla minima nulla ratione similique illum non asperiores error iusto." "Voluptatem soluta adipisci qui odio magnam fuga consequatur pariatur veniam aut quis ipsam quibusdam voluptatibus et sapiente." "Molestiae quaerat consectetur pariatur possimus." "Nisi et eum quia suscipit itaque magnam architecto porro ut earum vel possimus at commodi aliquid possimus est magni sit et molestias odit animi velit eos numquam animi voluptatum voluptatem." "Deserunt error id consequatur." "A quis." "Eius labore molestiae ut omnis." "Et nam." "Qui nobis aut. ""Consectetur eum aut non dolorem enim voluptas vitae." "Alias ullam voluptas est voluptatem dolore reprehenderit." "Voluptate qui possimus animi ut voluptatem quo asperiores." "Veritatis voluptatibus."]

f.Lorem.Paragraphs(5, 5) // example:
// "Veniam sed enim quidem blanditiis excepturi dicta molestias numquam enim. Nisi ipsum reiciendis vel voluptatum dolorum eum deleniti voluptas eum sed rem nulla."

```

### Medical

```go
f.Medical.BloodType() // example: "O"

f.Medical.BloodRhFactor() // example: "+"

f.Medical.BloodGroup() // example: "AB-"

```

### Payment

```go
f.Payment.CreditCardType() // example: "Visa"

f.Payment.CreditCardNumber("") // example: "4024007142538"

f.Payment.CreditCardNumber("Visa") // example: "4532463395885642"

f.Payment.CreditCardNumberFormatted("Visa", "-") // example: "4597-9650-4195-3089"

f.Payment.CreditCardExpirationDate(true) // example: `time.Time`: 2026-05-31 17:44:02.496343249 +0900 JST m=+7520088.956891708

f.Payment.CreditCardExpirationDateString(true, "") // example: 04/27

f.Payment.CreditCardDetailsResult(true, f.Person.Name()) // example: `payment.CreditCardDetails`: &{MasterCard 2590375307974525 John Doe 05/28}

f.Payment.IBAN("", "") // example: "DO125R9578724012889910398673"

f.Payment.IBAN("DE", "") // example: "DE78970818524005997887"

f.Payment.SWIFTBICNumber() // example: "UBAIYK48"

```


### User Agent

```go
f.UserAgent.RandomUserAgent() // example: "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/5342 (KHTML, like Gecko) Chrome/37.0.849.0 Mobile Safari/5342"

f.UserAgent.Chrome() // example: "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Mobile Safari/5312"

f.UserAgent.MSEdge() // example: "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/90.0.4664.61 Safari/535.2 Edg/90.01016.28"

f.UserAgent.Firefox() // example: "Mozilla/5.0 (X11; Linux i686; rv:7.0) Gecko/20101231 Firefox/36.0"

f.UserAgent.Safari() // example: "Mozilla/5.0 (Windows; U; Windows NT 5.1) AppleWebKit/534.12.5 (KHTML, like Gecko) Version/4.0 Safari/534.12.5"

f.UserAgent.Opera() // example: "Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00"

f.UserAgent.InternetExplorer() // example: "Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)"

f.UserAgent.WindowsPlatformToken() // example: "Windows NT 6.1"

f.UserAgent.MacPlatformToken() // example: "Macintosh; Intel Mac OS X 10_7_3"

f.UserAgent.IOSMobileToken() // example: "iPhone; CPU iPhone OS 14_1 like Mac OS X"

f.UserAgent.LinuxPlatformToken() // example: "X11; Linux x86_64"

f.UserAgent.MacProcessor() // example: "Intel"

f.UserAgent.LinuxProcessor() // example: "x86_64"

```

## Locale

インスタンス作成時に、渡すロケールによって作成されるデータが変わります。

### Address

```go
f.Address.CitySuffix() // example: "land"
// example ja_JP: "村"

f.Address.CityPrefix() // example: "Port"

f.Address.CityName() // example: "Crona"
// example ja_JP: "日高"

f.Address.City() // example: "Fayside"
// example ja_JP: "天童市"

f.Address.StreetSuffix() // example: "Parkways"

f.Address.StreetName() // example: "Mraz"

f.Address.Street() // example: "Considine Island"

f.Address.BuildingNumber() // example: "3114"

f.Address.SecondaryAddress() // example: "Suite 760"
// example ja_JP: "レジデンス加納8K号"

f.Address.StreetAddress() // example: "618 Lynch Apt. 42"

f.Address.Postcode() // example: "75602"
// example ja_JP: "6858030"

f.Address.StateAbbr() // example: "CA"

f.Address.State() // example: "Colorado"

f.Address.Address() // example:
// "0730 Gleason Apt. 34\n
// South West, NM 08956"
// example ja_JP: "334-7397  茨城県大仙町東区東夷川3-7-0"

f.Address.Country() // example: "United States Minor Outlying Islands"
// example ja_JP: "サウジアラビア"

f.Address.Prefecture() // example: "佐賀県"

f.Address.WardSuffix() // example: "区"

f.Address.WardName() // example: "東"

f.Address.Ward() // example: "北区"

f.Address.AreaName() // example: "北小路"

f.Address.AreaNumber() // example: "2-8-04"

f.Address.Area() // example: "谷口町7丁目1番地610"

f.Address.BuildingName() // example: "笹田"

f.Address.RoomNumber() // example: "804"

f.Address.Latitude() // example: 54.617171

f.Address.Longitude() // example: 74.851822

f.Address.LocalCoordinates() // example: 82.130718, -121.140770

```

### Company

```go
f.Company.CompanyName() // example: "Reichel"
// example ja_JP: "山口"
f.Company.CompanyPrefix() // example: ""
// example ja_JP: "株式会社"

f.Company.CompanySuffix() // example: "Group"
// example ja_JP: "有限会社"

f.Company.Name() // example: "Osinski-Schinner"
// example ja_JP: "株式会社 小林"
f.Company.JobTitleName() // example: "Telecommunications Equipment Installer"

f.Company.JobTitle() // example: "Airframe Mechanic"

f.Company.EINPrefix() // example: "54"

f.Company.EIN() // example: "52-9635645"

```

### Person


```go
f.Person.FirstNameMale() // example: "Xzavier"
// example ja_JP: "直樹"

f.Person.FirstNameFemale() // example: "Stephany"
// example ja_JP: "真綾"

f.Person.FirstName() // example: "Ernesto"
// example ja_JP: "学"

f.Person.LastName() // example: "Hauck"
// example ja_JP: "桐山"

f.Person.TitleMale() // example: "Prof."

f.Person.TitleFemale() // example: "Miss"

f.Person.Title() // example: "Dr."

f.Person.Suffix() // example: "Jr."

f.Person.MaleName() // example: "Charles Flatley"
// example ja_JP: "山口 康弘"

f.Person.FemaleName() // example: "Esther Rosenbaum"
// example ja_JP: "青山 あすか"

f.Person.Name() // example: "Gregory Braun III"
// example ja_JP: "佐藤 智也"

f.Person.SSN() // example: "733-20-4400"

f.Person.FirstKanaNameMale() // example: ""
// example ja_JP: "タクマ"

f.Person.FirstKanaNameFemale() // example: ""
// example ja_JP: "ミキ"

f.Person.FirstKanaName() // example: ""
// example ja_JP: "ツバサ"

f.Person.LastKanaName() // example: ""
// example ja_JP: "スギヤマ"

f.Person.MaleKanaName() // example: ""
// example ja_JP: "コンドウ ケンイチ"

f.Person.FemaleKanaName() // example: ""
// example ja_JP: "サトウ ミキ"

f.Person.KanaName() // example: ""
// example ja_JP: "イダカ サトミ"

```

## Global/Locale

### Phone Number

```go
f.PhoneNumber.PhoneNumber() // example: "201-886-0269"
// example ja_JP: "090-1234-5678"

f.PhoneNumber.E164PhoneNumber() // example: "+27113456789"

f.PhoneNumber.IMEI() // example: "354809024498147"

```
