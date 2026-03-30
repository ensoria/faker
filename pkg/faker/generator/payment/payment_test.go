package payment_test

import (
	"regexp"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/payment"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/global"
	"github.com/ensoria/gofake/pkg/faker/testutil"
)

var _ = Describe("Payment", func() {
	coreRand := core.NewRand(util.RandSeed())
	g := &provider.Global{
		Payments: global.CreatePayments(),
	}
	p := payment.New(coreRand, g)

	Describe("CreditCardType", func() {
		It("should return a valid card vendor", func() {
			r := p.CreditCardType()
			Expect(r).To(BeElementOf(g.Payments.CardVendors))
			testutil.Output("Payment.CreditCardType", r)
		})
	})

	Describe("CreditCardNumber", func() {
		It("should return a Luhn-valid number when type is not specified", func() {
			r := p.CreditCardNumber("")
			Expect(r).To(MatchRegexp(`^\d+$`))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			testutil.Output("Payment.CreditCardNumber (random)", r)
		})

		It("should return a valid Visa number", func() {
			r := p.CreditCardNumber("Visa")
			Expect(r).To(HavePrefix("4"))
			Expect(r).To(HaveLen(16))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			testutil.Output("Payment.CreditCardNumber (Visa)", r)
		})

		It("should return a valid Visa Retired number", func() {
			r := p.CreditCardNumber("Visa Retired")
			Expect(r).To(HavePrefix("4"))
			Expect(r).To(HaveLen(13))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			testutil.Output("Payment.CreditCardNumber (Visa Retired)", r)
		})

		It("should return a valid MasterCard number", func() {
			r := p.CreditCardNumber("MasterCard")
			Expect(r).To(HaveLen(16))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			// MasterCard starts with 2221-2720 or 51-55
			first2 := r[:2]
			validStart := (first2 >= "22" && first2 <= "27") || (first2 >= "51" && first2 <= "55")
			Expect(validStart).To(BeTrue())
			testutil.Output("Payment.CreditCardNumber (MasterCard)", r)
		})

		It("should return a valid American Express number", func() {
			r := p.CreditCardNumber("American Express")
			Expect(r).To(HaveLen(15))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			Expect(r[:2]).To(SatisfyAny(Equal("34"), Equal("37")))
			testutil.Output("Payment.CreditCardNumber (Amex)", r)
		})

		It("should return a valid JCB number", func() {
			r := p.CreditCardNumber("JCB")
			Expect(r).To(HaveLen(16))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			Expect(r[:4]).To(SatisfyAny(Equal("3528"), Equal("3589")))
			testutil.Output("Payment.CreditCardNumber (JCB)", r)
		})

		It("should return a valid Discover Card number", func() {
			r := p.CreditCardNumber("Discover Card")
			Expect(r).To(HaveLen(16))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			Expect(r[:4]).To(Equal("6011"))
			testutil.Output("Payment.CreditCardNumber (Discover)", r)
		})

		It("should return empty string for unknown type", func() {
			r := p.CreditCardNumber("Unknown")
			Expect(r).To(BeEmpty())
		})
	})

	Describe("CreditCardNumberFormatted", func() {
		It("should return a formatted card number with default separator", func() {
			r := p.CreditCardNumberFormatted("Visa", "")
			Expect(r).To(MatchRegexp(`^\d{4}-\d{4}-\d{4}-\d{4}$`))
			// Verify Luhn validity of the underlying number
			number := strings.ReplaceAll(r, "-", "")
			Expect(payment.IsLuhnValid(number)).To(BeTrue())
			testutil.Output("Payment.CreditCardNumberFormatted", r)
		})

		It("should return a formatted card number with custom separator", func() {
			r := p.CreditCardNumberFormatted("Visa", " ")
			Expect(r).To(MatchRegexp(`^\d{4} \d{4} \d{4} \d{4}$`))
			testutil.Output("Payment.CreditCardNumberFormatted (space)", r)
		})
	})

	Describe("CreditCardExpirationDate", func() {
		It("should return a future date when valid is true", func() {
			r := p.CreditCardExpirationDate(true)
			Expect(r).To(BeTemporally(">=", time.Now().Add(-1*time.Second)))
			Expect(r).To(BeTemporally("<=", time.Now().AddDate(0, 36, 0).Add(1*time.Second)))
		})

		It("should return a date within range when valid is false", func() {
			r := p.CreditCardExpirationDate(false)
			Expect(r).To(BeTemporally(">=", time.Now().AddDate(0, -36, 0).Add(-1*time.Second)))
			Expect(r).To(BeTemporally("<=", time.Now().AddDate(0, 36, 0).Add(1*time.Second)))
		})
	})

	Describe("CreditCardExpirationDateString", func() {
		It("should return date in default format (MM/YY)", func() {
			r := p.CreditCardExpirationDateString(true, "")
			Expect(r).To(MatchRegexp(`^\d{2}/\d{2}$`))
			testutil.Output("Payment.CreditCardExpirationDateString", r)
		})

		It("should return date in custom format", func() {
			r := p.CreditCardExpirationDateString(true, "2006-01")
			Expect(r).To(MatchRegexp(`^\d{4}-\d{2}$`))
			testutil.Output("Payment.CreditCardExpirationDateString (custom)", r)
		})
	})

	Describe("CreditCardDetailsResult", func() {
		It("should return complete credit card details", func() {
			r := p.CreditCardDetailsResult(true, "John Doe")
			Expect(r.Type).To(BeElementOf(g.Payments.CardVendors))
			Expect(r.Number).NotTo(BeEmpty())
			Expect(payment.IsLuhnValid(r.Number)).To(BeTrue())
			Expect(r.Name).To(Equal("John Doe"))
			Expect(r.ExpirationDate).To(MatchRegexp(`^\d{2}/\d{2}$`))
			testutil.Output("Payment.CreditCardDetailsResult", r)
		})
	})

	Describe("IBAN", func() {
		It("should generate a valid IBAN with random country code", func() {
			r := p.IBAN("", "")
			Expect(len(r)).To(BeNumerically(">=", 4))
			// First two chars should be uppercase letters (country code)
			Expect(r[:2]).To(MatchRegexp(`^[A-Z]{2}$`))
			// Next two chars should be digits (check digits)
			Expect(r[2:4]).To(MatchRegexp(`^\d{2}$`))
			Expect(payment.IsIBANValid(r)).To(BeTrue())
			testutil.Output("Payment.IBAN (random)", r)
		})

		It("should generate a valid German IBAN", func() {
			r := p.IBAN("DE", "")
			Expect(r).To(HavePrefix("DE"))
			// DE IBAN: DE + 2 check digits + 18 digits = 22 chars
			Expect(r).To(HaveLen(22))
			Expect(payment.IsIBANValid(r)).To(BeTrue())
			testutil.Output("Payment.IBAN (DE)", r)
		})

		It("should generate a valid British IBAN", func() {
			r := p.IBAN("GB", "")
			Expect(r).To(HavePrefix("GB"))
			// GB IBAN: GB + 2 check digits + 4 letters + 14 digits = 22 chars
			Expect(r).To(HaveLen(22))
			Expect(payment.IsIBANValid(r)).To(BeTrue())
			testutil.Output("Payment.IBAN (GB)", r)
		})

		It("should generate a valid IBAN with a prefix", func() {
			r := p.IBAN("DE", "3704")
			Expect(r).To(HavePrefix("DE"))
			// The BBAN should start with the prefix
			Expect(r[4:8]).To(Equal("3704"))
			Expect(payment.IsIBANValid(r)).To(BeTrue())
			testutil.Output("Payment.IBAN (DE, prefix)", r)
		})

		It("should handle lowercase country code", func() {
			r := p.IBAN("de", "")
			Expect(r).To(HavePrefix("DE"))
			Expect(payment.IsIBANValid(r)).To(BeTrue())
		})
	})

	Describe("SWIFTBICNumber", func() {
		It("should return a valid SWIFT/BIC number", func() {
			r := p.SWIFTBICNumber()
			// SWIFT format: 4 letters + 2 letters + 2 alphanumeric + optional 3 alphanumeric
			// Total: 8 or 11 characters
			validPattern := regexp.MustCompile(`^[A-Z]{4}[A-Z]{2}[0-9A-Z]{2}([0-9A-Z]{3})?$`)
			Expect(validPattern.MatchString(r)).To(BeTrue())
			Expect(len(r)).To(SatisfyAny(Equal(8), Equal(11)))
			testutil.Output("Payment.SWIFTBICNumber", r)
		})
	})
})
