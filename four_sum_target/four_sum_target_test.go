package four_sum_target

import "testing"

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}

	v := FourSum(nums, 0)
	t.Log(v)
}
