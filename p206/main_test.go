package p206

import (
	"testing"

	"github.com/alastairruhm/leetcode-go/datastructure"

	"github.com/stretchr/testify/assert"
)

var GenSinglylinkedList = datastructure.GenSinglylinkedList

func TestReverseListRecursively(t *testing.T) {
	assert := assert.New(t)
	l1 := GenSinglylinkedList([]int{1, 2, 3, 4, 5})
	assert.Equal("1->2->3->4->5", l1.String())
	r1 := reverseListRecursively(l1)
	assert.Equal("5->4->3->2->1", r1.String())
}
