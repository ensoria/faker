package address

import (
	"github.com/ensoria/faker/pkg/faker/common/log"
	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

// Address provides methods for generating random address data.
//
// ランダムな住所データを生成するメソッドを提供する構造体。
type Address struct {
	rand *core.Rand
	data *provider.Addresses
}

// New creates a new Address instance with the given random source and locale data.
//
// 指定されたランダムソースとロケールデータで新しいAddressインスタンスを作成する。
func New(rand *core.Rand, localized *provider.Localized) *Address {
	return &Address{
		rand,
		localized.Addresses,
	}
}

// CitySuffix returns a random city suffix.
// Example: "town"
//
// ランダムな市区町村の接尾辞を返す。
func (a *Address) CitySuffix() string {
	if len(a.data.CitySuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.CitySuffixes)
}

// CityPrefix returns a random city prefix.
//
// ランダムな市区町村の接頭辞を返す。
func (a *Address) CityPrefix() string {
	if len(a.data.CityPrefixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.CityPrefixes)
}

// CityName returns a random city name.
//
// ランダムな市区町村名を返す。
func (a *Address) CityName() string {
	if len(a.data.CityNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.CityNames)
}

// City returns a randomly formatted city name using locale-specific formats.
// Example: "Shieldsfurt"
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた市区町村名を返す。
func (a *Address) City() string {
	if len(a.data.CityFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.CityFormats)
	cityNameData := a.data.CreateCity(a)
	return util.RenderTemplate(format, cityNameData)
}

// StreetSuffix returns a random street suffix.
// Example: "Avenue"
//
// ランダムな通りの接尾辞を返す。
func (a *Address) StreetSuffix() string {
	if len(a.data.StreetSuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.StreetSuffixes)
}

// StreetName returns a random street name.
//
// ランダムな通り名を返す。
func (a *Address) StreetName() string {
	if len(a.data.StreetNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.StreetNames)
}

// Street returns a randomly formatted street name using locale-specific formats.
// Example: "Lindgren Dam"
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた通り名を返す。
func (a *Address) Street() string {
	if len(a.data.StreetFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.StreetFormats)
	streetNameData := a.data.CreateStreet(a)
	return util.RenderTemplate(format, streetNameData)
}

// BuildingNumber returns a random building number.
// Example: "791"
//
// ランダムな建物番号を返す。
func (a *Address) BuildingNumber() string {
	if len(a.data.BuildingNumbers) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.BuildingNumbers)
	return a.rand.Str.AlphaDigitsLike(format)
}

// SecondaryAddress returns a random secondary address (e.g. apartment number).
// Example: "Apt. 160"
//
// ランダムな補助住所（例: 部屋番号）を返す。
func (a *Address) SecondaryAddress() string {
	if len(a.data.SecondaryAddressFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.SecondaryAddressFormats)
	data := a.data.CreateSecondaryAddress(a)
	return util.RenderTemplate(format, data)
}

// StreetAddress returns a randomly formatted street address.
//
// ランダムにフォーマットされた通り住所を返す。
func (a *Address) StreetAddress() string {
	if len(a.data.StreetAddressFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.StreetAddressFormats)
	streetAddressData := a.data.CreateStreetAddress(a)
	return util.RenderTemplate(format, streetAddressData)
}

// Postcode returns a random postcode/zip code.
// Example: "87678"
//
// ランダムな郵便番号を返す。
func (a *Address) Postcode() string {
	if len(a.data.Postcodes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.Postcodes)
	return a.rand.Str.AlphaDigitsLike(format)
}

// StateAbbr returns a random state abbreviation.
// Example: "CA"
//
// ランダムな都道府県/州の略称を返す。
func (a *Address) StateAbbr() string {
	if len(a.data.StateAbbrs) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.StateAbbrs)
}

// State returns a random state or province name.
// Example: "California"
//
// ランダムな都道府県/州名を返す。
func (a *Address) State() string {
	if len(a.data.States) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.States)
}

// Address returns a randomly formatted full address.
// Example: "026 Rolfson Summit\nPollichfurt, AZ 34986"
//
// ランダムにフォーマットされた完全な住所を返す。
func (a *Address) Address() string {
	if len(a.data.AddressFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.AddressFormats)
	addressData := a.data.CreateAddress(a)
	return util.RenderTemplate(format, addressData)
}

// Country returns a random country name.
// Example: "United States of America"
//
// ランダムな国名を返す。
func (a *Address) Country() string {
	if len(a.data.Countries) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.Countries)
}

// Prefecture returns a random prefecture name (for locales that have prefectures, e.g. ja_JP).
//
// ランダムな都道府県名を返す（日本語ロケールなど、都道府県を持つロケール用）。
func (a *Address) Prefecture() string {
	if len(a.data.Prefectures) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.Prefectures)
}

// WardSuffix returns a random ward suffix (e.g. "区").
//
// ランダムな区の接尾辞を返す（例: "区"）。
func (a *Address) WardSuffix() string {
	if len(a.data.WardSuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.WardSuffixes)
}

// WardName returns a random ward name.
//
// ランダムな区名を返す。
func (a *Address) WardName() string {
	if len(a.data.WardNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.WardNames)
}

// Ward returns a randomly formatted ward name using locale-specific formats.
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた区名を返す。
func (a *Address) Ward() string {
	if len(a.data.WardFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.WardFormats)
	wardData := a.data.CreateWard(a)
	return util.RenderTemplate(format, wardData)
}

// AreaName returns a random area/town name.
//
// ランダムな地域名/町名を返す。
func (a *Address) AreaName() string {
	if len(a.data.AreaNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.AreaNames)
}

// AreaNumber returns a random area number (e.g. "3丁目2番地1").
//
// ランダムな番地を返す（例: "3丁目2番地1"）。
func (a *Address) AreaNumber() string {
	if len(a.data.AreaNumbers) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.AreaNumbers)
	return a.rand.Str.AlphaDigitsLike(format)
}

// Area returns a randomly formatted area name using locale-specific formats.
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた地域名を返す。
func (a *Address) Area() string {
	if len(a.data.AreaFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.AreaFormats)
	areaData := a.data.CreateArea(a)
	return util.RenderTemplate(format, areaData)

}

// BuildingName returns a random building name.
//
// ランダムな建物名を返す。
func (a *Address) BuildingName() string {
	if len(a.data.BuildingNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.BuildingNames)
}

// RoomNumber returns a random room number.
//
// ランダムな部屋番号を返す。
func (a *Address) RoomNumber() string {
	if len(a.data.RoomNumbers) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.RoomNumbers)
	return a.rand.Str.AlphaDigitsLike(format)
}

// Latitude returns a random latitude value with 6 decimal places.
// Example: 35.785163
//
// 小数点以下6桁のランダムな緯度を返す。
func (a *Address) Latitude() float64 {
	val := a.rand.Num.Float64Bt(-90, 90)
	return util.TruncateToPrecision(val, 6)
}

// Longitude returns a random longitude value with 6 decimal places.
// Example: -71.462048
//
// 小数点以下6桁のランダムな経度を返す。
func (a *Address) Longitude() float64 {
	val := a.rand.Num.Float64Bt(-180, 180)
	return util.TruncateToPrecision(val, 6)
}

// LocalCoordinates returns a random latitude and longitude pair.
// Example: 35.785163, -71.462048
//
// ランダムな緯度と経度のペアを返す。
func (a *Address) LocalCoordinates() (float64, float64) {
	return a.Latitude(), a.Longitude()
}
