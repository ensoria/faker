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
		It("Ean8 should return a barcode with 8 digits", func() {
			r := bc.Ean8()
			Expect(r).To(MatchRegexp(`^\d{8}$`))
			lastDigit, _ := strconv.Atoi(r[7:])
			heading7Digits := r[:7]
			Expect(lastDigit).To(Equal(barcode.CalcEanCheckDigit(heading7Digits)))
		})

		It("Ean13 should return a barcode with 13 digits", func() {
			r := bc.Ean13()
			Expect(r).To(MatchRegexp(`^\d{13}$`))
			lastDigit, _ := strconv.Atoi(r[12:])
			heading12Digits := r[:12]
			Expect(lastDigit).To(Equal(barcode.CalcEanCheckDigit(heading12Digits)))
		})

		It("CalcEanCheckDigit should calculate barcode check digit", func() {
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
				result := barcode.CalcEanCheckDigit(tc.digits)
				Expect(result).To(Equal(tc.want))
			}
		})
	})

	Describe("ISBN", func() {
		It("Isbn10 should return a ISBN with 10 digits", func() {
			r := bc.Isbn10()
			Expect(r).To(MatchRegexp(`^\d{9}[\dX]$`))

			lastDigit := r[9:]
			heading9Digits := r[:9]
			Expect(lastDigit).To(Equal(barcode.CalcIsbnCheckDigit(heading9Digits)))
		})

		It("Isbn13 should return a ISBN with 13 digits", func() {
			r := bc.Isbn13()
			Expect(r).To(MatchRegexp(`^97[89]\d{10}$`))

			lastDigit, _ := strconv.Atoi(r[12:])
			heading12Digits := r[:12]
			Expect(lastDigit).To(Equal(barcode.CalcEanCheckDigit(heading12Digits)))
		})

		It("CalcIsbnCheckDigit should calculate ISBN check digit", func() {
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
				got := barcode.CalcIsbnCheckDigit(tc.input)
				Expect(got).To(Equal(tc.want))
			}
		})
	})
})
