package three_sum_zero

import "testing"

func TestFindThreeSumZero(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(FindThreeSumZero(nums))
}
