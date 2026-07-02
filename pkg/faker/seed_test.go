package faker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
)

var _ = Describe("Seeded Faker", func() {
	// 同一シードなら同一の値列を生成する（決定的）ことを確認する。
	It("produces the same sequence for the same seed", func() {
		a := faker.CreateWithSeed(42)
		b := faker.CreateWithSeed(42)

		for i := 0; i < 20; i++ {
			Expect(a.Rand.Num.IntBt(0, 1_000_000)).To(Equal(b.Rand.Num.IntBt(0, 1_000_000)))
			Expect(a.Rand.Str.AlphaFixedLength(12)).To(Equal(b.Rand.Str.AlphaFixedLength(12)))
			Expect(a.Person.Name()).To(Equal(b.Person.Name()))
		}
	})

	It("produces different sequences for different seeds", func() {
		a := faker.CreateWithSeed(1)
		b := faker.CreateWithSeed(2)

		// 少なくとも1つは異なるはず（同一なら決定性の意味がない）
		differs := false
		for i := 0; i < 20; i++ {
			if a.Rand.Num.IntBt(0, 1_000_000) != b.Rand.Num.IntBt(0, 1_000_000) {
				differs = true
				break
			}
		}
		Expect(differs).To(BeTrue())
	})
})
