package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestInteger(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.Int, typemap.ToString[int]](t)
		assertImplements[typemap.Int, typemap.Compare[int]](t)
		assertImplements[typemap.Int, typemap.Order[int]](t)
	})
}
