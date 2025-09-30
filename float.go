package typemap

import (
	"cmp"
	"fmt"
)

type Float[T ~float32 | ~float64] struct {
	ToStringFunc[T]
}

func (f Float[T]) ToString(v T) string {
	if f.ToStringFunc != nil {
		return f.ToStringFunc(v)
	}
	return fmt.Sprintf("%v", v)
}

func (f Float[T]) Compare(a, b T) bool {
	return a == b
}

func (f Float[T]) Order(a, b T) int {
	return cmp.Compare[T](a, b)
}
