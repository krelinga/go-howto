package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestMapLike(t *testing.T) {
	type TestMap = map[int]string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.MapLike[TestMap, int, string], typemap.ToString[TestMap]](t)
		assertImplements[typemap.MapLike[TestMap, int, string], typemap.Length[TestMap]](t)
		assertImplements[typemap.MapLike[TestMap, int, string], typemap.HasKey[TestMap, int]](t)
	})
}

func TestMap(t *testing.T) {
	type TestMap = map[int]string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.Map[int, string], typemap.ToString[TestMap]](t)
		assertImplements[typemap.Map[int, string], typemap.Length[TestMap]](t)
		assertImplements[typemap.Map[int, string], typemap.HasKey[TestMap, int]](t)
	})
}
