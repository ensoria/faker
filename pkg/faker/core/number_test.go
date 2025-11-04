package core_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
)

var _ = Describe("Tests for random number functions", func() {
	randNum := core.NewRandNum(util.RandSeed())
	Describe("IntBt", func() {
		It("should return int value in given range", func() {
			r := randNum.IntBt(1, 10)
			Expect(r).To(BeNumerically(">=", 1))
			Expect(r).To(BeNumerically("<=", 10))
		})
	})

	Describe("Float32Bt", func() {
		It("should return float value in given range", func() {
			r := randNum.Float32Bt(1.0, 2.0)
			Expect(r).To(BeNumerically(">", 1.0))
			Expect(r).To(BeNumerically("<", 2.0))
		})
	})
})
