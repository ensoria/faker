package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

func Output(funcName string, value any) {
	fmt.Printf("%s: [%v]\n", funcName, value)
}

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
func VisibleDecimalPlaces(val float64) int {
	strVal := strconv.FormatFloat(val, 'f', -1, 64)
	decimalVal := strings.Split(strVal, ".")
	if len(decimalVal) < 2 {
		return 0
	}
	return len(decimalVal[1])
}
