package p001

import "testing"

var (
	nums   = []int{4, 6, 7, 8, 9, 1, 3, 11}
	target = 12
)

func benchmarkAddTwo(AddTwo func([]int, int) []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddTwo(nums, target)
	}
}

func BenchmarkAddTwo_1(b *testing.B) { benchmarkAddTwo(twoSum, b) }
