package howto_test

import (
	"testing"

	"github.com/krelinga/go-howto"
)

func TestInteger(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[howto.Int, howto.ToString[int]](t)
		assertImplements[howto.Int, howto.Compare[int]](t)
		assertImplements[howto.Int, howto.Order[int]](t)
	})
}