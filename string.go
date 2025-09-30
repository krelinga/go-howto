package howto

import (
	"cmp"
)

type StringLike[T ~string] struct {
	CustomToString func(string) string
}

func (s StringLike[T]) ToString(v string) string {
	if s.CustomToString != nil {
		return s.CustomToString(v)
	}
	return v
}

func (s StringLike[T]) Compare(a, b string) bool {
	return a == b
}

func (s StringLike[T]) Order(a, b string) int {
	return cmp.Compare(a, b)
}

func (s StringLike[T]) Length(v string) int {
	return len(v)
}

// TODO: implement the following on String:
// - HasKey(string, int)
// - GetValue(string, int)
// - AllKeys(string)
// - AllValues(string)
// - AllKeysAndValues(string)

type String = StringLike[string]
