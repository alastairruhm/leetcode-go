package medianOfTwoSortedArrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {

	var cases = []struct {
		num1     []int
		num2     []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, tt := range cases {
		result := findMedianSortedArrays(tt.num1, tt.num2)
		assert.Equal(t, tt.expected, result)
	}
}
