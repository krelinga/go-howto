package typemap

import (
	"cmp"
	"fmt"
)

type IntegerType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer[T IntegerType] struct {
	ToStringFunc[T]
}

func (i Integer[T]) ToString(v T) string {
	if i.ToStringFunc != nil {
		return i.ToStringFunc(v)
	}
	return fmt.Sprintf("%v", v)
}

func (i Integer[T]) Compare(a, b T) bool {
	return a == b
}

func (i Integer[T]) Order(a, b T) int {
	return cmp.Compare[T](a, b)
}

type Int = Integer[int]
