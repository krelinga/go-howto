package typemap

type ForPointer[T any] struct{
	StringFunc[*T]
}

func (_ ForPointer[T]) Compare(a, b *T) bool {
	return a == b
}

func (_ ForPointer[T]) IsNil(v *T) bool {
	return v == nil
}

func (_ ForPointer[T]) Deref(v *T) T {
	return *v
}
