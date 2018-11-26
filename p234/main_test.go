package p234

import (
	"testing"

	"github.com/alastairruhm/leetcode-go/datastructure"

	"github.com/stretchr/testify/assert"
)

var GenSinglylinkedList = datastructure.GenSinglylinkedList

func TestIsPalindrome(t *testing.T) {
	assert := assert.New(t)
	assert.True(isPalindrome(GenSinglylinkedList([]int{1, 2, 2, 1})))
	assert.True(isPalindrome(GenSinglylinkedList([]int{1, 2, 3, 2, 1})))
	assert.False(isPalindrome(GenSinglylinkedList([]int{1, 2, 3, 4, 5})))
	assert.True(isPalindrome(GenSinglylinkedList([]int{})))
	assert.True(isPalindrome(GenSinglylinkedList([]int{1})))
}
