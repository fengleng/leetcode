package two_sorted_array

import "math"

func findMedianSortedArrays(num1, num2 []int) float64 {
	len1 := len(num1)
	len2 := len(num2)
	total := len1 + len2
	left := (total + 1) / 2
	right := (total + 1) / 2
	return (float64(findTopK(num1, 0, num2, 0, left)) + float64(findTopK(num1, 0, num2, 0, right))) / 2.0
}

func findTopK(num1 []int, i int, num2 []int, j int, k int) int {
	num1Len := len(num1)
	if i > num1Len {
		return num2[j+k-1]
	}
	num2Len := len(num2)
	if j > num2Len {
		return num1[i+k-1]
	}
	if k == 1 {
		return int(math.Min(float64(num1[i]), float64(num2[j])))
	}
	var mid1, mid2 int
	if i+k/2-1 < num1Len {
		mid1 = num1[i+k/2-1]
	} else {
		mid1 = int(math.MaxInt64)
	}

	if j+k/2-1 < num2Len {
		mid2 = num2[j+k/2-1]
	} else {
		mid2 = int(math.MaxInt64)
	}
	if mid1 < mid2 {
		return findTopK(num1, i+k/2, num2, j, k-k/2)
	}
	return findTopK(num1, i, num2, j+k/2, k-k/2)
}