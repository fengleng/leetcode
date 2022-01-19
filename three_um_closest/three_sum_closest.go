package three_um_closest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, res, diff := len(nums), 0, math.MaxInt32
	for i := 0; i < n-2; i++ {
		if i > 1 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if int(math.Abs(float64(sum-target))) < diff {
				res = sum
				diff = int(math.Abs(float64(sum - target)))
			}
			if res == target {
				return res
			}
			if res > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
