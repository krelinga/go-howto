package typemap

import "iter"

type ForStringLike[T ~string] struct {
	StringFunc[T]
	DefaultCompare[T]
	DefaultOrder[T]
}

func (s ForStringLike[T]) Length(v T) int {
	return len(v)
}

func (s ForStringLike[T]) HasKey(v T, key int) bool {
	return key >= 0 && key < len(v)
}

func (s ForStringLike[T]) GetValue(v T, key int) (rune, bool) {
	if !s.HasKey(v, key) {
		return 0, false
	}
	return rune(v[key]), true
}

func (s ForStringLike[T]) AllKeys(v T) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range v {
			if !yield(i) {
				return
			}
		}
	}
}

func (s ForStringLike[T]) AllValues(v T) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range v {
			if !yield(r) {
				return
			}
		}
	}
}

func (s ForStringLike[T]) AllKeysAndValues(v T) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range v {
			if !yield(i, r) {
				return
			}
		}
	}
}

type ForString = ForStringLike[string]
