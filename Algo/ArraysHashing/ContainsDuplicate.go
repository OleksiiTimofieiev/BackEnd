package arrayshashing

func ContainsDuplicate(nums []int) bool {
	list := make(map[int]int)

	list[nums[0]] = nums[0]

	for i := 1; i < len(nums); i++ {
		val, ok := list[nums[i]]

		if ok {
			return true
		} else {
			list[nums[i]] = val
		}
	}
	return false
}
