package faker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker"
	"github.com/ensoria/faker/pkg/faker/provider/locale/ja_JP"
)

var _ = Describe("Faker", func() {

	Context("faker with ja_JP locale", func() {
		j := ja_JP.New()
		f := faker.CreateWithLocale(j)
		It("just tinkering; not testing, just experimenting with outputs", func() {
			f.Rand.Bool.Evenly()
			f.Rand.Bool.WeightedToTrue(0.8)

			f.Rand.Num.IntBt(1, 10)
			f.Rand.Num.Int32Bt(1, 10)
			f.Rand.Num.Int64Bt(1, 10)
			f.Rand.Num.Float32Bt(1.0, 10.0)
			f.Rand.Num.Float64Bt(1.0, 10.0)

			// rand.Randのメソッドを使いたい場合は、エイリアスが用意されています
			f.Rand.Num.Int()
			f.Rand.Num.Intn(10)
			f.Rand.Str.Char()
			f.Rand.Time.PastFuture()

			f.Rand.Slice.IntElem([]int{1, 2, 3})
			f.Barcode.Ean13()
			f.Color.Hex()
			f.File.MimeType()
			f.Image.Binary(100, 100, "jpg")
			f.Internet.FirstName()
			f.Lorem.Word()
			f.Address.City()
			f.Company.Name()

			// Just a placeholder expectation to ensure the test runs
			Expect(true).To(Equal(true))
		})
	})

})
