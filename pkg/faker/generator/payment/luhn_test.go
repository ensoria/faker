package payment_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/generator/payment"
)

var _ = Describe("Luhn", func() {
	Describe("CalcLuhnChecksum", func() {
		It("should return 0 for a valid Luhn number", func() {
			// 79927398713 is a well-known valid Luhn number
			Expect(payment.CalcLuhnChecksum("79927398713")).To(Equal(0))
		})

		It("should return non-zero for an invalid number", func() {
			Expect(payment.CalcLuhnChecksum("79927398710")).NotTo(Equal(0))
		})
	})

	Describe("ComputeLuhnCheckDigit", func() {
		It("should compute correct check digits", func() {
			testCases := []struct {
				partial string
				want    string
			}{
				{"7992739871", "3"},
				{"0", "0"},
				{"1", "8"},
				{"12", "5"},
				{"400000000000000", "2"},
				{"453200000000000", "9"},
			}

			for _, tc := range testCases {
				Expect(payment.ComputeLuhnCheckDigit(tc.partial)).To(Equal(tc.want))
			}
		})
	})

	Describe("IsLuhnValid", func() {
		It("should return true for valid Luhn numbers", func() {
			validNumbers := []string{
				"79927398713",
				"4532000000000009",
				"4000000000000002",
			}
			for _, n := range validNumbers {
				Expect(payment.IsLuhnValid(n)).To(BeTrue(), "expected %s to be valid", n)
			}
		})

		It("should return false for invalid Luhn numbers", func() {
			invalidNumbers := []string{
				"79927398710",
				"1234567890",
			}
			for _, n := range invalidNumbers {
				Expect(payment.IsLuhnValid(n)).To(BeFalse(), "expected %s to be invalid", n)
			}
		})
	})
})
