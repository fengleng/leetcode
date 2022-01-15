package two_sum

func TwoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i, v := range arr {
		another := target - v
		if ai, ok := m[another]; ok {
			return []int{ai, i}
		}
		m[v] = v
	}
	return nil
}
