package typemap_test

import "testing"

func assertImplements[Got, Want any](t *testing.T) {
	var gotZero Got
	_, ok := any(gotZero).(Want)
	if !ok {
		var wantZero Want
		t.Errorf("Expected %T to implement %T", gotZero, wantZero)
	}
}
