package core_test

import (
	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests for random string functions", func() {
	randStr := core.NewRandStr(util.RandSeed())

	Describe("Char", func() {
		It("should return a single char includes alpha, num, special", func() {
			r := randStr.Char()

			Expect(r).ToNot(BeEmpty())
			Expect(len(r)).To(Equal(1))

			// check output
			testutil.Output("RandStr.Char", r)
		})
	})

	Describe("Letter", func() {
		It("should return a single letter", func() {
			r := randStr.Letter()

			Expect(r).ToNot(BeEmpty())
			Expect(len(r)).To(Equal(1))

			// check output
			testutil.Output("RandStr.Letter", r)
		})
	})

	Describe("Digit", func() {
		It("should return a single digit", func() {
			r := randStr.Digit()

			Expect(r).ToNot(BeEmpty())
			Expect(len(r)).To(Equal(1))

			// check output
			testutil.Output("RandStr.Digit", r)
		})
	})

	Describe("AlphaRange", func() {
		It("should return a string with random length", func() {
			r := randStr.AlphaRange(1, 5)
			length := len(r)

			Expect(length).To(BeNumerically(">", 0))
			Expect(length).To(BeNumerically("<=", 5))

			// check output
			testutil.Output("RandStr.AlphaRange", r)
		})
	})

	Describe("AlphaFixedLength", func() {
		It("should return a string with fixed length", func() {
			r := randStr.AlphaFixedLength(5)
			length := len(r)

			Expect(length).To(Equal(5))

			// check output
			testutil.Output("RandStr.AlphaFixedLength", r)
		})
	})
	Describe("AlphaDigitsLike", func() {
		When("like is 'abc-###-???'", func() {
			It("should return a string with specified alpha and digits", func() {
				r := randStr.AlphaDigitsLike("abc-###-???")
				Expect(r).To(MatchRegexp("abc-[0-9]{3}-[a-zA-Z]{3}"))
			})
		})
		When("like is '***'", func() {
			It("should return a string with specified alpha and digits", func() {
				r := randStr.AlphaDigitsLike("***")
				Expect(r).To(MatchRegexp(`[\d\w]{3}`))

				testutil.Output("RandStr.AlphaDigitsLike", r)
			})
		})

	})
})
