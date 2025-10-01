package typemap

import (
	"fmt"
	"iter"
	"maps"
)

type ForMapLike[T ~map[K]V, K comparable, V any] struct {
	StringFunc[T]
}

func (m ForMapLike[T, K, V]) ToString(v T) string {
	if m.StringFunc != nil {
		return m.StringFunc(v)
	}
	return fmt.Sprintf("%v", v)
}

func (m ForMapLike[T, K, V]) Length(v T) int {
	return len(v)
}

func (m ForMapLike[T, K, V]) HasKey(v T, k K) bool {
	_, ok := v[k]
	return ok
}

func (m ForMapLike[T, K, V]) IsNil(v T) bool {
	return v == nil
}

func (m ForMapLike[T, K, V]) GetValue(v T, k K) (V, bool) {
	got, ok := v[k]
	return got, ok
}

func (m ForMapLike[T, K, V]) AllKeys(v T) iter.Seq[K] {
	return maps.Keys(v)
}

func (m ForMapLike[T, K, V]) AllValues(v T) iter.Seq[V] {
	return maps.Values(v)
}

func (m ForMapLike[T, K, V]) AllKeyValues(v T) iter.Seq2[K, V] {
	return maps.All(v)
}

type ForMap[K comparable, V any] = ForMapLike[map[K]V, K, V]
