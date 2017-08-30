package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func main() {
	ints := []int{7, 2, 4}
	var ints2 [3]int
	copy(ints2[:], ints)
	sort.Ints(ints)
	fmt.Printf("%v\n", ints)
	fmt.Printf("%v\n", ints2)
	c := sort.Search(len(ints2), func(i int) bool { return ints2[i] >= 4 })
	fmt.Printf("%d\n", c)
}
