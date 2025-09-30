package typemap

type ForIntegerLike[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr] struct {
	StringFunc[T]
	DefaultCompare[T]
	DefaultOrder[T]
}

type ForInt = ForIntegerLike[int]
type ForInt8 = ForIntegerLike[int8]
type ForInt16 = ForIntegerLike[int16]
type ForInt32 = ForIntegerLike[int32]
type ForInt64 = ForIntegerLike[int64]
type ForUint = ForIntegerLike[uint]
type ForUint8 = ForIntegerLike[uint8]
type ForUint16 = ForIntegerLike[uint16]
type ForUint32 = ForIntegerLike[uint32]
type ForUint64 = ForIntegerLike[uint64]
type ForUintptr = ForIntegerLike[uintptr]
