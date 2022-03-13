package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	v := RemoveElement(nums)
	t.Log(v)

}
