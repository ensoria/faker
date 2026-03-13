package company

import (
	"fmt"

	"github.com/ensoria/faker/pkg/faker/common/log"
	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

// Company provides methods for generating random company data.
//
// ランダムな会社データを生成するメソッドを提供する構造体。
type Company struct {
	rand *core.Rand
	data *provider.Companies
}

// New creates a new Company instance with the given random source and locale data.
//
// 指定されたランダムソースとロケールデータで新しいCompanyインスタンスを作成する。
func New(rand *core.Rand, localized *provider.Localized) *Company {
	return &Company{
		rand,
		localized.Companies,
	}
}

// CompanyName returns a random company name from the locale data.
//
// ロケールデータからランダムな会社名を返す。
func (c *Company) CompanyName() string {
	if len(c.data.CompanyNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.CompanyNames)

}

// CompanyPrefix returns a random company name prefix.
//
// ランダムな会社名の接頭辞を返す。
func (c *Company) CompanyPrefix() string {
	if len(c.data.CompanyPrefixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.CompanyPrefixes)
}

// CompanySuffix returns a random company name suffix.
//
// ランダムな会社名の接尾辞を返す。
func (c *Company) CompanySuffix() string {
	if len(c.data.CompanySuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.CompanySuffixes)
}

// Name returns a randomly formatted company name using locale-specific formats.
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた会社名を返す。
func (c *Company) Name() string {
	if len(c.data.CompanyFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := c.rand.Slice.StrElem(c.data.CompanyFormats)
	nameData := c.data.CreateCompany(c)
	return util.RenderTemplate(format, nameData)

}

// JobTitleName returns a random job title name.
//
// ランダムな役職名を返す。
func (c *Company) JobTitleName() string {
	if len(c.data.JobTitleNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.JobTitleNames)
}

// JobTitle returns a randomly formatted job title using locale-specific formats.
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた役職を返す。
func (c *Company) JobTitle() string {
	if len(c.data.JobTitleFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := c.rand.Slice.StrElem(c.data.JobTitleFormats)
	nameData := c.data.CreateJobTitle(c)
	return util.RenderTemplate(format, nameData)
}

// EINPrefix returns a random EIN (Employer Identification Number) prefix.
//
// ランダムなEIN（雇用者識別番号）の接頭辞を返す。
func (c *Company) EINPrefix() string {
	if len(c.data.EINPrefixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.EINPrefixes)
}

// EIN returns a random Employer Identification Number.
// See: https://en.wikipedia.org/wiki/Employer_Identification_Number
// Example: "12-3456789"
//
// ランダムな雇用者識別番号を返す。
func (c *Company) EIN() string {
	prefix := c.EINPrefix()
	suffixNum := c.rand.Num.Int64Bt(0, 9999999)
	return prefix + "-" + fmt.Sprintf("%07d", suffixNum)
}
