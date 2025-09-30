package typemap

import (
	"cmp"
)

type ForStringLike[T ~string] struct {
	CustomToString func(string) string
}

func (s ForStringLike[T]) ToString(v string) string {
	if s.CustomToString != nil {
		return s.CustomToString(v)
	}
	return v
}

func (s ForStringLike[T]) Compare(a, b string) bool {
	return a == b
}

func (s ForStringLike[T]) Order(a, b string) int {
	return cmp.Compare(a, b)
}

func (s ForStringLike[T]) Length(v string) int {
	return len(v)
}

// TODO: implement the following on String:
// - HasKey(string, int)
// - GetValue(string, int)
// - AllKeys(string)
// - AllValues(string)
// - AllKeysAndValues(string)

type ForString = ForStringLike[string]
