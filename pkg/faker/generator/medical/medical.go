package medical

import (
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

// Medical provides methods for generating random medical data.
//
// ランダムな医療データを生成するメソッドを提供する構造体。
type Medical struct {
	rand *core.Rand
	data *provider.Medicals
}

// New creates a new Medical instance with the given random source and global data.
//
// 指定されたランダムソースとグローバルデータで新しいMedicalインスタンスを作成する。
func New(rand *core.Rand, global *provider.Global) *Medical {
	return &Medical{
		rand: rand,
		data: global.Medicals,
	}
}

// BloodType returns a random blood type.
// Example: "A", "B", "AB", "O"
//
// ランダムな血液型を返す。
func (m *Medical) BloodType() string {
	return m.rand.Slice.StrElem(m.data.BloodTypes)
}

// BloodRhFactor returns a random Rh factor.
// Example: "+", "-"
//
// ランダムなRh因子を返す。
func (m *Medical) BloodRhFactor() string {
	return m.rand.Slice.StrElem(m.data.BloodRhFactors)
}

// BloodGroup returns a random blood group (type + Rh factor).
// Example: "A+", "O-"
//
// ランダムな血液型グループ（型 + Rh因子）を返す。
func (m *Medical) BloodGroup() string {
	return m.BloodType() + m.BloodRhFactor()
}
