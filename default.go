package typemap

import "cmp"

type DefaultCompare[T comparable] struct{}

func (d DefaultCompare[T]) Compare(a, b T) bool {
	return a == b
}

type DefaultOrder[T cmp.Ordered] struct{}

func (d DefaultOrder[T]) Order(a, b T) int {
	return cmp.Compare(a, b)
}