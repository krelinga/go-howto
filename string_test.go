package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestStringLike(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForStringLike[string], typemap.String[string]](t)
		assertImplements[typemap.ForStringLike[string], typemap.Compare[string]](t)
		assertImplements[typemap.ForStringLike[string], typemap.Order[string]](t)
	})
}

func TestString(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForString, typemap.String[string]](t)
		assertImplements[typemap.ForString, typemap.Compare[string]](t)
		assertImplements[typemap.ForString, typemap.Order[string]](t)
	})
}
