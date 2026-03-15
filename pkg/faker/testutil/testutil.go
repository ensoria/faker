package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

// Output prints the function name and its return value in a formatted string for test debugging.
//
// テストデバッグ用に、関数名とその戻り値をフォーマットされた文字列で出力する。
func Output(funcName string, value any) {
	fmt.Printf("%s: [%v]\n", funcName, value)
}

// IsInSlice reports whether val is contained in slice.
//
// valがスライスに含まれているかどうかを返す。
func IsInSlice[T comparable](val T, slice []T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// VisibleDecimalPlaces returns the number of visible decimal places of a float64 value.
// Note: trailing zeros are not preserved in float64, so this returns the apparent
// number of decimal digits (e.g. 1.500000 → 1.5 → 1, not 6).
//
// float64値の見かけ上の小数点以下の桁数を返す。
// 注意: float64では末尾のゼロは保持されないため、見かけ上の小数桁数を返す（例: 1.500000 → 1.5 → 1であり、6ではない）。
func VisibleDecimalPlaces(val float64) int {
	strVal := strconv.FormatFloat(val, 'f', -1, 64)
	decimalVal := strings.Split(strVal, ".")
	if len(decimalVal) < 2 {
		return 0
	}
	return len(decimalVal[1])
}
