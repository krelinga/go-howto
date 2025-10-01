package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestStringLike(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		type myString string
		type Got = typemap.ForStringLike[myString]
		assertImplements[Got, typemap.String[myString]](t)
		assertImplements[Got, typemap.Compare[myString]](t)
		assertImplements[Got, typemap.Order[myString]](t)
		assertImplements[Got, typemap.Length[myString]](t)
		assertImplements[Got, typemap.HasKey[myString, int]](t)
		assertImplements[Got, typemap.GetValue[myString, int, rune]](t)
		assertImplements[Got, typemap.AllKeys[myString, int]](t)
		assertImplements[Got, typemap.AllValues[myString, rune]](t)
		assertImplements[Got, typemap.AllKeysAndValues[myString, int, rune]](t)
	})
}

func TestString(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForString, typemap.String[string]](t)
		assertImplements[typemap.ForString, typemap.Compare[string]](t)
		assertImplements[typemap.ForString, typemap.Order[string]](t)
		assertImplements[typemap.ForString, typemap.Length[string]](t)
		assertImplements[typemap.ForString, typemap.HasKey[string, int]](t)
		assertImplements[typemap.ForString, typemap.GetValue[string, int, rune]](t)
		assertImplements[typemap.ForString, typemap.AllKeys[string, int]](t)
		assertImplements[typemap.ForString, typemap.AllValues[string, rune]](t)
		assertImplements[typemap.ForString, typemap.AllKeysAndValues[string, int, rune]](t)
	})
}
