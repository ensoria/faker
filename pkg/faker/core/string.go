package core

import (
	"fmt"
	"math/rand"

	"github.com/ensoria/faker/pkg/faker/common/log"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const nonZeroNumbers = "123456789"
const specialChars = "!@#$%^&*()_+{}|:<>?-=[]\\;',./"

var letterRunes = []rune(letters)
var numberRunes = []rune(numbers)
var nonZeroNumberRunes = []rune(nonZeroNumbers)
var specialCharRunes = []rune(specialChars)
var allRunes = []rune(letters + numbers + specialChars)

// RandStr provides methods for generating random strings.
// ランダムな文字列を生成するメソッドを提供する構造体。
type RandStr struct {
	rand *rand.Rand
}

// NewRandStr creates a new RandStr instance with the given random source.
// 指定されたランダムソースで新しいRandStrインスタンスを作成する。
func NewRandStr(rand *rand.Rand) *RandStr {
	return &RandStr{
		rand,
	}
}

// Char returns a random character from letters, digits, and special characters.
// アルファベット、数字、特殊文字を含むランダムな1文字を返す。
func (r *RandStr) Char() string {
	return string(allRunes[r.rand.Intn(len(allRunes))])
}

// Letter returns a random alphabetic character (a-zA-Z).
// ランダムなアルファベット1文字を返す（a-zA-Z）。
func (r *RandStr) Letter() string {
	return string(letterRunes[r.rand.Intn(len(letterRunes))])
}

// Digit returns a random digit character (0-9).
// ランダムな数字1文字を返す（0-9）。
func (r *RandStr) Digit() string {
	return string(numberRunes[r.rand.Intn(len(numberRunes))])
}

// NonZeroDigit returns a random non-zero digit character (1-9).
// ゼロを除くランダムな数字1文字を返す（1-9）。
func (r *RandStr) NonZeroDigit() string {
	return string(nonZeroNumberRunes[r.rand.Intn(len(nonZeroNumberRunes))])
}

// SpecialChar returns a random special character (!@#$%^&*()_+{}|:<>?-=[]\;',./).
//
// ランダムな特殊文字を返す（!@#$%^&*()_+{}|:<>?-=[]\;',./）。
func (r *RandStr) SpecialChar() string {
	return string(specialCharRunes[r.rand.Intn(len(specialCharRunes))])
}

// RandomASCII returns a random string of the given length containing
// letters, digits, and special characters.
// 指定された長さのアルファベット、数字、特殊文字を含むランダムな文字列を返す。
func (r *RandStr) RandomASCII(length int) string {
	if length < 0 {
		errMsg := fmt.Sprintf("Invalid length: %d", length)
		log.WrongUsage(errMsg, 1)
		return ""
	}
	var result string
	for i := 0; i < length; i++ {
		result += string(allRunes[r.rand.Intn(len(allRunes))])
	}
	return result
}

// AlphaRange returns a random alphabetic string with a length between min and max (inclusive).
// min以上max以下のランダムな長さのアルファベット文字列を返す。maxで指定した長さは含まれる。
func (r *RandStr) AlphaRange(min int, max int) string {
	if (min < 0) || (max < 0) || (min > max) || (min == max) {
		errMsg := fmt.Sprintf("Invalid range: min=%d, max=%d", min, max)
		log.WrongUsage(errMsg, 1)
		return ""
	}
	randomMax := r.rand.Intn(max-min+1) + min
	var result string
	for i := 0; i < randomMax; i++ {
		result += r.Letter()
	}
	return result
}

// AlphaFixedLength returns a random alphabetic string of the specified fixed length.
// 指定された固定長のランダムなアルファベット文字列を返す。
func (r *RandStr) AlphaFixedLength(length int) string {
	if length < 0 {
		errMsg := fmt.Sprintf("Invalid length: %d", length)
		log.WrongUsage(errMsg, 1)
		return ""
	}
	var result string
	for i := 0; i < length; i++ {
		result += r.Letter()
	}
	return result
}

// AlphaDigitsLike replaces special symbols in the given string with random characters:
//   - '?' -> random letter (a-zA-Z)
//   - '#' -> random digit (0-9)
//   - '%' -> random non-zero digit (1-9)
//   - '!' -> random special character (!@#$%^&*()_+{}|:<>?-=[]\;',./)
//   - '*' -> random letter or digit
//
// For example, "??-##1??X##" may return "ab-351efX35".
//
// 指定された文字列の特殊記号をランダムな文字に置き換えて返す:
//   - '?' -> ランダムなアルファベット (a-zA-Z)
//   - '#' -> ランダムな数字 (0-9)
//   - '%' -> ゼロを除くランダムな数字 (1-9)
//   - '!' -> ランダムな特殊文字 (!@#$%^&*()_+{}|:<>?-=[]\;',./)
//   - '*' -> アルファベットか数字のどちらかにランダムに置き換える
//
// 例えば、"??-##1??X##"の場合、"ab-351efX35"のような文字列を返す。
func (r *RandStr) AlphaDigitsLike(like string) string {
	result := ""
	for _, char := range []rune(like) {
		switch char {
		case '?':
			result += r.Letter()
		case '#':
			result += r.Digit()
		case '%':
			result += r.NonZeroDigit()
		case '!':
			result += r.SpecialChar()
		case '*':
			tmp := r.Letter()
			if r.rand.Intn(2) == 0 {
				tmp = r.Digit()
			}
			result += tmp
		default:
			result += string(char)
		}
	}
	return result
}
