package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenSinglyLinkedList(t *testing.T) {
	assert := assert.New(t)
	v := GenSinglylinkedList([]int{1, 2, 3, 4, 5})
	output := v.String()
	assert.Equal("1->2->3->4->5", output)
	v = GenSinglylinkedList([]int{1, 2, 2, 1})
	output = v.String()
	assert.Equal("1->2->2->1", output)
}
