package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestMapLike(t *testing.T) {
	type TestMap = map[int]string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForMapLike[TestMap, int, string], typemap.String[TestMap]](t)
		assertImplements[typemap.ForMapLike[TestMap, int, string], typemap.Length[TestMap]](t)
		assertImplements[typemap.ForMapLike[TestMap, int, string], typemap.IsNil[TestMap]](t)
		assertImplements[typemap.ForMapLike[TestMap, int, string], typemap.HasKey[TestMap, int]](t)
	})
}

func TestMap(t *testing.T) {
	type TestMap = map[int]string
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForMap[int, string], typemap.String[TestMap]](t)
		assertImplements[typemap.ForMap[int, string], typemap.Length[TestMap]](t)
		assertImplements[typemap.ForMap[int, string], typemap.IsNil[TestMap]](t)
		assertImplements[typemap.ForMap[int, string], typemap.HasKey[TestMap, int]](t)
	})
}
