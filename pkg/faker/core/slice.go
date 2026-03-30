package core

import (
	"math/rand"

	"github.com/ensoria/gofake/pkg/faker/common/log"
)

// RandSlice provides methods for selecting random elements from slices.
// スライスからランダムな要素を選択するメソッドを提供する構造体。
type RandSlice struct {
	rand *rand.Rand
}

// NewRandSlice creates a new RandSlice instance with the given random source.
// 指定されたランダムソースで新しいRandSliceインスタンスを作成する。
func NewRandSlice(rand *rand.Rand) *RandSlice {
	return &RandSlice{
		rand,
	}
}

// StrElem returns a random string element from the given slice.
// 指定されたスライスからランダムな文字列要素を返す。
func (r *RandSlice) StrElem(slice []string) string {
	if len(slice) == 0 {
		log.WrongUsage("Given slice is empty.", 1)
		return ""
	}
	return slice[r.rand.Intn(len(slice))]
}

// IntElem returns a random int element from the given slice.
// 指定されたスライスからランダムなint要素を返す。
func (r *RandSlice) IntElem(slice []int) int {
	if len(slice) == 0 {
		log.WrongUsage("Given slice is empty.", 1)
		return 0
	}
	return slice[r.rand.Intn(len(slice))]
}
