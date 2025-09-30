package howto_test

import (
	"testing"

	"github.com/krelinga/go-howto"
)

func TestStringLike(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.StringLike[string], howto.ToString[string]](t)
		assertImplements[howto.StringLike[string], howto.Compare[string]](t)
		assertImplements[howto.StringLike[string], howto.Order[string]](t)
	})
}

func TestString(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.String, howto.ToString[string]](t)
		assertImplements[howto.String, howto.Compare[string]](t)
		assertImplements[howto.String, howto.Order[string]](t)
	})
}
