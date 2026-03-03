package payment

import (
	"fmt"
	"strings"
	"time"

	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

const (
	defaultExpirationDateFormat = "01/06" // Go time format: MM/YY
	defaultCardSeparator        = "-"
	defaultIbanLength           = 24
	expirationValidMonths       = 36
	upperLetters                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// CreditCardDetails holds credit card information.
type CreditCardDetails struct {
	Type           string
	Number         string
	Name           string
	ExpirationDate string
}

type Payment struct {
	rand *core.Rand
	data *provider.Payments
}

func New(rand *core.Rand, global *provider.Global) *Payment {
	return &Payment{
		rand,
		global.Payments,
	}
}

// CreditCardType returns a random credit card vendor name.
// Example: "MasterCard"
func (p *Payment) CreditCardType() string {
	return p.rand.Slice.StrElem(p.data.CardVendors)
}

// CreditCardNumber returns a credit card number string.
// cardType supports "Visa", "MasterCard", "American Express", "Discover Card", "JCB", and "Visa Retired".
// If cardType is empty, a random type is selected.
// Example: "4485480221084675"
func (p *Payment) CreditCardNumber(cardType string) string {
	if cardType == "" {
		cardType = p.CreditCardType()
	}

	masks, ok := p.data.CardParams[cardType]
	if !ok {
		return ""
	}

	mask := p.rand.Slice.StrElem(masks)
	number := p.rand.Str.AlphaDigitsLike(mask)
	number += ComputeLuhnCheckDigit(number)

	return number
}

// CreditCardNumberFormatted returns a formatted credit card number with separators every 4 digits.
// If cardType is empty, a random type is selected.
// Example: "4485-4802-2108-4675"
func (p *Payment) CreditCardNumberFormatted(cardType string, separator string) string {
	number := p.CreditCardNumber(cardType)
	if number == "" {
		return ""
	}

	if separator == "" {
		separator = defaultCardSeparator
	}

	var parts []string
	for i := 0; i < len(number); i += 4 {
		end := i + 4
		if end > len(number) {
			end = len(number)
		}
		parts = append(parts, number[i:end])
	}

	return strings.Join(parts, separator)
}

// CreditCardExpirationDate returns a random expiration date for a credit card.
// If valid is true, the date will be between now and 36 months in the future.
// If valid is false, the date will be between 36 months ago and 36 months in the future.
func (p *Payment) CreditCardExpirationDate(valid bool) time.Time {
	now := time.Now()
	future := now.AddDate(0, expirationValidMonths, 0)

	if valid {
		return p.rand.Time.TimeRange(now, future)
	}

	past := now.AddDate(0, -expirationValidMonths, 0)
	return p.rand.Time.TimeRange(past, future)
}

// CreditCardExpirationDateString returns a formatted expiration date string.
// If valid is true, the date will be valid (in the future).
// format uses Go time format layout. If empty, defaults to "01/06" (MM/YY).
// Example: "04/26"
func (p *Payment) CreditCardExpirationDateString(valid bool, format string) string {
	if format == "" {
		format = defaultExpirationDateFormat
	}
	return p.CreditCardExpirationDate(valid).Format(format)
}

// CreditCardDetailsResult returns credit card details including type, number, and expiration date.
// The name field is set from the provided holderName parameter.
// If valid is true, the expiration date will be valid.
func (p *Payment) CreditCardDetailsResult(valid bool, holderName string) *CreditCardDetails {
	cardType := p.CreditCardType()

	return &CreditCardDetails{
		Type:           cardType,
		Number:         p.CreditCardNumber(cardType),
		Name:           holderName,
		ExpirationDate: p.CreditCardExpirationDateString(valid, ""),
	}
}

// Iban generates an International Bank Account Number (IBAN).
// countryCode is an ISO 3166-1 alpha-2 country code. If empty, a random country is selected.
// prefix is an optional prefix for the bank account number portion.
// See: http://en.wikipedia.org/wiki/International_Bank_Account_Number
func (p *Payment) Iban(countryCode string, prefix string) string {
	if countryCode == "" {
		countryCode, _ = core.GetRandomKeyValue(p.rand.Map, p.data.IbanFormats)
	} else {
		countryCode = strings.ToUpper(countryCode)
	}

	format, hasFormat := p.data.IbanFormats[countryCode]

	// Calculate total length from format
	length := 0
	if !hasFormat {
		length = defaultIbanLength
		format = [][2]any{{"n", defaultIbanLength}}
	} else {
		for _, part := range format {
			count := part[1].(int)
			length += count
		}
	}

	// Expand format into a string of character classes
	expandedFormat := p.expandIbanFormat(format)

	// Generate result starting with prefix
	result := prefix
	remaining := expandedFormat[len(result):]

	for _, class := range remaining {
		switch class {
		case 'a':
			result += p.randomUpperLetter()
		case 'n':
			result += p.rand.Str.Digit()
		default: // 'c' - alphanumeric
			if p.rand.Bool.Evenly() {
				result += p.rand.Str.Digit()
			} else {
				result += p.randomUpperLetter()
			}
		}
	}

	checksum := CalcIbanChecksum(countryCode + "00" + result)

	return countryCode + checksum + result
}

// SwiftBicNumber generates a random SWIFT/BIC number.
// See: http://en.wikipedia.org/wiki/ISO_9362
// Example: "RZTIAT22263"
func (p *Payment) SwiftBicNumber() string {
	// SWIFT/BIC format: 4 letters (bank) + 2 letters (country) + 2 alphanumeric (location) + optional 3 alphanumeric (branch)
	var result strings.Builder

	// 4 uppercase letters (bank code)
	for i := 0; i < 4; i++ {
		result.WriteString(p.randomUpperLetter())
	}

	// 2 uppercase letters (country code)
	for i := 0; i < 2; i++ {
		result.WriteString(p.randomUpperLetter())
	}

	// 2 alphanumeric (location code)
	for i := 0; i < 2; i++ {
		if p.rand.Bool.Evenly() {
			result.WriteString(p.rand.Str.Digit())
		} else {
			result.WriteString(p.randomUpperLetter())
		}
	}

	// Optional 3 alphanumeric (branch code) - 50% chance
	if p.rand.Bool.Evenly() {
		for i := 0; i < 3; i++ {
			if p.rand.Bool.Evenly() {
				result.WriteString(p.rand.Str.Digit())
			} else {
				result.WriteString(p.randomUpperLetter())
			}
		}
	}

	return result.String()
}

// expandIbanFormat expands an IBAN format definition into a string of character classes.
func (p *Payment) expandIbanFormat(format [][2]any) string {
	var expanded strings.Builder
	for _, item := range format {
		class := item[0].(string)
		count := item[1].(int)
		expanded.WriteString(strings.Repeat(class, count))
	}
	return expanded.String()
}

// randomUpperLetter returns a random uppercase letter (A-Z).
func (p *Payment) randomUpperLetter() string {
	idx := p.rand.Num.Intn(len(upperLetters))
	return fmt.Sprintf("%c", upperLetters[idx])
}
