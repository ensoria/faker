package core_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
)

var _ = Describe("Tests for random slice functions", func() {

	randSlice := core.NewRandSlice(util.RandSeed())
	Describe("StrElem", func() {
		slice := []string{"abc", "efg", "hij", "klm"}
		It("should return a random string element from the slice", func() {
			r := randSlice.StrElem(slice)
			Expect(r).To(BeElementOf(slice))
		})
	})

	Describe("IntElem", func() {
		slice := []int{123, 456, 789, 101, 102, 103}
		It("should return a random int element from the slice", func() {
			r := randSlice.IntElem(slice)
			Expect(r).To(BeElementOf(slice))
		})
	})
})
