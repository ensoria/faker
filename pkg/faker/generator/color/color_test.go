package color_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/color"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
)

var _ = Describe("Color", func() {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Colors: global.CreateColors(),
	}
	cl := color.New(coreRand, global)

	Describe("SafeName", func() {
		It("should return random safe color name", func() {
			colorName := cl.SafeName()
			Expect(colorName).To(BeElementOf(global.Colors.SafeColorNames))
		})
	})

	Describe("Name", func() {
		It("should return random color name", func() {
			colorName := cl.Name()
			Expect(colorName).To(BeElementOf(global.Colors.AllColorNames))
		})
	})

	Describe("Hex", func() {
		It("should return random hex color string", func() {
			r := cl.Hex()
			Expect(r).To(MatchRegexp(`^#[0-9a-f]{6}$`))
		})
	})

	Describe("SafeHex", func() {
		It("should return random safe hex color string", func() {
			r := cl.SafeHex()
			Expect(r).To(MatchRegexp(`^#(00|33|66|99|[Cc][Cc]|[Ff][Ff]){3}$`))
		})
	})

	Describe("RgbAsNum", func() {
		It("should return random rgb color as numbers", func() {
			r, g, b := cl.RgbAsNum()

			Expect(r).To(BeNumerically(">", -1))
			Expect(r).To(BeNumerically("<", 256))
			Expect(g).To(BeNumerically(">", -1))
			Expect(g).To(BeNumerically("<", 256))
			Expect(b).To(BeNumerically(">", -1))
			Expect(b).To(BeNumerically("<", 256))
		})
	})

	regex255 := `([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])`
	regexAlpha := `([01]?(\.\d+)?)`

	Describe("RgbAsStr", func() {
		It("should return random rgb color string", func() {
			r := cl.RgbAsStr()
			Expect(r).To(MatchRegexp("^" + regex255 + "," + regex255 + "," + regex255 + "$"))
		})
	})

	Describe("RgbAsArr", func() {
		It("should return random rgb color as array", func() {
			r := cl.RgbAsArr()

			Expect(r[0]).To(BeNumerically(">", -1))
			Expect(r[0]).To(BeNumerically("<", 256))
			Expect(r[1]).To(BeNumerically(">", -1))
			Expect(r[1]).To(BeNumerically("<", 256))
			Expect(r[2]).To(BeNumerically(">", -1))
			Expect(r[2]).To(BeNumerically("<", 256))
		})
	})

	Describe("RgbCss", func() {
		It("should return random rgb css color string", func() {
			r := cl.RgbCss()
			Expect(r).To(MatchRegexp(`^rgb\(` + regex255 + `,` + regex255 + `,` + regex255 + `\)$`))
		})
	})

	Describe("RgbaCss", func() {
		It("should return random rgba css color string", func() {
			r := cl.RgbaCss()
			Expect(r).To(MatchRegexp(`^rgba\(` + regex255 + `,` + regex255 + `,` + regex255 + `,` + regexAlpha + `\)$`))
		})
	})

	Describe("HslAsNum", func() {
		It("should return random hsl color as numbers", func() {
			h, s, l := cl.HslAsNum()
			Expect(h).To(BeNumerically(">", -1))
			Expect(h).To(BeNumerically("<", 361))
			Expect(s).To(BeNumerically(">", -1))
			Expect(s).To(BeNumerically("<", 101))
			Expect(l).To(BeNumerically(">", -1))
			Expect(l).To(BeNumerically("<", 101))
		})
	})

	regex360 := `(?:36[0]|3[0-5][0-9]|[12][0-9][0-9]|[1-9]?[0-9])`
	regex100 := `(?:100|[1-9]?[0-9])`

	Describe("HslAsStr", func() {
		It("should return random hsl color string", func() {
			r := cl.HslAsStr()
			Expect(r).To(MatchRegexp("^" + regex360 + "," + regex100 + "," + regex100 + "$"))
		})
	})

	Describe("HslAsArr", func() {
		It("should return random hsl color as array", func() {
			r := cl.HslAsArr()
			Expect(r[0]).To(BeNumerically(">", -1))
			Expect(r[0]).To(BeNumerically("<", 361))
			Expect(r[1]).To(BeNumerically(">", -1))
			Expect(r[1]).To(BeNumerically("<", 101))
			Expect(r[2]).To(BeNumerically(">", -1))
			Expect(r[2]).To(BeNumerically("<", 101))
		})
	})
})
