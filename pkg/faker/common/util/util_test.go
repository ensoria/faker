package util_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("TruncateToPrecision", func() {
	rand := util.RandSeed()
	It("should truncate to precision", func() {
		precision := rand.Intn(8)
		r := util.TruncateToPrecision(1.23456789123456, precision)
		fmt.Println(r)
		decimalLength := testutil.GetDecimalLength(r)
		fmt.Println(decimalLength)

		Expect(decimalLength).To(Equal(precision))
	})
})

var _ = Describe("CapFirstLetter", func() {
	It("should capitalize the first letter", func() {
		r := util.CapFirstLetter("hello")
		Expect(r).To(Equal("Hello"))

		r2 := util.CapFirstLetter("world")
		Expect(r2).To(Equal("World"))
	})
})
