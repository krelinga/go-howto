package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestStringLike(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.StringLike[string], typemap.ToString[string]](t)
		assertImplements[typemap.StringLike[string], typemap.Compare[string]](t)
		assertImplements[typemap.StringLike[string], typemap.Order[string]](t)
	})
}

func TestString(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.String, typemap.ToString[string]](t)
		assertImplements[typemap.String, typemap.Compare[string]](t)
		assertImplements[typemap.String, typemap.Order[string]](t)
	})
}
