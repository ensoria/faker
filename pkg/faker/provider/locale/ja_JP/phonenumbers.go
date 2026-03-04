package ja_JP

import "github.com/ensoria/faker/pkg/faker/provider"

func CreatePhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{
		Formats: formats,
	}
}

// See: http://www.soumu.go.jp/main_sosiki/joho_tsusin/top/tel_number/number_shitei.html#kotei-denwa
var formats = []string{
	"080-####-####",
	"090-####-####",
	"0#-####-####",
	"0####-#-####",
	"0###-##-####",
	"0##-###-####",
	"0##0-###-###",
	"0#########",
	"0##########",
}
