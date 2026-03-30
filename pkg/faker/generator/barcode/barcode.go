package barcode

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/ensoria/gofake/pkg/faker/core"
)

// Barcode provides methods for generating random barcode strings.
//
// ランダムなバーコード文字列を生成するメソッドを提供する構造体。
type Barcode struct {
	rand *core.Rand
}

// New creates a new Barcode instance with the given random source.
//
// 指定されたランダムソースで新しいBarcodeインスタンスを作成する。
func New(rand *core.Rand) *Barcode {
	return &Barcode{
		rand,
	}
}

// EAN8 returns a random valid EAN-8 barcode string.
// Example: "12345678"
//
// ランダムな有効なEAN-8バーコード文字列を返す。
func (b *Barcode) EAN8() string {
	return b.ean(8)
}

// EAN13 returns a random valid EAN-13 barcode string.
// Example: "1234567890123"
//
// ランダムな有効なEAN-13バーコード文字列を返す。
func (b *Barcode) EAN13() string {
	return b.ean(13)
}

func (b *Barcode) ean(length int) string {
	if length != 8 && length != 13 {
		panic(fmt.Sprintf("Invalid length for EAN barcode. EAN code length must be 8 or 13. Given length: [%d]", length))
	}
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, length-1))
	code := b.rand.Str.AlphaDigitsLike(format)
	return code + fmt.Sprint(CalcEANCheckDigit(code))
}

// CalcEANCheckDigit computes the checksum of an EAN number.
// See: https://en.wikipedia.org/wiki/International_Article_Number
//
// EAN番号のチェックサムを計算する。
func CalcEANCheckDigit(digits string) int {
	sequence := []int{1, 3}
	if len(digits)+1 == 8 {
		sequence = []int{3, 1}
	}
	sums := 0

	for n, digit := range digits {
		sums += int(digit-'0') * sequence[n%2]
	}

	return (10 - sums%10) % 10
}

// ISBN10 returns a random valid ISBN-10 code.
// See: http://en.wikipedia.org/wiki/International_Standard_Book_Number
// Example: "3254681223"
//
// ランダムな有効なISBN-10コードを返す。
func (b *Barcode) ISBN10() string {
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, 9))
	code := b.rand.Str.AlphaDigitsLike(format)
	return code + fmt.Sprint(CalcISBNCheckDigit(code))
}

// ISBN13 returns a random valid ISBN-13 code.
// See: http://en.wikipedia.org/wiki/International_Standard_Book_Number
// Example: "9783254681223"
//
// ランダムな有効なISBN-13コードを返す。
func (b *Barcode) ISBN13() string {
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, 9))
	prefX := b.rand.Num.IntBt(8, 9)
	prefix := "97" + fmt.Sprint(prefX)

	pubBookCode := b.rand.Str.AlphaDigitsLike(format)
	code := prefix + pubBookCode
	// EAN check digit is used here because it's 12 digits
	return code + fmt.Sprint(CalcEANCheckDigit(code))
}

// CalcISBNCheckDigit calculates the ISBN-10 check digit.
// See: http://en.wikipedia.org/wiki/International_Standard_Book_Number#ISBN-10_check_digits
//
// ISBN-10のチェックディジットを計算する。
func CalcISBNCheckDigit(input string) string {
	// We're calculating check digit for ISBN-10
	// so, the length of the input should be 9
	length := 9
	if len(input) != length {
		panic(fmt.Sprintf("input length should be equal to %d", length))
	}

	sum := 0
	for i, digit := range input {
		sum += (10 - i) * int(digit-'0')
	}

	result := (11 - sum%11) % 11

	// 10 is replaced by X
	if result < 10 {
		return strconv.Itoa(result)
	}
	return "X"
}
