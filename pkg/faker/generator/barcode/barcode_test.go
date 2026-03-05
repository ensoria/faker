package barcode_test

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/barcode"
)

var _ = Describe("Barcode", func() {

	coreRand := core.NewRand(util.RandSeed())
	bc := barcode.New(coreRand)

	Describe("EAN", func() {
		It("EAN8 should return a barcode with 8 digits", func() {
			r := bc.EAN8()
			Expect(r).To(MatchRegexp(`^\d{8}$`))
			lastDigit, _ := strconv.Atoi(r[7:])
			heading7Digits := r[:7]
			Expect(lastDigit).To(Equal(barcode.CalcEANCheckDigit(heading7Digits)))
		})

		It("EAN13 should return a barcode with 13 digits", func() {
			r := bc.EAN13()
			Expect(r).To(MatchRegexp(`^\d{13}$`))
			lastDigit, _ := strconv.Atoi(r[12:])
			heading12Digits := r[:12]
			Expect(lastDigit).To(Equal(barcode.CalcEANCheckDigit(heading12Digits)))
		})

		It("CalcEANCheckDigit should calculate barcode check digit", func() {
			testCases := []struct {
				digits string
				want   int
			}{
				// 7 digits -> EAN-8
				{"1234567", 0},
				{"7654321", 0},
				{"3897546", 2},
				{"7653573", 4},
				{"3264902", 4},
				// 12 digits -> EAN-13
				{"764564239870", 7},
				{"233209246782", 8},
				{"876456876876", 6},
				{"272549071238", 4},
				{"986126758742", 7},
			}
			for _, tc := range testCases {
				result := barcode.CalcEANCheckDigit(tc.digits)
				Expect(result).To(Equal(tc.want))
			}
		})
	})

	Describe("ISBN", func() {
		It("ISBN10 should return a ISBN with 10 digits", func() {
			r := bc.ISBN10()
			Expect(r).To(MatchRegexp(`^\d{9}[\dX]$`))

			lastDigit := r[9:]
			heading9Digits := r[:9]
			Expect(lastDigit).To(Equal(barcode.CalcISBNCheckDigit(heading9Digits)))
		})

		It("ISBN13 should return a ISBN with 13 digits", func() {
			r := bc.ISBN13()
			Expect(r).To(MatchRegexp(`^97[89]\d{10}$`))

			lastDigit, _ := strconv.Atoi(r[12:])
			heading12Digits := r[:12]
			Expect(lastDigit).To(Equal(barcode.CalcEANCheckDigit(heading12Digits)))
		})

		It("CalcISBNCheckDigit should calculate ISBN check digit", func() {
			testCases := []struct {
				input string
				want  string
			}{
				{"765757657", "X"},
				{"111111111", "1"},
				{"764350980", "8"},
				{"325468122", "3"},
				{"753697616", "X"},
				{"456893287", "4"},
			}

			for _, tc := range testCases {
				got := barcode.CalcISBNCheckDigit(tc.input)
				Expect(got).To(Equal(tc.want))
			}
		})
	})
})
