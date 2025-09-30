package typemap

type FloatLike[T ~float32 | ~float64] struct {
	StringFunc[T]
	DefaultCompare[T]
	DefaultOrder[T]
}

type Float32 = FloatLike[float32]
type Float64 = FloatLike[float64]
