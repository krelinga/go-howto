package howto_test

import (
	"testing"

	"github.com/krelinga/go-howto"
)

func TestSliceLike(t *testing.T) {
	type TestSlice = []string
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.SliceLike[TestSlice, string], howto.ToString[TestSlice]](t)
		assertImplements[howto.SliceLike[TestSlice, string], howto.Length[TestSlice]](t)
		assertImplements[howto.SliceLike[TestSlice, string], howto.HasKey[TestSlice, int]](t)
	})
}

func TestSlice(t *testing.T) {
	type TestSlice = []string
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.Slice[string], howto.ToString[TestSlice]](t)
		assertImplements[howto.Slice[string], howto.Length[TestSlice]](t)
		assertImplements[howto.Slice[string], howto.HasKey[TestSlice, int]](t)
	})
}
