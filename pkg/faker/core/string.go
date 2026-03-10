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

type RandStr struct {
	rand *rand.Rand
}

func NewRandStr(rand *rand.Rand) *RandStr {
	return &RandStr{
		rand,
	}
}

// 以下のは全部coreに入れる
// アルファベット、数字、特殊文字を含むランダムな文字を返す
func (r *RandStr) Char() string {
	return string(allRunes[r.rand.Intn(len(allRunes))])
}

// これはアルファベットのみ
func (r *RandStr) Letter() string {
	return string(letterRunes[r.rand.Intn(len(letterRunes))])
}

func (r *RandStr) Digit() string {
	return string(numberRunes[r.rand.Intn(len(numberRunes))])
}

func (r *RandStr) NonZeroDigit() string {
	return string(nonZeroNumberRunes[r.rand.Intn(len(nonZeroNumberRunes))])
}

// min以上max以下のランダムな長さの文字列を返す。maxで指定した長さは含まれる。
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

// 固定の長さの文字列を返す
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

// 指定した文字列の特殊記号をランダムな文字に置き換えて返す。
//   - '?' → ランダムなアルファベット (a-zA-Z)
//   - '#' → ランダムな数字 (0-9)
//   - '%' → ゼロを除くランダムな数字 (1-9)
//   - '*' → アルファベットと数字のどちらかにランダムに置き換える
//
// 例えば、likeが"??-??1??X##"の場合、"ab-cd1efX35"のような文字列を返す。
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
