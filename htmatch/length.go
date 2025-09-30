package htmatch

import (
	"github.com/krelinga/go-howto"
	"github.com/krelinga/go-match"
	"github.com/krelinga/go-match/matchutil"
)

func Length[T any](matcher match.Matcher[int], ht howto.Length[T]) match.Matcher[T] {
	return match.MatcherFunc[T](func(v T) (matched bool, explanation string) {
		length := ht.Length(v)
		matched, e := matcher.Match(length)
		explanation = matchutil.Explain(matched, "hlmatch.Length", e)
		return
	})
}
