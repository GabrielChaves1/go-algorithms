package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected %v, got %v", want, got)
	}
}
