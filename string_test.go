package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestStringLike(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		type myString string
		assertImplements[typemap.ForStringLike[myString], typemap.String[myString]](t)
		assertImplements[typemap.ForStringLike[myString], typemap.Compare[myString]](t)
		assertImplements[typemap.ForStringLike[myString], typemap.Order[myString]](t)
		assertImplements[typemap.ForStringLike[myString], typemap.Length[myString]](t)
	})
}

func TestString(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForString, typemap.String[string]](t)
		assertImplements[typemap.ForString, typemap.Compare[string]](t)
		assertImplements[typemap.ForString, typemap.Order[string]](t)
		assertImplements[typemap.ForString, typemap.Length[string]](t)
	})
}
