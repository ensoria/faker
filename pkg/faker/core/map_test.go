package core_test

import (
	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests for random map functions", func() {
	randMap := core.NewRandMap(util.RandSeed())

	mockMap1 := map[any]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	mockMap2 := map[any][]any{
		1: {"value11", "value12"},
		2: {"value21", "value22"},
	}

	Describe("GetRandomKeyValue", func() {
		It("should return string key and string value", func() {
			kStr, vStr := core.GetRandomKeyValue(randMap, mockMap1)
			vExpected := mockMap1[kStr]
			Expect(vStr).To(Equal(vExpected))
		})

		It("should return int key and slice value", func() {
			kInt, vSlice := core.GetRandomKeyValue(randMap, mockMap2)
			vExpected2 := mockMap2[kInt]
			Expect(vSlice).To(Equal(vExpected2))
		})
	})

	Describe("KeyValue", func() {
		It("should return key and the value", func() {
			k, v := randMap.KeyValue(mockMap1)

			vExpected := mockMap1[k]
			Expect(v).To(Equal(vExpected))
		})
	})

	Describe("KeySliceValue", func() {
		It("should return key and the value", func() {
			k, v := randMap.KeySliceValue(mockMap2)

			vExpected := mockMap2[k]
			Expect(v).To(Equal(vExpected))
		})
	})

})
