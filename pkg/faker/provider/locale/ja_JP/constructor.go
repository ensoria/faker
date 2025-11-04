package ja_JP

import (
	"github.com/ensoria/faker/pkg/faker/provider"
)

func New() *provider.Localized {
	return &provider.Localized{
		People:    CreatePeople(),
		Addresses: CreateAddresses(),
		Companies: CreateCompanies(),
	}
}
