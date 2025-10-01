package typemap

type ForStringLike[T ~string] struct {
	StringFunc[T]
	DefaultCompare[T]
	DefaultOrder[T]
}

func (s ForStringLike[T]) Length(v T) int {
	return len(v)
}

// TODO: implement the following on String:
// - HasKey(string, int)
// - GetValue(string, int)
// - AllKeys(string)
// - AllValues(string)
// - AllKeysAndValues(string)

type ForString = ForStringLike[string]
