package global

import "github.com/ensoria/faker/pkg/faker/provider"

func CreateMedicals() *provider.Medicals {
	return &provider.Medicals{
		BloodTypes:     BloodTypes,
		BloodRhFactors: BloodRhFactors,
	}

}

var BloodTypes = []string{
	"A", "B", "AB", "O",
}

var BloodRhFactors = []string{
	"+", "-",
}
