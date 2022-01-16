package two_sorted_array

import "testing"

func TestFindMedia(t *testing.T) {
	num1 := []int{1, 34, 67, 444}
	num2 := []int{45, 67, 88, 888}
	v := findMedianSortedArrays(num1, num2)
	t.Log(v)
}
