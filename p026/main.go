package p026

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	symbol := nums[0]
	index := 1

	for i := range nums {
		if symbol != nums[i] {
			nums[index] = nums[i]
			symbol = nums[i]
			index = index + 1
		}
	}

	return index
}
