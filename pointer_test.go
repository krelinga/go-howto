package typemap_test

import (
	"testing"

	"github.com/krelinga/go-typemap"
)

func TestForPointer(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[typemap.ForPointer[int], typemap.Deref[int]](t)
		assertImplements[typemap.ForPointer[int], typemap.Compare[*int]](t)
		assertImplements[typemap.ForPointer[int], typemap.IsNil[*int]](t)
		assertImplements[typemap.ForPointer[int], typemap.String[*int]](t)
	})
}