package payment_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/generator/payment"
)

var _ = Describe("Iban", func() {
	Describe("CalcIbanChecksum", func() {
		It("should compute correct IBAN checksums", func() {
			testCases := []struct {
				iban string
				want string
			}{
				// GB82WEST12345698765432 is a well-known valid IBAN
				{"GB00WEST12345698765432", "82"},
				// DE89370400440532013000 is a well-known valid German IBAN
				{"DE00370400440532013000", "89"},
			}

			for _, tc := range testCases {
				Expect(payment.CalcIbanChecksum(tc.iban)).To(Equal(tc.want))
			}
		})
	})

	Describe("IsIbanValid", func() {
		It("should return true for valid IBANs", func() {
			validIbans := []string{
				"GB82WEST12345698765432",
				"DE89370400440532013000",
				"FR7630006000011234567890189",
				"ES9121000418450200051332",
			}
			for _, iban := range validIbans {
				Expect(payment.IsIbanValid(iban)).To(BeTrue(), "expected %s to be valid", iban)
			}
		})

		It("should return false for invalid IBANs", func() {
			invalidIbans := []string{
				"GB00WEST12345698765432",
				"DE00370400440532013000",
			}
			for _, iban := range invalidIbans {
				Expect(payment.IsIbanValid(iban)).To(BeFalse(), "expected %s to be invalid", iban)
			}
		})
	})
})
