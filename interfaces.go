package howto

import "iter"

type ToString[T any] interface {
	ToString(T) string
}

type ToStringFunc[T any] func(T) string

func (f ToStringFunc[T]) ToString(v T) string {
	return f(v)
}

type Compare[T any] interface {
	Compare(T, T) bool
}

type CompareFunc[T any] func(T, T) bool

func (f CompareFunc[T]) Compare(a, b T) bool {
	return f(a, b)
}

type Order[T any] interface {
	Order(T, T) int
}

type OrderFunc[T any] func(T, T) int

func (f OrderFunc[T]) Order(a, b T) int {
	return f(a, b)
}

type Length[T any] interface {
	Length(T) int
}

type LengthFunc[T any] func(T) int

func (f LengthFunc[T]) Length(v T) int {
	return f(v)
}

type HasKey[T any, K any] interface {
	HasKey(T, K) bool
}

type HasKeyFunc[T any, K any] func(T, K) bool

func (f HasKeyFunc[T, K]) HasKey(v T, k K) bool {
	return f(v, k)
}

type IsNil interface {
	IsNil() bool
}

type IsNilFunc[T any] func(T) bool

func (f IsNilFunc[T]) IsNil(v T) bool {
	return f(v)
}

type GetValue[T any, K any, V any] interface {
	GetValue(T, K) (V, bool)
}

type GetValueFunc[T any, K any, V any] func(T, K) (V, bool)

func (f GetValueFunc[T, K, V]) GetValue(v T, k K) (V, bool) {
	return f(v, k)
}

type AllKeys[T any, K any] interface {
	AllKeys(T) iter.Seq[K]
}

type AllKeysFunc[T any, K any] func(T) iter.Seq[K]

func (f AllKeysFunc[T, K]) AllKeys(v T) iter.Seq[K] {
	return f(v)
}

type AllValues[T any, V any] interface {
	AllValues(T) iter.Seq[V]
}

type AllValuesFunc[T any, V any] func(T) iter.Seq[V]

func (f AllValuesFunc[T, V]) AllValues(v T) iter.Seq[V] {
	return f(v)
}

type AllKeysAndValues[T any, K any, V any] interface {
	AllKeysAndValues(T) iter.Seq2[K, V]
}

type AllKeysAndValuesFunc[T any, K any, V any] func(T) iter.Seq2[K, V]

func (f AllKeysAndValuesFunc[T, K, V]) AllKeysAndValues(v T) iter.Seq2[K, V] {
	return f(v)
}
