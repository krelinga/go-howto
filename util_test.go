package typemap_test

import (
	"reflect"
	"testing"
)

func assertImplements[Got, Want any](t *testing.T) {
	var gotZero Got
	_, ok := any(gotZero).(Want)
	if !ok {
		gotName := reflect.TypeFor[Got]().Name()
		wantName := reflect.TypeFor[Want]().Name()
		t.Errorf("Type %s does not implement %s", gotName, wantName)
	}
}
