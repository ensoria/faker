package payment

import "strconv"

// CalcLuhnChecksum calculates the Luhn checksum of a numeric string.
// See: http://en.wikipedia.org/wiki/Luhn_algorithm
func CalcLuhnChecksum(number string) int {
	length := len(number)
	sum := 0

	// Sum odd digits from right (0-indexed from right: positions 0, 2, 4, ...)
	for i := length - 1; i >= 0; i -= 2 {
		sum += int(number[i] - '0')
	}

	// Sum even digits from right, doubling each
	for i := length - 2; i >= 0; i -= 2 {
		doubled := int(number[i]-'0') * 2
		// Sum the digits of the doubled value (e.g., 14 -> 1+4=5)
		sum += doubled/10 + doubled%10
	}

	return sum % 10
}

// ComputeLuhnCheckDigit calculates the check digit that should be appended
// to make a number valid according to the Luhn algorithm.
func ComputeLuhnCheckDigit(partialNumber string) string {
	checkDigit := CalcLuhnChecksum(partialNumber + "0")

	if checkDigit == 0 {
		return "0"
	}

	return strconv.Itoa(10 - checkDigit)
}

// IsLuhnValid checks whether a number (partial number + check digit) is Luhn compliant.
func IsLuhnValid(number string) bool {
	return CalcLuhnChecksum(number) == 0
}
