package faker

import (
	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/address"
	"github.com/ensoria/gofake/pkg/faker/generator/barcode"
	"github.com/ensoria/gofake/pkg/faker/generator/color"
	"github.com/ensoria/gofake/pkg/faker/generator/company"
	"github.com/ensoria/gofake/pkg/faker/generator/file"
	"github.com/ensoria/gofake/pkg/faker/generator/image"
	"github.com/ensoria/gofake/pkg/faker/generator/internet"
	"github.com/ensoria/gofake/pkg/faker/generator/lorem"
	"github.com/ensoria/gofake/pkg/faker/generator/medical"
	"github.com/ensoria/gofake/pkg/faker/generator/payment"
	"github.com/ensoria/gofake/pkg/faker/generator/person"
	"github.com/ensoria/gofake/pkg/faker/generator/phonenumber"
	"github.com/ensoria/gofake/pkg/faker/generator/useragent"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/global"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_US"
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
	return newFaker(core.NewRand(util.RandSeed()), localized)
}

// CreateWithSeed creates a new Faker instance with the default en_US locale and
// a fixed seed, producing deterministic output (useful for tests / golden files).
//
// デフォルトのen_USロケールと固定シードで新しいFakerインスタンスを作成する。
// 出力は決定的になる（テストやゴールデンファイル向け）。
func CreateWithSeed(seed int64) *Faker {
	return CreateWithLocaleAndSeed(en_US.New(), seed)
}

// CreateWithLocaleAndSeed creates a new Faker instance with the specified locale
// data and a fixed seed, producing deterministic output.
//
// 指定されたロケールデータと固定シードで新しいFakerインスタンスを作成する。
func CreateWithLocaleAndSeed(localized *provider.Localized, seed int64) *Faker {
	return newFaker(core.NewRand(util.SeededRand(seed)), localized)
}

// newFaker wires all generators onto the given core.Rand and locale.
//
// 与えられた core.Rand とロケールで全ジェネレーターを組み立てる内部ヘルパー。
func newFaker(coreRand *core.Rand, localized *provider.Localized) *Faker {
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
