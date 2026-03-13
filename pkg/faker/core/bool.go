package core

import (
	"fmt"
	"math/rand"

	"github.com/ensoria/faker/pkg/faker/common/log"
)

// RandBool provides methods for generating random boolean values.
// ランダムな真偽値を生成するメソッドを提供する構造体。
type RandBool struct {
	rand *rand.Rand
}

// NewRandBool creates a new RandBool instance with the given random source.
// 指定されたランダムソースで新しいRandBoolインスタンスを作成する。
func NewRandBool(rand *rand.Rand) *RandBool {
	return &RandBool{
		rand,
	}
}

// Evenly returns true or false with a 50% probability each.
//
// 50%の確率でtrueかfalseを返す。
func (r *RandBool) Evenly() bool {
	return r.rand.Intn(2) == 0
}

// WeightedToTrue returns true with the given weight probability (0.0 to 1.0).
// 指定された重み（0.0〜1.0）の確率でtrueを返す。
func (r *RandBool) WeightedToTrue(weight float32) bool {
	if weight < 0 || weight > 1 {
		errMsg := fmt.Sprintf("Invalid weight: %f", weight)
		log.WrongUsage(errMsg, 1)
		return false
	}
	return r.rand.Float32() < weight
}
