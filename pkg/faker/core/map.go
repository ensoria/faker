package core

import "math/rand"

// RandMap provides methods for selecting random elements from maps.
// マップからランダムな要素を選択するメソッドを提供する構造体。
type RandMap struct {
	rand *rand.Rand
}

// NewRandMap creates a new RandMap instance with the given random source.
// 指定されたランダムソースで新しいRandMapインスタンスを作成する。
func NewRandMap(rand *rand.Rand) *RandMap {
	return &RandMap{
		rand,
	}
}

// KeyValue returns a random key and value from the given map.
// 指定されたマップからランダムなキーと値を返す。
func (r *RandMap) KeyValue(m map[any]any) (any, any) {
	return GetRandomKeyValue(r, m)
}

// KeySliceValue returns a random key and its slice value from the given map.
// 指定されたマップからランダムなキーとそのスライス値を返す。
func (r *RandMap) KeySliceValue(m map[any][]any) (any, any) {
	return GetRandomKeyValue(r, m)
}

// GetRandomKeyValue returns a random key and value from a map with any comparable key type.
// 任意のcomparableなキー型を持つマップからランダムなキーと値を返す。
func GetRandomKeyValue[K comparable, V any](r *RandMap, m map[K]V) (K, V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	key := keys[r.rand.Intn(len(keys))]
	value := m[key]
	return key, value
}
