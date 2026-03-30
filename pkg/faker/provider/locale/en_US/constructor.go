package en_US

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
)

func New() *provider.Localized {
	return &provider.Localized{
		People:       CreatePeople(),
		Addresses:    CreateAddresses(),
		Companies:    CreateCompanies(),
		PhoneNumbers: CreatePhoneNumbers(),
	}
}
