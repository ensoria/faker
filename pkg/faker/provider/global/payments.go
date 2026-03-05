package global

import "github.com/ensoria/faker/pkg/faker/provider"

func CreatePayments() *provider.Payments {
	return &provider.Payments{
		CardVendors: cardVendors,
		CardParams:  cardParams,
		IBANFormats: ibanFormats,
	}
}

// cardVendors defines the weighted list of card vendors.
// Duplicates increase the probability of selection.
var cardVendors = []string{
	"Visa", "Visa", "Visa", "Visa", "Visa",
	"MasterCard", "MasterCard", "MasterCard", "MasterCard", "MasterCard",
	"American Express", "Discover Card", "Visa Retired", "JCB",
}

// cardParams defines card brand masks for generating valid credit card numbers.
// '#' characters are replaced with random digits.
//
// See: https://en.wikipedia.org/wiki/Payment_card_number
// See: https://www.mastercard.us/en-us/issuers/get-support/2-series-bin-expansion.html
var cardParams = map[string][]string{
	"Visa": {
		"4539###########",
		"4556###########",
		"4916###########",
		"4532###########",
		"4929###########",
		"40240071#######",
		"4485###########",
		"4716###########",
		"4##############",
	},
	"Visa Retired": {
		"4539########",
		"4556########",
		"4916########",
		"4532########",
		"4929########",
		"40240071####",
		"4485########",
		"4716########",
		"4###########",
	},
	"MasterCard": {
		"2221###########",
		"23#############",
		"24#############",
		"25#############",
		"26#############",
		"2720###########",
		"51#############",
		"52#############",
		"53#############",
		"54#############",
		"55#############",
	},
	"American Express": {
		"34############",
		"37############",
	},
	"Discover Card": {
		"6011###########",
	},
	"JCB": {
		"3528###########",
		"3589###########",
	},
}

// ibanFormats defines IBAN formats by country code.
// Each element is [charClass, count] where charClass is:
//   - "n": numeric digit
//   - "a": uppercase letter
//   - "c": alphanumeric (uppercase letter or digit)
//
// Source: https://www.swift.com/standards/data-standards/iban
var ibanFormats = map[string][][2]any{
	"AD": {{"n", 4}, {"n", 4}, {"c", 12}},
	"AE": {{"n", 3}, {"n", 16}},
	"AL": {{"n", 8}, {"c", 16}},
	"AT": {{"n", 5}, {"n", 11}},
	"AZ": {{"a", 4}, {"c", 20}},
	"BA": {{"n", 3}, {"n", 3}, {"n", 8}, {"n", 2}},
	"BE": {{"n", 3}, {"n", 7}, {"n", 2}},
	"BG": {{"a", 4}, {"n", 4}, {"n", 2}, {"c", 8}},
	"BH": {{"a", 4}, {"c", 14}},
	"BR": {{"n", 8}, {"n", 5}, {"n", 10}, {"a", 1}, {"c", 1}},
	"CH": {{"n", 5}, {"c", 12}},
	"CR": {{"n", 4}, {"n", 14}},
	"CY": {{"n", 3}, {"n", 5}, {"c", 16}},
	"CZ": {{"n", 4}, {"n", 6}, {"n", 10}},
	"DE": {{"n", 8}, {"n", 10}},
	"DK": {{"n", 4}, {"n", 9}, {"n", 1}},
	"DO": {{"c", 4}, {"n", 20}},
	"EE": {{"n", 2}, {"n", 2}, {"n", 11}, {"n", 1}},
	"EG": {{"n", 4}, {"n", 4}, {"n", 17}},
	"ES": {{"n", 4}, {"n", 4}, {"n", 1}, {"n", 1}, {"n", 10}},
	"FI": {{"n", 6}, {"n", 7}, {"n", 1}},
	"FR": {{"n", 5}, {"n", 5}, {"c", 11}, {"n", 2}},
	"GB": {{"a", 4}, {"n", 6}, {"n", 8}},
	"GE": {{"a", 2}, {"n", 16}},
	"GI": {{"a", 4}, {"c", 15}},
	"GR": {{"n", 3}, {"n", 4}, {"c", 16}},
	"GT": {{"c", 4}, {"c", 20}},
	"HR": {{"n", 7}, {"n", 10}},
	"HU": {{"n", 3}, {"n", 4}, {"n", 1}, {"n", 15}, {"n", 1}},
	"IE": {{"a", 4}, {"n", 6}, {"n", 8}},
	"IL": {{"n", 3}, {"n", 3}, {"n", 13}},
	"IS": {{"n", 4}, {"n", 2}, {"n", 6}, {"n", 10}},
	"IT": {{"a", 1}, {"n", 5}, {"n", 5}, {"c", 12}},
	"KW": {{"a", 4}, {"n", 22}},
	"KZ": {{"n", 3}, {"c", 13}},
	"LB": {{"n", 4}, {"c", 20}},
	"LI": {{"n", 5}, {"c", 12}},
	"LT": {{"n", 5}, {"n", 11}},
	"LU": {{"n", 3}, {"c", 13}},
	"LV": {{"a", 4}, {"c", 13}},
	"MC": {{"n", 5}, {"n", 5}, {"c", 11}, {"n", 2}},
	"MD": {{"c", 2}, {"c", 18}},
	"ME": {{"n", 3}, {"n", 13}, {"n", 2}},
	"MK": {{"n", 3}, {"c", 10}, {"n", 2}},
	"MR": {{"n", 5}, {"n", 5}, {"n", 11}, {"n", 2}},
	"MT": {{"a", 4}, {"n", 5}, {"c", 18}},
	"MU": {{"a", 4}, {"n", 2}, {"n", 2}, {"n", 12}, {"n", 3}, {"a", 3}},
	"NL": {{"a", 4}, {"n", 10}},
	"NO": {{"n", 4}, {"n", 6}, {"n", 1}},
	"PK": {{"a", 4}, {"c", 16}},
	"PL": {{"n", 8}, {"n", 16}},
	"PS": {{"a", 4}, {"c", 21}},
	"PT": {{"n", 4}, {"n", 4}, {"n", 11}, {"n", 2}},
	"RO": {{"a", 4}, {"c", 16}},
	"RS": {{"n", 3}, {"n", 13}, {"n", 2}},
	"SA": {{"n", 2}, {"c", 18}},
	"SE": {{"n", 3}, {"n", 16}, {"n", 1}},
	"SI": {{"n", 5}, {"n", 8}, {"n", 2}},
	"SK": {{"n", 4}, {"n", 6}, {"n", 10}},
	"SM": {{"a", 1}, {"n", 5}, {"n", 5}, {"c", 12}},
	"TN": {{"n", 2}, {"n", 3}, {"n", 13}, {"n", 2}},
	"TR": {{"n", 5}, {"n", 1}, {"c", 16}},
	"VG": {{"a", 4}, {"n", 16}},
}
