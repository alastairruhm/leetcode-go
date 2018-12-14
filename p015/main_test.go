package p015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal([][]int{[]int{0, 0, 0}}, threeSum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// assert.Equal(5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
