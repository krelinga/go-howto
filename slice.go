package howto

import (
	"fmt"
	"iter"
	"slices"
)

type SliceLike[T ~[]E, E any] struct {
	ToStringFunc[T]
}

func (s SliceLike[T, E]) ToString(v T) string {
	if s.ToStringFunc != nil {
		return s.ToStringFunc(v)
	}
	return fmt.Sprintf("%v", v)
}

func (s SliceLike[T, E]) Length(v T) int {
	return len(v)
}

func (s SliceLike[T, E]) HasKey(v T, k int) bool {
	return k >= 0 && k < len(v)
}

func (s SliceLike[T, E]) IsNil(v T) bool {
	return v == nil
}

func (s SliceLike[T, E]) GetValue(v T, k int) (E, bool) {
	var zero E
	if k < 0 || k >= len(v) {
		return zero, false
	}
	return v[k], true
}

func (s SliceLike[T, E]) AllKeys(v T) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range v {
			if !yield(i) {
				return
			}
		}
	}
}

func (s SliceLike[T, E]) AllValues(v T) iter.Seq[E] {
	return slices.Values(v)
}

func (s SliceLike[T, E]) AllKeysAndValues(v T) iter.Seq2[int, E] {
	return slices.All(v)
}

type Slice[E any] = SliceLike[[]E, E]
