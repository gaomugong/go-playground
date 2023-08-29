package arrayslice

import "testing"

func TestSum(t *testing.T) {

	// nums := [5]int{1, 2, 3, 4, 5}
	nums := [...]int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("%#v: want %d, got %d", nums, want, got)
	}
}
