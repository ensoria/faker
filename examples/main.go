package main

import (
	"time"

	"github.com/ensoria/faker/pkg/faker"
	"github.com/ensoria/faker/pkg/faker/provider/locale/ja_JP"
)

func main() {
	f := faker.Create()
	jf := faker.CreateWithLocale(ja_JP.New())

	// bool
	f.Rand.Bool.Evenly()
	f.Rand.Bool.WeightedToTrue(0.8)

	// int
	f.Rand.Num.IntBt(1, 10)
	f.Rand.Num.Int32Bt(1, 10)
	f.Rand.Num.Int64Bt(1, 10)
	f.Rand.Num.Float32Bt(1.0, 10.0)
	f.Rand.Num.Float64Bt(1.0, 10.0)
	f.Rand.Num.Int()
	f.Rand.Num.Intn(10)

	// string
	f.Rand.Str.Char()
	f.Rand.Str.Letter()
	f.Rand.Str.Digit()
	f.Rand.Str.AlphaRange(5, 10)
	f.Rand.Str.AlphaFixedLength(10)
	f.Rand.Str.AlphaDigitsLike("###-???-***")

	// time
	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)
	f.Rand.Time.PastFuture()
	f.Rand.Time.PastFrom(past30Years)
	f.Rand.Time.Past()
	f.Rand.Time.FutureTo(future30Years)
	f.Rand.Time.Future()
	f.Rand.Time.TimeRange(past30Years, future30Years)
	f.Rand.Time.Duration()
	f.Rand.Time.DurationMilliSec()
	f.Rand.Time.DurationMin()
	f.Rand.Time.DurationHour()
	f.Rand.Time.DurationTo(1 * time.Second)
	f.Rand.Time.DurationRange(1*time.Second, 2*time.Second)

	// slice
	f.Rand.Slice.IntElem([]int{1, 2, 3})
	f.Rand.Slice.StrElem([]string{"foo", "bar", "bazz"})

	// map
	simpleValues := map[any]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	f.Rand.Map.KeyValue(simpleValues)

	sliceValues := map[any][]any{
		1: {"value11", "value12"},
		2: {"value21", "value22"},
	}
	f.Rand.Map.KeySliceValue(sliceValues)

	// barcode
	f.Barcode.EAN8()
	f.Barcode.EAN13()
	f.Barcode.ISBN10()
	f.Barcode.ISBN13()

	// color
	f.Color.SafeName()
	f.Color.Name()
	f.Color.Hex()
	f.Color.SafeHex()
	f.Color.RGBAsNum()
	f.Color.RGBAsStr()
	f.Color.RGBAsArr()
	f.Color.RGBCSS()
	f.Color.RGBACSS()
	f.Color.HSLAsNum()
	f.Color.HSLAsStr()
	f.Color.HSLAsArr()

	//  file
	f.File.MIMEType()
	f.File.Extension()

	// destDir := "./tmp"
	// content := "Hello, World!"
	// returnFullPath := false
	// f.File.WriteWithText(destDir, content, "txt", returnFullPath)

	// srcFilePath := "./file/sample.txt"
	// f.File.CopyFrom(destDir, srcFilePath, "txt", returnFullPath)

	// internet
	f.Internet.UserName()
	f.Internet.DomainWord()
	f.Internet.TLD()
	f.Internet.DomainName()
	f.Internet.Email()
	f.Internet.Password()
	f.Internet.IPv4()
	f.Internet.IPv6()
	f.Internet.LocalIPv4()
	f.Internet.MACAddress()

	// lorem
	f.Lorem.Word()
	f.Lorem.WordSliceFixedLength(5)
	f.Lorem.WordSlice(5)
	f.Lorem.Words(5)
	f.Lorem.SentenceFixedLength(5)
	f.Lorem.Sentence(5)
	f.Lorem.SentenceSliceFixedLength(5, 5)
	f.Lorem.SentenceSlice(5, 5)
	f.Lorem.Sentences(5, 5)
	f.Lorem.ParagraphSliceFixedLength(5, 5)
	f.Lorem.ParagraphSlice(5, 5)
	f.Lorem.Paragraphs(5, 5)

	// medical
	f.Medical.BloodType()
	f.Medical.BloodRhFactor()
	f.Medical.BloodGroup()

	// payment
	f.Payment.CreditCardType()
	f.Payment.CreditCardNumber("")
	f.Payment.CreditCardNumber("Visa")
	f.Payment.CreditCardNumberFormatted("Visa", "-")
	f.Payment.CreditCardExpirationDate(true)
	f.Payment.CreditCardExpirationDateString(true, "")
	f.Payment.CreditCardDetailsResult(true, f.Person.Name())
	f.Payment.Iban("", "")
	f.Payment.Iban("DE", "")
	f.Payment.SwiftBicNumber()

	// phone number
	f.PhoneNumber.PhoneNumber()
	f.PhoneNumber.E164PhoneNumber()
	f.PhoneNumber.IMEI()

	// user agent
	f.UserAgent.RandomUserAgent()
	f.UserAgent.Chrome()
	f.UserAgent.MsEdge()
	f.UserAgent.Firefox()
	f.UserAgent.Safari()
	f.UserAgent.Opera()
	f.UserAgent.InternetExplorer()
	f.UserAgent.WindowsPlatformToken()
	f.UserAgent.MacPlatformToken()
	f.UserAgent.IosMobileToken()
	f.UserAgent.LinuxPlatformToken()
	f.UserAgent.MacProcessor()
	f.UserAgent.LinuxProcessor()

	// address
	f.Address.CitySuffix()
	f.Address.CityPrefix()
	f.Address.CityName()
	f.Address.City()
	f.Address.StreetSuffix()
	f.Address.StreetName()
	f.Address.Street()
	f.Address.BuildingNumber()
	f.Address.SecondaryAddress()
	f.Address.StreetAddress()
	jf.Address.Postcode()
	f.Address.StateAbbr()
	f.Address.State()
	f.Address.Address()
	f.Address.Country()

	jf.Address.Prefecture()
	jf.Address.WardSuffix()
	jf.Address.WardName()
	jf.Address.Ward()
	jf.Address.AreaName()
	jf.Address.AreaNumber()
	jf.Address.Area()
	jf.Address.BuildingName()
	jf.Address.RoomNumber()

	f.Address.Latitude()
	f.Address.Longitude()
	f.Address.LocalCoordinates()

	// company
	f.Company.CompanyName()
	jf.Company.CompanyPrefix()
	f.Company.CompanySuffix()
	f.Company.Name()
	f.Company.JobTitleName()
	f.Company.JobTitle()
	f.Company.EINPrefix()
	f.Company.EIN()

	// person
	f.Person.FirstNameMale()
	f.Person.FirstNameFemale()
	f.Person.FirstName()
	f.Person.LastName()
	f.Person.TitleMale()
	f.Person.TitleFemale()
	f.Person.Title()
	f.Person.Suffix()
	f.Person.MaleName()
	f.Person.FemaleName()
	f.Person.Name()
	f.Person.Ssn()
	// ja_JP only
	jf.Person.FirstKanaNameMale()
	jf.Person.FirstKanaNameFemale()
	jf.Person.FirstKanaName()
	jf.Person.LastKanaName()
	jf.Person.MaleKanaName()
	jf.Person.FemaleKanaName()
	jf.Person.KanaName()

}
