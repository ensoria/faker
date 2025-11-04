package address_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/address"
	"github.com/ensoria/faker/pkg/faker/provider/locale/en_US"
	"github.com/ensoria/faker/pkg/faker/provider/locale/ja_JP"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Address", func() {
	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	addressUs := address.New(coreRand, localized)

	Describe("Country", func() {
		It("should return a country", func() {
			r := addressUs.Country()
			Expect(r).To(BeElementOf(en_US.Countries))
		})
	})

	Describe("Postcode", func() {
		It("should return a postcode", func() {
			r := addressUs.Postcode()
			testutil.Output("Address.Postcode", r)
		})
	})

	Describe("StateAbbr", func() {
		It("should return a state abbreviation", func() {
			r := addressUs.StateAbbr()
			Expect(r).To(BeElementOf(en_US.StateAbbrs))
		})
	})

	Describe("State", func() {
		It("should return a state", func() {
			r := addressUs.State()
			Expect(r).To(BeElementOf(en_US.States))
		})
	})

	Describe("City", func() {
		It("CityName should return a city name", func() {
			r := addressUs.CityName()
			Expect(r).To(BeElementOf(*en_US.CityNames))
		})

		It("CitySuffix should return a city suffix", func() {
			r := addressUs.CitySuffix()
			Expect(r).To(BeElementOf(en_US.CitySuffixes))
		})

		It("CityPrefix should return a city prefix", func() {
			r := addressUs.CityPrefix()
			Expect(r).To(BeElementOf(en_US.CityPrefixes))
		})

		It("City should return a city", func() {
			r := addressUs.City()
			testutil.Output("Address.City", r)
		})
	})

	Describe("Street", func() {
		It("StreetName should return a street name", func() {
			r := addressUs.StreetName()
			Expect(r).To(BeElementOf(*en_US.StreetNames))
		})

		It("StreetSuffix should return a street suffix", func() {
			r := addressUs.StreetSuffix()
			Expect(r).To(BeElementOf(en_US.StreetSuffixes))
		})

		It("Street should return a street name", func() {
			r := addressUs.Street()
			testutil.Output("Address.Street", r)
		})
	})

	Describe("SecondaryAddress", func() {
		It("BuildingNumber should return a building number", func() {
			r := addressUs.BuildingNumber()
			testutil.Output("Address.BuildingNumber", r)
		})

		It("SecondaryAddress should return a secondary address", func() {
			r := addressUs.SecondaryAddress()
			testutil.Output("Address.SecondaryAddress", r)
		})
	})

	Describe("StreetAddress", func() {
		It("should return a street address", func() {
			r := addressUs.StreetAddress()
			testutil.Output("Address.StreetAddress", r)
		})
	})

	Describe("Address", func() {
		It("should return an address", func() {
			r := addressUs.Address()
			testutil.Output("Address.Address", r)
		})
	})

	Describe("Latitude", func() {
		It("should return a latitude", func() {
			r := addressUs.Latitude()
			Expect(r).To(BeNumerically(">=", -90.0))
			Expect(r).To(BeNumerically("<=", 90.0))

			length := testutil.GetDecimalLength(r)
			// when the end of the number is 0, float64 may have less than 5 digits instead of 6
			Expect(length).To(BeNumerically("<=", 6))
			testutil.Output("Address.Latitude", r)
		})
	})

	Describe("Longitude", func() {
		It("should return a longitude", func() {
			r := addressUs.Longitude()
			Expect(r).To(BeNumerically(">=", -180.0))
			Expect(r).To(BeNumerically("<=", 180.0))

			length := testutil.GetDecimalLength(r)
			// when the end of the number is 0, float64 may have less than 5 digits instead of 6
			Expect(length).To(BeNumerically("<=", 6))
			testutil.Output("Address.Longitude", r)
		})
	})

	Describe("LocalCoordinates", func() {
		It("should return a coordinate", func() {
			lat, lon := addressUs.LocalCoordinates()
			Expect(lat).To(BeNumerically(">=", -90.0))
			Expect(lat).To(BeNumerically("<=", 90.0))
			latLength := testutil.GetDecimalLength(lat)
			Expect(latLength).To(BeNumerically("<=", 6))

			Expect(lon).To(BeNumerically(">=", -180.0))
			Expect(lon).To(BeNumerically("<=", 180.0))
			lonLength := testutil.GetDecimalLength(lon)
			Expect(lonLength).To(BeNumerically("<=", 6))
		})
	})

	localizedJaJp := ja_JP.New()
	addressJaJp := address.New(coreRand, localizedJaJp)

	Describe("Prefecture", func() {
		It("should return a prefecture", func() {
			r := addressJaJp.Prefecture()
			Expect(r).To(BeElementOf(ja_JP.Prefectures))
		})
	})

	Describe("Ward", func() {
		It("WardSuffix should return a ward suffix", func() {
			r := addressJaJp.WardSuffix()
			Expect(r).To(BeElementOf(ja_JP.WardSuffixes))
		})

		It("WardName should return a ward name", func() {
			r := addressJaJp.WardName()
			Expect(r).To(BeElementOf(ja_JP.WardNames))
		})

		It("Ward should return a ward", func() {
			r := addressJaJp.Ward()
			testutil.Output("Address.Ward", r)
		})
	})

	Describe("Area", func() {
		It("AreaNumber should return an area number", func() {
			r := addressJaJp.AreaNumber()
			testutil.Output("Address.AreaNumber", r)
		})

		It("AreaName should return an area name", func() {
			r := addressJaJp.AreaName()
			Expect(r).To(BeElementOf(ja_JP.AreaNames))
		})

		It("Area should return an area", func() {
			r := addressJaJp.Area()
			testutil.Output("Address.Area", r)
		})
	})

	Describe("SecondaryAddress for ja_JP", func() {
		It("BuildingName should return a building name", func() {
			r := addressJaJp.BuildingName()
			Expect(r).To(BeElementOf(ja_JP.BuildingNames))
		})

		It("RoomNumber should return a room number", func() {
			r := addressJaJp.RoomNumber()
			testutil.Output("Address.RoomNumber", r)
		})

		It("SecondaryAddress for ja_JP should return a secondary address", func() {
			r := addressJaJp.SecondaryAddress()
			testutil.Output("Address.SecondaryAddress", r)
		})
	})

	Describe("Address for ja_JP", func() {
		It("should return an address", func() {
			r := addressJaJp.Address()
			testutil.Output("Address.Address", r)
		})
	})

})
