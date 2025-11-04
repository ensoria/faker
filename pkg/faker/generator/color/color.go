package color

import (
	"fmt"

	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

type Color struct {
	rand *core.Rand
	data *provider.Colors
}

func New(rand *core.Rand, global *provider.Global) *Color {
	return &Color{
		rand,
		global.Colors,
	}
}

// example 'blue'
func (c *Color) SafeName() string {
	return c.rand.Slice.StrElem(c.data.SafeColorNames)
}

// example 'NavajoWhite'
func (c *Color) Name() string {
	return c.rand.Slice.StrElem(c.data.AllColorNames)
}

// example '#fa3cc2'
func (c *Color) Hex() string {
	number := c.rand.Num.IntBt(1, 16777215)
	return fmt.Sprintf("#%06x", number)
}

// example '#ff3300'
func (c *Color) SafeHex() string {
	// Web-safe values in decimal: 0, 51, 102, 153, 204, 255
	safeValues := []int{0, 51, 102, 153, 204, 255}

	r := safeValues[c.rand.Num.Intn(len(safeValues))]
	g := safeValues[c.rand.Num.Intn(len(safeValues))]
	b := safeValues[c.rand.Num.Intn(len(safeValues))]

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// example 0, 255, 122
func (c *Color) RgbAsNum() (int, int, int) {
	return c.rand.Num.IntBt(0, 255),
		c.rand.Num.IntBt(0, 255),
		c.rand.Num.IntBt(0, 255)
}

// example '0,255,122'
func (c *Color) RgbAsStr() string {
	r, g, b := c.RgbAsNum()
	return fmt.Sprintf("%d,%d,%d", r, g, b)
}

// example [0, 255, 122]
func (c *Color) RgbAsArr() [3]int {
	r, g, b := c.RgbAsNum()
	return [3]int{r, g, b}

}

// example 'rgb(0,255,122)'
func (c *Color) RgbCss() string {
	return "rgb(" + c.RgbAsStr() + ")"
}

// example 'rgba(0,255,122,0.8)'
func (c *Color) RgbaCss() string {
	return "rgba(" + c.RgbAsStr() + "," + fmt.Sprintf("%.1f", c.rand.Num.Float32Bt(0, 1)) + ")"
}

// example 340, 50, 20
func (c *Color) HslAsNum() (int, int, int) {
	return c.rand.Num.IntBt(0, 360),
		c.rand.Num.IntBt(0, 100),
		c.rand.Num.IntBt(0, 100)
}

// example '340,50,20'
func (c *Color) HslAsStr() string {
	h, s, l := c.HslAsNum()
	return fmt.Sprintf("%d,%d,%d", h, s, l)
}

// example [340, 50, 20]
func (c *Color) HslAsArr() [3]int {
	h, s, l := c.HslAsNum()
	return [3]int{h, s, l}
}
