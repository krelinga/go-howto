package howto_test

import (
	"testing"

	"github.com/krelinga/go-howto"
)

func TestMapLike(t *testing.T) {
	type TestMap = map[int]string
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.MapLike[TestMap, int, string], howto.ToString[TestMap]](t)
		assertImplements[howto.MapLike[TestMap, int, string], howto.Length[TestMap]](t)
		assertImplements[howto.MapLike[TestMap, int, string], howto.HasKey[TestMap, int]](t)
	})
}

func TestMap(t *testing.T) {
	type TestMap = map[int]string
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.Map[int, string], howto.ToString[TestMap]](t)
		assertImplements[howto.Map[int, string], howto.Length[TestMap]](t)
		assertImplements[howto.Map[int, string], howto.HasKey[TestMap, int]](t)
	})
}
