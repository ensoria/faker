package core

import (
	"math/rand"
)

// Rand aggregates all random value generators.
// すべてのランダム値ジェネレーターを集約する構造体。
type Rand struct {
	Str   *RandStr
	Num   *RandNum
	Bool  *RandBool
	Slice *RandSlice
	Map   *RandMap
	Time  *RandTime
}

// NewRand creates a new Rand instance with the given random source.
// 指定されたランダムソースで新しいRandインスタンスを作成する。
func NewRand(rand *rand.Rand) *Rand {
	return &Rand{
		Str:   NewRandStr(rand),
		Num:   NewRandNum(rand),
		Bool:  NewRandBool(rand),
		Slice: NewRandSlice(rand),
		Map:   NewRandMap(rand),
		Time:  NewRandTime(rand),
	}
}
