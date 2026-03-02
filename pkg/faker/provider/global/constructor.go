package global

import "github.com/ensoria/faker/pkg/faker/provider"

func New() *provider.Global {
	return &provider.Global{
		Colors:    CreateColors(),
		Files:     CreateFiles(),
		Images:    CreateImages(),
		Internets: CreateInternets(),
		Lorems:    CreateLorems(),
		Medicals:  CreateMedicals(),
	}
}
