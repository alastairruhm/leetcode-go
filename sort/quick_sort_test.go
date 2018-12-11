package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	x := []int{1}
	QuickSort(x)
	assert.Equal([]int{1}, x)

	x = []int{5, 4, 2, 3, 1}
	QuickSort(x)
	assert.Equal([]int{1, 2, 3, 4, 5}, x)
}
