package company_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/company"
	"github.com/ensoria/faker/pkg/faker/provider/locale/en_US"
	"github.com/ensoria/faker/pkg/faker/provider/locale/ja_JP"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Company", func() {

	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	comp := company.New(coreRand, localized)

	Describe("Company", func() {
		It("CompanyName should return a company name", func() {
			r := comp.CompanyName()
			Expect(r).To(BeElementOf(en_US.CompanyNames))
		})

		It("CompanySuffix should return a company suffix", func() {
			r := comp.CompanySuffix()
			Expect(r).To(BeElementOf(en_US.CompanySuffixes))
		})

		It("Name should return a company name", func() {
			r := comp.Name()
			testutil.Output("Company.Name", r)
		})
	})

	Describe("JobTitle", func() {
		It("jobTitleName should return a job title name", func() {
			r := comp.JobTitleName()
			Expect(r).To(BeElementOf(en_US.JobTitleNames))
		})

		It("JobTitle should return a job title", func() {
			r := comp.JobTitle()
			testutil.Output("Company.JobTitle", r)
		})
	})

	Describe("EIN", func() {
		It("EinPrefix should return a EIN prefix", func() {
			r := comp.EinPrefix()
			Expect(r).To(BeElementOf(en_US.EinPrefixes))
		})
		It("Ein should return a EIN", func() {
			r := comp.Ein()
			Expect(r).To(MatchRegexp(`\d{2}-\d{7}`))
		})
	})

	jaJp := ja_JP.New()
	compJaJp := company.New(coreRand, jaJp)
	Describe("ja_JP Company", func() {
		It("CompanyPrefix should return a company prefix", func() {
			r := compJaJp.CompanyPrefix()
			Expect(r).To(BeElementOf(ja_JP.CompanyPrefixes))
		})

		It("Name should return a company name", func() {
			r := compJaJp.Name()
			testutil.Output("Company.Name", r)
		})
	})
})
