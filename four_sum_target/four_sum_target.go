package four_sum_target

import "sort"

func FourSum(nums []int, target int) [][]int {
	var res = make([][]int, 0)
	var counter = make(map[int]int)

	for _, num := range nums {
		counter[num] += 1
	}
	var uniqNums = make([]int, 0, len(counter))
	for num, _ := range counter {
		uniqNums = append(uniqNums, num)
	}

	sort.Ints(uniqNums)

	uniqNumsLen := len(uniqNums)

	for i := 0; i < uniqNumsLen; i++ {
		if counter[uniqNums[i]] >= 4 && uniqNums[i]*4 == target {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < uniqNumsLen; j++ {
			if uniqNums[i]*3+uniqNums[j] == target && counter[uniqNums[i]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if uniqNums[j]*3+uniqNums[i] == target && counter[uniqNums[j]] >= 3 {
				res = append(res, []int{uniqNums[j], uniqNums[j], uniqNums[j], uniqNums[i]})
			}
			if uniqNums[i]*2+uniqNums[j]*2 == target && counter[uniqNums[i]] >= 2 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}

			for k := j + 1; k < uniqNumsLen; k++ {
				if uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target && counter[uniqNums[i]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[j]*2+uniqNums[i]+uniqNums[k] == target && counter[uniqNums[j]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[k]*2+uniqNums[i]+uniqNums[j] == target && counter[uniqNums[k]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[c] >= 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}

	}
	return res
}
