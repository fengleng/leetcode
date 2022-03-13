package remove_element

func RemoveElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var last, finder = 0, 0

	for last < len(nums) {
		for nums[last] == nums[finder] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

func RemoveElementTarget(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			if index != i {
				nums[i], nums[index] = nums[index], nums[i]
			}
			index++
		}
	}
	return index
}
