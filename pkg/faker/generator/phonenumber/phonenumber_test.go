package phonenumber_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/payment"
	"github.com/ensoria/gofake/pkg/faker/generator/phonenumber"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/global"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_US"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/ja_JP"
	"github.com/ensoria/gofake/pkg/faker/testutil"
)

var _ = Describe("PhoneNumber", func() {
	coreRand := core.NewRand(util.RandSeed())
	g := &provider.Global{
		PhoneNumbers: global.CreatePhoneNumbers(),
	}

	Describe("with en_US locale", func() {
		l := &provider.Localized{
			PhoneNumbers: en_US.CreatePhoneNumbers(),
		}
		p := phonenumber.New(coreRand, g, l)

		It("PhoneNumber should return a US-formatted phone number", func() {
			r := p.PhoneNumber()
			Expect(r).NotTo(BeEmpty())
			// Should contain only digits, dashes, dots, parentheses, spaces, and 'x'
			Expect(r).To(MatchRegexp(`^[\d\-\(\)\. x]+$`))
			testutil.Output("PhoneNumber.PhoneNumber (en_US)", r)
		})
	})

	Describe("with ja_JP locale", func() {
		l := &provider.Localized{
			PhoneNumbers: ja_JP.CreatePhoneNumbers(),
		}
		p := phonenumber.New(coreRand, g, l)

		It("PhoneNumber should return a JP-formatted phone number", func() {
			r := p.PhoneNumber()
			Expect(r).NotTo(BeEmpty())
			// Should start with 0 and contain digits and dashes
			Expect(r).To(MatchRegexp(`^0[\d\-]+$`))
			testutil.Output("PhoneNumber.PhoneNumber (ja_JP)", r)
		})
	})

	Describe("E164PhoneNumber", func() {
		l := &provider.Localized{
			PhoneNumbers: en_US.CreatePhoneNumbers(),
		}
		p := phonenumber.New(coreRand, g, l)

		It("should return a phone number starting with +", func() {
			r := p.E164PhoneNumber()
			Expect(r).To(MatchRegexp(`^\+\d+$`))
			// E.164 numbers are max 15 digits (including country code, excluding +)
			digits := r[1:] // strip leading +
			Expect(len(digits)).To(BeNumerically(">=", 8))
			Expect(len(digits)).To(BeNumerically("<=", 15))
			testutil.Output("PhoneNumber.E164PhoneNumber", r)
		})
	})

	Describe("IMEI", func() {
		l := &provider.Localized{
			PhoneNumbers: en_US.CreatePhoneNumbers(),
		}
		p := phonenumber.New(coreRand, g, l)

		It("should return a 15-digit Luhn-valid number", func() {
			r := p.IMEI()
			Expect(r).To(HaveLen(15))
			Expect(r).To(MatchRegexp(`^\d{15}$`))
			Expect(payment.IsLuhnValid(r)).To(BeTrue())
			testutil.Output("PhoneNumber.IMEI", r)
		})
	})
})
