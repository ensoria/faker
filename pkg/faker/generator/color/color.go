package color

import (
	"fmt"

	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

// Color provides methods for generating random color values.
//
// ランダムな色の値を生成するメソッドを提供する構造体。
type Color struct {
	rand *core.Rand
	data *provider.Colors
}

// New creates a new Color instance with the given random source and global data.
//
// 指定されたランダムソースとグローバルデータで新しいColorインスタンスを作成する。
func New(rand *core.Rand, global *provider.Global) *Color {
	return &Color{
		rand,
		global.Colors,
	}
}

// SafeName returns a random web-safe color name.
// Example: "blue"
//
// ランダムなウェブセーフカラー名を返す。
func (c *Color) SafeName() string {
	return c.rand.Slice.StrElem(c.data.SafeColorNames)
}

// Name returns a random color name.
// Example: "NavajoWhite"
//
// ランダムな色名を返す。
func (c *Color) Name() string {
	return c.rand.Slice.StrElem(c.data.AllColorNames)
}

// Hex returns a random hex color code.
// Example: "#fa3cc2"
//
// ランダムな16進カラーコードを返す。
func (c *Color) Hex() string {
	number := c.rand.Num.IntBt(1, 16777215)
	return fmt.Sprintf("#%06x", number)
}

// SafeHex returns a random web-safe hex color code.
// Example: "#ff3300"
//
// ランダムなウェブセーフ16進カラーコードを返す。
func (c *Color) SafeHex() string {
	// Web-safe values in decimal: 0, 51, 102, 153, 204, 255
	safeValues := []int{0, 51, 102, 153, 204, 255}

	r := safeValues[c.rand.Num.Intn(len(safeValues))]
	g := safeValues[c.rand.Num.Intn(len(safeValues))]
	b := safeValues[c.rand.Num.Intn(len(safeValues))]

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// RGBAsNum returns random RGB values as three integers (0-255).
// Example: 0, 255, 122
//
// ランダムなRGB値を3つの整数（0-255）で返す。
func (c *Color) RGBAsNum() (int, int, int) {
	return c.rand.Num.IntBt(0, 255),
		c.rand.Num.IntBt(0, 255),
		c.rand.Num.IntBt(0, 255)
}

// RGBAsStr returns random RGB values as a comma-separated string.
// Example: "0,255,122"
//
// ランダムなRGB値をカンマ区切りの文字列で返す。
func (c *Color) RGBAsStr() string {
	r, g, b := c.RGBAsNum()
	return fmt.Sprintf("%d,%d,%d", r, g, b)
}

// RGBAsArr returns random RGB values as a 3-element array.
// Example: [0, 255, 122]
//
// ランダムなRGB値を3要素の配列で返す。
func (c *Color) RGBAsArr() [3]int {
	r, g, b := c.RGBAsNum()
	return [3]int{r, g, b}

}

// RGBCSS returns a random RGB CSS function string.
// Example: "rgb(0,255,122)"
//
// ランダムなRGB CSS関数文字列を返す。
func (c *Color) RGBCSS() string {
	return "rgb(" + c.RGBAsStr() + ")"
}

// RGBACSS returns a random RGBA CSS function string.
// Example: "rgba(0,255,122,0.8)"
//
// ランダムなRGBA CSS関数文字列を返す。
func (c *Color) RGBACSS() string {
	return "rgba(" + c.RGBAsStr() + "," + fmt.Sprintf("%.1f", c.rand.Num.Float32Bt(0, 1)) + ")"
}

// HSLAsNum returns random HSL values as three integers.
// Example: 340, 50, 20
//
// ランダムなHSL値を3つの整数で返す。
func (c *Color) HSLAsNum() (int, int, int) {
	return c.rand.Num.IntBt(0, 360),
		c.rand.Num.IntBt(0, 100),
		c.rand.Num.IntBt(0, 100)
}

// HSLAsStr returns random HSL values as a comma-separated string.
// Example: "340,50,20"
//
// ランダムなHSL値をカンマ区切りの文字列で返す。
func (c *Color) HSLAsStr() string {
	h, s, l := c.HSLAsNum()
	return fmt.Sprintf("%d,%d,%d", h, s, l)
}

// HSLAsArr returns random HSL values as a 3-element array.
// Example: [340, 50, 20]
//
// ランダムなHSL値を3要素の配列で返す。
func (c *Color) HSLAsArr() [3]int {
	h, s, l := c.HSLAsNum()
	return [3]int{h, s, l}
}
