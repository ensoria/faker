package global

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_US"
)

func CreateInternets() *provider.Internets {
	return &provider.Internets{
		// username
		FirstNames:      FirstNames,
		LastNames:       LastNames,
		UserNameFormats: UserNameFormats,
		CreateUserName:  CreateUserName,
		// email
		EmailFormats: EmailFormats,
		CreateEmail:  CreateEmail,
		// url
		TLD:           TLD,
		URLFormats:    URLFormats,
		LocalIPBlocks: LocalIPBlocks,
	}
}

// UserName
var LastNames = en_US.LastNames

var FirstNames = append(en_US.FirstNameMales, en_US.FirstNameFemales...)
var UserNameFormats = []string{
	"{{.LastName}}.{{.FirstName}}",
	"{{.FirstName}}.{{.LastName}}",
	"{{.FirstName}}##",
	"{{.LastName}}##",
}

type UserName struct {
	FirstName string
	LastName  string
}

type UserNameGenerator interface {
	FirstName() string
	LastName() string
}

func CreateUserName(i any) any {
	g := i.(UserNameGenerator)
	return &UserName{
		g.FirstName(),
		g.LastName(),
	}
}

var TLD = []string{
	"com", "biz", "info", "net", "org",
	"edu", "gov", "mil", "int", "co",
	"io", "dev", "app", "ai", "me",
	"tv", "cc", "us", "uk", "de",
	"fr", "jp", "cn", "au", "ca",
	"br", "in", "ru", "nl", "se",
	"no", "fi", "dk", "ch", "at",
	"be", "es", "it", "pt", "pl",
	"cz", "kr", "mx", "za", "nz",
	"ie", "sg", "hk", "tw", "cloud",
	"xyz", "online", "site", "tech", "store",
}

var EmailFormats = []string{
	"{{.UserName}}@{{.DomainName}}",
}

type Email struct {
	UserName   string
	DomainName string
}

type EmailGenerator interface {
	UserName() string
	DomainName() string
}

func CreateEmail(i any) any {
	g := i.(EmailGenerator)
	return &Email{
		UserName:   g.UserName(),
		DomainName: g.DomainName(),
	}
}

var URLFormats = []string{
	"http://www.{{.DomainName}}/",
	"http://{{.DomainName}}/",
	"http://www.{{.DomainName}}/{{.Slug}}",
	"http://www.{{.DomainName}}/{{.Slug}}",
	"https://www.{{.DomainName}}/{{.Slug}}",
	"http://www.{{.DomainName}}/{{.Slug}}.html",
	"http://{{.DomainName}}/{{.Slug}}",
	"http://{{.DomainName}}/{{.Slug}}",
	"http://{{.DomainName}}/{{.Slug}}.html",
	"https://{{.DomainName}}/{{.Slug}}.html",
}

var LocalIPBlocks = [][]string{
	{"10.0.0.0", "10.255.255.255"},
	{"172.16.0.0", "172.31.255.255"},
	{"192.168.0.0", "192.168.255.255"},
}
