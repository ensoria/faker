package medical_test

import (
	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/medical"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
	"github.com/ensoria/faker/pkg/faker/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Medical", func() {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Medicals: global.CreateMedicals(),
	}

	md := medical.New(coreRand, global)

	It("BloodType should return a blood type", func() {
		r := md.BloodType()

		Expect(r).To(BeElementOf(global.Medicals.BloodTypes))
		testutil.Output("Medical.BloodType", r)
	})

	It("BloodRhFactor should return a blood rh factor", func() {
		r := md.BloodRhFactor()

		Expect(r).To(BeElementOf(global.Medicals.BloodRhFactors))
		testutil.Output("Medical.BloodRhFactor", r)
	})

	It("BloodGroup should return a blood group", func() {
		r := md.BloodGroup()
		Expect(r).To(MatchRegexp(`^(A|B|AB|O)(\+|\-)$`))
		testutil.Output("Medical.BloodGroup", r)
	})
})
