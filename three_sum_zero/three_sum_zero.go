package three_sum_zero

import "sort"

func FindThreeSumZero(nums []int) [][]int {
	counter := make(map[int]int)
	var uniqNums []int
	var res [][]int

	for _, num := range nums {
		counter[num] += 1
	}

	for key, _ := range counter {
		uniqNums = append(uniqNums, key)
	}

	sort.Ints(uniqNums)
	lenUniqNums := len(uniqNums)

	for i := 0; i < lenUniqNums; i++ {
		if uniqNums[i]*3 == 0 && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < lenUniqNums; j++ {
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}

			if uniqNums[j]*2+uniqNums[i] == 0 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}

			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] >= 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}

	return res
}
