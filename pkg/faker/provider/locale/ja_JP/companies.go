package ja_JP

import "github.com/ensoria/gofake/pkg/faker/provider"

func CreateCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames:    CompanyNames,
		CompanyPrefixes: CompanyPrefixes,
		CompanySuffixes: CompanySuffixes,
		CompanyFormats:  CompanyFormats,
		CreateCompany:   CreateJaJPCompany,
	}
}

var CompanyNames = LastNames
var CompanyPrefixes = []string{"株式会社", "有限会社"}
var CompanySuffixes = CompanyPrefixes

var CompanyFormats = []string{
	"{{.CompanyPrefix}} {{.CompanyName}}",
	"{{.CompanyName}} {{.CompanySuffix}}",
}

type JaJPCompany struct {
	CompanyName   string
	CompanyPrefix string
	CompanySuffix string
}

type JaJPCompanyGenerator interface {
	CompanyName() string
	CompanySuffix() string
	CompanyPrefix() string
}

func CreateJaJPCompany(c any) any {
	g := c.(JaJPCompanyGenerator)
	return &JaJPCompany{
		CompanyName:   g.CompanyName(),
		CompanyPrefix: g.CompanyPrefix(),
		CompanySuffix: g.CompanySuffix(),
	}
}
