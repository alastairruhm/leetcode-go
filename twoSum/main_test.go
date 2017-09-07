package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwo(t *testing.T) {

	var cases = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{1, 2, 3}, 9, []int{0, 0}},
	}

	for _, tt := range cases {
		result := twoSum(tt.nums, tt.target)
		assert.ObjectsAreEqualValues(result, tt.expected)
	}
}
