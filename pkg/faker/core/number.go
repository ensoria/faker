package core

import (
	"fmt"
	"math/rand"

	"github.com/ensoria/gofake/pkg/faker/common/log"
)

// RandNum provides methods for generating random numeric values.
// ランダムな数値を生成するメソッドを提供する構造体。
type RandNum struct {
	rand *rand.Rand
}

type orderable interface {
	int |
		int8 |
		int16 |
		int32 |
		int64 |
		uint |
		uint8 |
		uint16 |
		uint32 |
		uint64 |
		uintptr |
		float32 |
		float64
}

// NewRandNum creates a new RandNum instance with the given random source.
// 指定されたランダムソースで新しいRandNumインスタンスを作成する。
func NewRandNum(rand *rand.Rand) *RandNum {
	return &RandNum{
		rand,
	}
}

// IntBt returns a random integer between min and max (inclusive).
// min以上max以下のランダムな整数を返す。maxは含まれる。
func (r *RandNum) IntBt(min int, max int) int {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Intn(randMax+1) + min
}

// Int32Bt returns a random int32 between min and max (inclusive).
// min以上max以下のランダムなint32を返す。maxは含まれる。
func (r *RandNum) Int32Bt(min int32, max int32) int32 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Int31n(randMax+1) + min
}

// Int64Bt returns a random int64 between min and max (inclusive).
// min以上max以下のランダムなint64を返す。maxは含まれる。
func (r *RandNum) Int64Bt(min int64, max int64) int64 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Int63n(randMax+1) + min
}

// Float32Bt returns a random float32 between min (inclusive) and max (exclusive).
// The result may approximate max but will never equal it.
// min以上max未満のランダムなfloat32を返す。maxの近似値になることはあるが、maxを返すことはない。
func (r *RandNum) Float32Bt(min float32, max float32) float32 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Float32()*randMax + min
}

// Float64Bt returns a random float64 between min (inclusive) and max (exclusive).
// The result may approximate max but will never equal it.
// min以上max未満のランダムなfloat64を返す。maxの近似値になることはあるが、maxを返すことはない。
func (r *RandNum) Float64Bt(min float64, max float64) float64 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Float64()*randMax + min
}

// randMaxRange returns the range (max - min) for use with Intn() and similar functions.
// This is a shared helper for all numeric types.
// Intn()などに渡すためのランダム値の範囲（max - min）を返す。全ての数値型で共通の処理。
func randMaxRange[N orderable](min N, max N) (N, error) {
	var err error
	if min >= max {
		errMsg := fmt.Sprintf("Invalid range: min=%v, max=%v", min, max)
		log.WrongUsage(errMsg, 2)
		err = fmt.Errorf("%s", errMsg)
	}

	return max - min, err
}

// Int returns a non-negative pseudo-random int.
// 非負の疑似乱数intを返す。rand.Rand.Intのエイリアス。
func (r *RandNum) Int() int {
	return r.rand.Int()
}

// Intn returns a non-negative pseudo-random int in the half-open interval [0,n).
// [0,n)の範囲の非負の疑似乱数intを返す。rand.Rand.Intnのエイリアス。
func (r *RandNum) Intn(n int) int {
	return r.rand.Intn(n)
}

// Float64 returns a pseudo-random float64 in the half-open interval [0.0,1.0).
// [0.0,1.0)の範囲の疑似乱数float64を返す。rand.Rand.Float64のエイリアス。
func (r *RandNum) Float64() float64 {
	return r.rand.Float64()
}

// Float32 returns a pseudo-random float32 in the half-open interval [0.0,1.0).
// [0.0,1.0)の範囲の疑似乱数float32を返す。rand.Rand.Float32のエイリアス。
func (r *RandNum) Float32() float32 {
	return r.rand.Float32()
}
