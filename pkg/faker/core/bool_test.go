package core_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Tests for random bool functions", func() {
	randBool := core.NewRandBool(util.RandSeed())
	Describe("Evenly", func() {
		It("should return true or false", func() {
			r := randBool.Evenly()
			// Check with output
			testutil.Output("RandBool.Evenly", r)
			Expect(r).To(BeAssignableToTypeOf(true))
		})
	})

	Describe("Weighted", func() {
		It("should return weighted bool", func() {
			r := randBool.WeightedToTrue(0.7)
			// Check with output
			testutil.Output("RandBool.WeightedToTrue", r)
			Expect(r).To(BeAssignableToTypeOf(true))
		})

		It("should return true when 1", func() {
			r := randBool.WeightedToTrue(1)
			Expect(r).To(BeTrue())
		})

		It("should return false when 0", func() {
			r := randBool.WeightedToTrue(0)
			Expect(r).To(BeFalse())
		})
	})
})
