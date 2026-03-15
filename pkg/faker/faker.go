package faker

import (
	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/address"
	"github.com/ensoria/faker/pkg/faker/generator/barcode"
	"github.com/ensoria/faker/pkg/faker/generator/color"
	"github.com/ensoria/faker/pkg/faker/generator/company"
	"github.com/ensoria/faker/pkg/faker/generator/file"
	"github.com/ensoria/faker/pkg/faker/generator/image"
	"github.com/ensoria/faker/pkg/faker/generator/internet"
	"github.com/ensoria/faker/pkg/faker/generator/lorem"
	"github.com/ensoria/faker/pkg/faker/generator/medical"
	"github.com/ensoria/faker/pkg/faker/generator/payment"
	"github.com/ensoria/faker/pkg/faker/generator/person"
	"github.com/ensoria/faker/pkg/faker/generator/phonenumber"
	"github.com/ensoria/faker/pkg/faker/generator/useragent"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
	"github.com/ensoria/faker/pkg/faker/provider/locale/en_US"
)

// Faker is the main entry point that aggregates all fake data generators.
//
// すべてのフェイクデータジェネレーターを集約するメインのエントリーポイント構造体。
type Faker struct {
	Rand        *core.Rand
	Person      *person.Person
	Color       *color.Color
	Address     *address.Address
	Barcode     *barcode.Barcode
	Company     *company.Company
	File        *file.File
	Image       *image.Image
	Internet    *internet.Internet
	Lorem       *lorem.Lorem
	Medical     *medical.Medical
	Payment     *payment.Payment
	PhoneNumber *phonenumber.PhoneNumber
	UserAgent   *useragent.UserAgent
}

// REF: https://fakerphp.github.io/

// Create creates a new Faker instance with the default en_US locale.
//
// デフォルトのen_USロケールで新しいFakerインスタンスを作成する。
func Create() *Faker {
	localized := en_US.New()
	return CreateWithLocale(localized)
}

// CreateWithLocale creates a new Faker instance with the specified locale data.
//
// 指定されたロケールデータで新しいFakerインスタンスを作成する。
func CreateWithLocale(localized *provider.Localized) *Faker {
	coreRand := core.NewRand(util.RandSeed())
	global := global.New()
	return &Faker{
		Rand:        coreRand,
		Barcode:     barcode.New(coreRand),
		Color:       color.New(coreRand, global),
		Person:      person.New(coreRand, localized),
		Address:     address.New(coreRand, localized),
		Company:     company.New(coreRand, localized),
		File:        file.New(coreRand, global),
		Image:       image.New(coreRand, global),
		Internet:    internet.New(coreRand, global),
		Lorem:       lorem.New(coreRand, global),
		Medical:     medical.New(coreRand, global),
		Payment:     payment.New(coreRand, global),
		PhoneNumber: phonenumber.New(coreRand, global, localized),
		UserAgent:   useragent.New(coreRand, global),
	}
}
