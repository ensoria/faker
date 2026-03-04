package en_US

import "github.com/ensoria/faker/pkg/faker/provider"

func CreatePhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{
		Formats: formats,
	}
}

var formats = []string{
	// Standard formats
	"###-###-####",
	"(###) ###-####",
	"1-###-###-####",
	"###.###.####",
	// With extension
	"###-###-#### x###",
	"(###) ###-#### x###",
	"1-###-###-#### x###",
	"###.###.#### x###",
	"###-###-#### x####",
	"(###) ###-#### x####",
	"1-###-###-#### x####",
	"###.###.#### x####",
	"###-###-#### x#####",
	"(###) ###-#### x#####",
	"1-###-###-#### x#####",
	"###.###.#### x#####",
}
