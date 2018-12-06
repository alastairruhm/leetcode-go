package p026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, removeDuplicates([]int{1, 1, 2}))
	assert.Equal(5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
