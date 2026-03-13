package util

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"text/template"
	"time"

	"github.com/ensoria/faker/pkg/faker/common/log"
)

// ConvertToAnySlice converts a typed slice (e.g. []string, []int) to []any.
// Note: This is convenient but has a performance cost, so use it judiciously.
//
// 型付きスライス（例: []string, []int）を[]anyに変換する。
// 注: 便利だが、パフォーマンスコストがかかるので、使いどころに注意。
func ConvertToAnySlice[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}

// RenderTemplate treats the format string as a Go template, embeds the data, and returns the result.
//
// フォーマット文字列をGoテンプレートとして扱い、データを埋め込み、結果の文字列を返す。
func RenderTemplate(format string, data interface{}) string {
	errMsgTmpl := "Failed to render locale template: %v"
	tmpl, err := template.New(format).Parse(format)
	if err != nil {
		errMsg := fmt.Sprintf(errMsgTmpl, err)
		log.GeneralError(errMsg, 2)
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		errMsg := fmt.Sprintf(errMsgTmpl, err)
		log.GeneralError(errMsg, 2)
		return ""
	}

	return buf.String()
}

// RandSeed creates a new rand.Rand instance seeded with the current time.
//
// 現在時刻をシードとして新しいrand.Randインスタンスを作成する。
func RandSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// TruncateToPrecision truncates a float64 value to the specified number of decimal places.
//
// float64の値を指定された小数桁数で切り捨てる。
func TruncateToPrecision(val float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Trunc(val*multiplier) / multiplier
}

// CapFirstLetter capitalizes the first letter of the given string.
//
// 指定された文字列の最初の文字を大文字にする。
func CapFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
