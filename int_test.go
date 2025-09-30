package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestInteger(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForInt, typemap.String[int]](t)
		assertImplements[typemap.ForInt, typemap.Compare[int]](t)
		assertImplements[typemap.ForInt, typemap.Order[int]](t)
	})
}
