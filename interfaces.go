package typemap

import (
	"fmt"
	"iter"
)

type String[T any] interface {
	String(T) string
}

type StringFunc[T any] func(T) string

func (f StringFunc[T]) String(v T) string {
	if f == nil {
		return fmt.Sprintf("%v", v)
	}
	return f(v)
}

type Compare[T any] interface {
	Compare(T, T) bool
}

type Order[T any] interface {
	Order(T, T) int
}

type Length[T any] interface {
	Length(T) int
}

type HasKey[T any, K any] interface {
	HasKey(T, K) bool
}

type IsNil interface {
	IsNil() bool
}

type GetValue[T any, K any, V any] interface {
	GetValue(T, K) (V, bool)
}

type AllKeys[T any, K any] interface {
	AllKeys(T) iter.Seq[K]
}

type AllValues[T any, V any] interface {
	AllValues(T) iter.Seq[V]
}

type AllKeyValues[T any, K any, V any] interface {
	AllKeyValues(T) iter.Seq2[K, V]
}
