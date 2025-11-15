package arrays

import (
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	want := 49

	got := ContainerWithMostWater(input)

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
