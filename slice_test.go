package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestSliceLike(t *testing.T) {
	type TestSlice = []string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForSliceLike[TestSlice, string], typemap.String[TestSlice]](t)
		assertImplements[typemap.ForSliceLike[TestSlice, string], typemap.Length[TestSlice]](t)
		assertImplements[typemap.ForSliceLike[TestSlice, string], typemap.HasKey[TestSlice, int]](t)
	})
}

func TestSlice(t *testing.T) {
	type TestSlice = []string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForSlice[string], typemap.String[TestSlice]](t)
		assertImplements[typemap.ForSlice[string], typemap.Length[TestSlice]](t)
		assertImplements[typemap.ForSlice[string], typemap.HasKey[TestSlice, int]](t)
	})
}
