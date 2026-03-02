package medical

import (
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

type Medical struct {
	rand *core.Rand
	data *provider.Medicals
}

func New(rand *core.Rand, global *provider.Global) *Medical {
	return &Medical{
		rand: rand,
		data: global.Medicals,
	}
}

// example 'A', 'B', 'AB', 'O'
func (m *Medical) BloodType() string {
	return m.rand.Slice.StrElem(m.data.BloodTypes)
}

// example '+', '-'
func (m *Medical) BloodRhFactor() string {
	return m.rand.Slice.StrElem(m.data.BloodRhFactors)
}

// example 'A+', 'O-', etc.
func (m *Medical) BloodGroup() string {
	return m.BloodType() + m.BloodRhFactor()
}
