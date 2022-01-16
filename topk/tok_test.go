package topk

import "testing"

func TestFindTopK(t *testing.T) {
	num1 := []int{1, 34, 67, 444}
	num2 := []int{45, 67, 88, 888}
	k := findTopK(num1, 0, num2, 0, 7)
	t.Log(k)
}
