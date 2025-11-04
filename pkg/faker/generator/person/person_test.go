package person_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/person"
	"github.com/ensoria/faker/pkg/faker/provider/locale/en_US"
	"github.com/ensoria/faker/pkg/faker/provider/locale/ja_JP"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Person", func() {
	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	p := person.New(coreRand, localized)

	Describe("FirstNameMale", func() {
		It("should return a male first name", func() {
			r := p.FirstNameMale()
			Expect(r).To(BeElementOf(en_US.FirstNameMales))
		})
	})

	Describe("FirstNameFemale", func() {
		It("should return a female first name", func() {
			r := p.FirstNameFemale()
			Expect(r).To(BeElementOf(en_US.FirstNameFemales))
		})
	})

	Describe("FirstName", func() {
		It("should return a first name", func() {
			r := p.FirstName()

			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FirstName", r)
		})
	})

	Describe("LastName", func() {
		It("should return a last name", func() {
			r := p.LastName()
			Expect(r).To(BeElementOf(en_US.LastNames))
		})
	})

	Describe("TitleMale", func() {
		It("should return a male title", func() {
			r := p.TitleMale()
			Expect(r).To(BeElementOf(en_US.TitleMales))
		})
	})

	Describe("TitleFemale", func() {
		It("should return a female title", func() {
			r := p.TitleFemale()
			Expect(r).To(BeElementOf(en_US.TitleFemales))
		})
	})

	Describe("Title", func() {
		It("should return a title", func() {
			r := p.Title()

			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.Title", r)
		})
	})

	Describe("Suffix", func() {
		It("should return a suffix", func() {
			r := p.Suffix()
			Expect(r).To(BeElementOf(en_US.Suffixes))
		})
	})

	Describe("MaleName", func() {
		It("should return a male name", func() {
			r := p.MaleName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.MaleName", r)
		})
	})

	Describe("FemaleName", func() {
		It("should return a female name", func() {
			r := p.FemaleName()

			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FemaleName", r)
		})
	})

	Describe("Name", func() {
		It("should return a person's full name", func() {
			r := p.Name()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.Name", r)
		})
	})

	Describe("Ssn", func() {
		It("should return a person's SSN", func() {
			r := p.Ssn()
			Expect(r).To(MatchRegexp(`^\d{3}-\d{2}-\d{4}$`))
		})
	})

	jaJp := ja_JP.New()
	pJaJp := person.New(coreRand, jaJp)
	Describe("FirstKanaNameMale", func() {
		It("should return a male first kana name", func() {
			r := pJaJp.FirstKanaNameMale()
			Expect(r).To(BeElementOf(ja_JP.FirstKanaNameMales))
		})
	})

	Describe("FirstKanaNameFemale", func() {
		It("should return a female first kana name", func() {
			r := pJaJp.FirstKanaNameFemale()
			Expect(r).To(BeElementOf(ja_JP.FirstKanaNameFemales))
		})
	})

	Describe("FirstKanaName", func() {
		It("should return a first kana name", func() {
			r := pJaJp.FirstKanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FirstKanaName", r)
		})
	})

	Describe("LastKanaName", func() {
		It("should return a last kana name", func() {
			r := pJaJp.LastKanaName()
			Expect(r).To(BeElementOf(ja_JP.LastKanaNames))
		})
	})

	Describe("MaleKanaName", func() {
		It("should return a male kana name", func() {
			r := pJaJp.MaleKanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.MaleKanaName", r)
		})
	})

	Describe("FemaleKanaName", func() {
		It("should return a female kana name", func() {
			r := pJaJp.FemaleKanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FemaleKanaName", r)
		})
	})

	Describe("KanaName", func() {
		It("should return a kana name", func() {
			r := pJaJp.KanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.KanaName", r)
		})
	})
})
