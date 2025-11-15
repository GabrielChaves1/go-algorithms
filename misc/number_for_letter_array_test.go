package misc

import (
	"reflect"
	"testing"
)

func TestNumberForLetterArray(t *testing.T) {
	got := NumberForLetterArray(5)
	want := []string{"a", "b", "c", "d", "e"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}
