package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/ensoria/gofake/pkg/faker/common/log"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

// Image provides methods for generating random dummy images.
//
// ランダムなダミー画像を生成するメソッドを提供する構造体。
type Image struct {
	rand    *core.Rand
	data    *provider.Images
	logSkip int
}

// New creates a new Image instance with the given random source and global data.
//
// 指定されたランダムソースとグローバルデータで新しいImageインスタンスを作成する。
func New(rand *core.Rand, global *provider.Global) *Image {
	return &Image{
		rand,
		global.Images,
		1,
	}
}

// ImageFormat represents the image file format.
//
// 画像ファイルのフォーマットを表す型。
type ImageFormat string

const (
	JPG = ImageFormat("jpg")
	PNG = ImageFormat("png")
	GIF = ImageFormat("gif")
)

// MaxHeightWidth is the maximum allowed width and height in pixels (3840px).
//
// 許可される最大の幅と高さ（3840ピクセル）。
const MaxHeightWidth = 3840

// Binary generates a white blank image as binary data.
// Width and height should be 3840px or less.
//
// 白い空白画像をバイナリデータとして生成する。
// 幅と高さは3840ピクセル以下にすること。
func (i *Image) Binary(width int, height int, format ImageFormat) ([]byte, error) {
	imgWidth := width
	if width > MaxHeightWidth {
		imgWidth = MaxHeightWidth
		log.GeneralError("Image width is too large, it will be set to 3840px", i.logSkip)
	}
	imgHeight := height
	if height > MaxHeightWidth {
		imgHeight = MaxHeightWidth
		log.GeneralError("Image height is too large, it will be set to 3840px", i.logSkip)
	}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Set each pixel color to white
	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	buf := new(bytes.Buffer)
	var err error
	switch format {
	case JPG:
		err = jpeg.Encode(buf, img, nil)
	case PNG:
		err = png.Encode(buf, img)
	case GIF:
		err = gif.Encode(buf, img, nil)
	default:
		// default to jpg
		err = jpeg.Encode(buf, img, nil)
		log.GeneralError("Image format must be either [image.JPG], [image.PNG] or [image.GIF]. If not any of those, it defaults to JPG.", 1)
	}

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Object generates a white blank image as an image.Image object.
// Width and height should be 3840px or less.
//
// 白い空白画像をimage.Imageオブジェクトとして生成する。
// 幅と高さは3840ピクセル以下にすること。
func (i *Image) Object(width int, height int, format ImageFormat) (image.Image, error) {
	i.logSkip = 2
	binary, err := i.Binary(width, height, format)
	i.logSkip = 1 // reset logSkip
	if err != nil {
		return nil, err
	}

	obj, _, errDecode := image.Decode(bytes.NewReader(binary))
	return obj, errDecode
}

// Base64 generates a white blank image as a Base64-encoded string.
// Width and height should be 3840px or less.
//
// 白い空白画像をBase64エンコードされた文字列として生成する。
// 幅と高さは3840ピクセル以下にすること。
func (i *Image) Base64(width int, height int, format ImageFormat) (string, error) {
	i.logSkip = 2
	img, err := i.Binary(width, height, format)
	i.logSkip = 1 // reset logSkip

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}
