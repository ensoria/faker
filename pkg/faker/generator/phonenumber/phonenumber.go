package phonenumber

import (
	"bytes"

	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/payment"
	"github.com/ensoria/faker/pkg/faker/provider"
)

const (
	imeiDigitCount = 14 // 14 random digits + 1 Luhn check digit = 15
)

type PhoneNumber struct {
	rand       *core.Rand
	globalData *provider.PhoneNumbers
	localData  *provider.PhoneNumbers
}

func New(rand *core.Rand, global *provider.Global, localized *provider.Localized) *PhoneNumber {
	return &PhoneNumber{
		rand:       rand,
		globalData: global.PhoneNumbers,
		localData:  localized.PhoneNumbers,
	}
}

// PhoneNumber generates a locale-specific phone number.
// Example (en_US): "201-886-0269"
// Example (ja_JP): "090-1234-5678"
func (p *PhoneNumber) PhoneNumber() string {
	format := p.rand.Slice.StrElem(p.localData.Formats)
	return p.rand.Str.AlphaDigitsLike(format)
}

// E164PhoneNumber generates an E.164 compliant international phone number.
// Example: "+27113456789"
func (p *PhoneNumber) E164PhoneNumber() string {
	format := p.rand.Slice.StrElem(p.globalData.E164Formats)
	return p.rand.Str.AlphaDigitsLike(format)
}

// IMEI generates a valid IMEI (International Mobile Equipment Identity) number.
// The IMEI is 15 digits, with the last digit being a Luhn check digit.
// See: http://en.wikipedia.org/wiki/International_Mobile_Station_Equipment_Identity
// Example: "354809024498147"
func (p *PhoneNumber) IMEI() string {
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, imeiDigitCount))
	code := p.rand.Str.AlphaDigitsLike(format)
	return code + payment.ComputeLuhnCheckDigit(code)
}
