package typemap

import (
	"fmt"
	"iter"
	"maps"
)

type MapLike[T ~map[K]V, K comparable, V any] struct {
	ToStringFunc[T]
}

func (m MapLike[T, K, V]) ToString(v T) string {
	if m.ToStringFunc != nil {
		return m.ToStringFunc(v)
	}
	return fmt.Sprintf("%v", v)
}

func (m MapLike[T, K, V]) Length(v T) int {
	return len(v)
}

func (m MapLike[T, K, V]) HasKey(v T, k K) bool {
	_, ok := v[k]
	return ok
}

func (m MapLike[T, K, V]) IsNil(v T) bool {
	return v == nil
}

func (m MapLike[T, K, V]) GetValue(v T, k K) (V, bool) {
	got, ok := v[k]
	return got, ok
}

func (m MapLike[T, K, V]) AllKeys(v T) iter.Seq[K] {
	return maps.Keys(v)
}

func (m MapLike[T, K, V]) AllValues(v T) iter.Seq[V] {
	return maps.Values(v)
}

func (m MapLike[T, K, V]) AllKeysAndValues(v T) iter.Seq2[K, V] {
	return maps.All(v)
}

type Map[K comparable, V any] = MapLike[map[K]V, K, V]
