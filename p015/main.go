package p015

func threeSum(nums []int) [][]int {
	r := [][]int{}
	insertSort(nums)

	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 { // 排序后，fixed 的数不能是正数
			break
		}

		if k > 0 && nums[k] == nums[k-1] { // fixed 的数字只需要一次
			continue
		}
		target := 0 - nums[k]
		i := k + 1
		j := len(nums) - 1
		for i < j {
			if nums[i]+nums[j] == target {
				r = append(r, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}
	}
	return r

}

func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		j := i - 1
		for j >= 0 {
			if nums[j] > v {
				nums[j+1] = nums[j]
			} else {
				break
			}
			j--
		}
		nums[j+1] = v
	}
}
