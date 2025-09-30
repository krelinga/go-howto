package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestSliceLike(t *testing.T) {
	type TestSlice = []string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.SliceLike[TestSlice, string], typemap.ToString[TestSlice]](t)
		assertImplements[typemap.SliceLike[TestSlice, string], typemap.Length[TestSlice]](t)
		assertImplements[typemap.SliceLike[TestSlice, string], typemap.HasKey[TestSlice, int]](t)
	})
}

func TestSlice(t *testing.T) {
	type TestSlice = []string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.Slice[string], typemap.ToString[TestSlice]](t)
		assertImplements[typemap.Slice[string], typemap.Length[TestSlice]](t)
		assertImplements[typemap.Slice[string], typemap.HasKey[TestSlice, int]](t)
	})
}
