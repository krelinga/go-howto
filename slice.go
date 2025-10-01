package typemap

import (
	"fmt"
	"iter"
	"slices"
)

type ForSliceLike[T ~[]E, E any] struct {
	StringFunc[T]
}

func (s ForSliceLike[T, E]) ToString(v T) string {
	if s.StringFunc != nil {
		return s.StringFunc(v)
	}
	return fmt.Sprintf("%v", v)
}

func (s ForSliceLike[T, E]) Length(v T) int {
	return len(v)
}

func (s ForSliceLike[T, E]) HasKey(v T, k int) bool {
	return k >= 0 && k < len(v)
}

func (s ForSliceLike[T, E]) IsNil(v T) bool {
	return v == nil
}

func (s ForSliceLike[T, E]) GetValue(v T, k int) (E, bool) {
	var zero E
	if k < 0 || k >= len(v) {
		return zero, false
	}
	return v[k], true
}

func (s ForSliceLike[T, E]) AllKeys(v T) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range v {
			if !yield(i) {
				return
			}
		}
	}
}

func (s ForSliceLike[T, E]) AllValues(v T) iter.Seq[E] {
	return slices.Values(v)
}

func (s ForSliceLike[T, E]) AllKeyValues(v T) iter.Seq2[int, E] {
	return slices.All(v)
}

type ForSlice[E any] = ForSliceLike[[]E, E]
