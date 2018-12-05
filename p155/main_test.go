package p155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	assert := assert.New(t)
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Equal(-3, stack.GetMin())
	assert.Equal(-3, stack.Top())
	stack.Pop()
	assert.Equal(-2, stack.GetMin())
	assert.Equal(0, stack.Top())
}
