package image_test

import (
	"bytes"
	"encoding/base64"
	stdimage "image"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/image"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
)

var _ = Describe("Image", func() {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Images: global.CreateImages(),
	}

	imgFk := image.New(coreRand, global)

	Describe("Binary", func() {
		It("should return specified image binary data", func() {
			r, err := imgFk.Binary(100, 100, image.JPG)
			img, format, _ := stdimage.Decode(bytes.NewReader(r))

			Expect(err).To(BeNil())
			Expect(format).To(Equal("jpeg"))
			bounds := img.Bounds()
			width := bounds.Dx()
			Expect(width).To(Equal(100))
			height := bounds.Dy()
			Expect(height).To(Equal(100))
		})

		When("more than 3840px width is specified", func() {
			It("should return 3840px width image", func() {
				r, _ := imgFk.Binary(4000, 100, image.JPG)

				img, _, _ := stdimage.Decode(bytes.NewReader(r))
				bounds := img.Bounds()
				width := bounds.Dx()
				Expect(width).To(Equal(image.MaxHeightWidth))
			})
		})

		When("more than 3840px height is specified", func() {
			It("should return 3840px height image", func() {
				r, _ := imgFk.Binary(100, 4000, image.JPG)

				img, _, _ := stdimage.Decode(bytes.NewReader(r))
				bounds := img.Bounds()
				height := bounds.Dy()
				Expect(height).To(Equal(image.MaxHeightWidth))
			})
		})

		When("an unsupported format is specified", func() {
			It("should return default jpg image", func() {
				r, _ := imgFk.Binary(100, 100, "unsupported_format")

				_, format, _ := stdimage.Decode(bytes.NewReader(r))
				Expect(format).To(Equal("jpeg"))
			})
		})
	})

	Describe("Object", func() {
		It("should return specified image object", func() {
			img, err := imgFk.Object(100, 100, image.JPG)

			Expect(err).To(BeNil())
			bounds := img.Bounds()
			width := bounds.Dx()
			Expect(width).To(Equal(100))
			height := bounds.Dy()
			Expect(height).To(Equal(100))
		})

		When("more than 3840px width is specified", func() {
			It("should log an error and return 3840px width image", func() {
				img, _ := imgFk.Object(4000, 100, image.JPG)

				bounds := img.Bounds()
				width := bounds.Dx()
				Expect(width).To(Equal(image.MaxHeightWidth))
			})
		})
	})

	Describe("Base64", func() {
		It("should return base64 string image", func() {
			r, err := imgFk.Base64(100, 100, image.JPG)

			img, _ := base64ToImage(r)
			Expect(err).To(BeNil())
			bounds := img.Bounds()
			width := bounds.Dx()
			Expect(width).To(Equal(100))
			height := bounds.Dy()
			Expect(height).To(Equal(100))
		})

		When("more than 3840px width is specified", func() {
			It("should log an error and return 3840px width image", func() {
				r, _ := imgFk.Base64(4000, 100, image.JPG)

				img, _ := base64ToImage(r)
				bounds := img.Bounds()
				width := bounds.Dx()
				Expect(width).To(Equal(image.MaxHeightWidth))
			})
		})
	})
})

func base64ToImage(base64Str string) (stdimage.Image, error) {
	// Decode the base64 string to binary data
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	img, _, err := stdimage.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}
