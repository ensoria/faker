package payment

import "fmt"

// CalcIbanChecksum generates the two-digit IBAN checksum for a given IBAN string.
// The input iban should have the country code as the first two characters,
// followed by "00" as placeholder check digits, then the BBAN.
func CalcIbanChecksum(iban string) string {
	// Move first four characters to end and set check digits to "00"
	checkString := iban[4:] + iban[0:2] + "00"

	// Replace all letters with their numeric equivalents (A=10, B=11, ..., Z=35)
	numeric := ""
	for _, ch := range checkString {
		if ch >= 'A' && ch <= 'Z' {
			numeric += fmt.Sprintf("%d", ibanAlphaToNumber(ch))
		} else {
			numeric += string(ch)
		}
	}

	// Perform mod 97 and subtract from 98
	checksum := 98 - ibanMod97(numeric)

	return fmt.Sprintf("%02d", checksum)
}

// IsIbanValid checks whether an IBAN has a valid checksum.
func IsIbanValid(iban string) bool {
	return CalcIbanChecksum(iban) == iban[2:4]
}

// ibanAlphaToNumber converts an uppercase letter to its IBAN numeric equivalent.
// A=10, B=11, ..., Z=35
func ibanAlphaToNumber(ch rune) int {
	return int(ch) - 55 // 'A' is 65, so 65-55=10
}

// ibanMod97 calculates mod 97 on a numeric string.
// It processes the string digit by digit to avoid integer overflow.
func ibanMod97(number string) int {
	checksum := int(number[0] - '0')

	for i := 1; i < len(number); i++ {
		checksum = (10*checksum + int(number[i]-'0')) % 97
	}

	return checksum
}
