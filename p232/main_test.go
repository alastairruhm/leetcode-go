package p232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	assert.Equal(2, stack.Peek())
	assert.Equal(2, stack.Pop())
	assert.False(stack.Empty())
	assert.Equal(1, stack.Pop())
	assert.True(stack.Empty())
}
func TestMyQueue(t *testing.T) {
	assert := assert.New(t)
	queue := Constructor()
	queue.Push(1)
	// fmt.Printf("%+v\n", queue)
	queue.Push(2)
	// fmt.Printf("%+v\n", queue)
	assert.Equal(1, queue.Peek()) // returns 1
	// fmt.Printf("%+v\n", queue)
	assert.Equal(1, queue.Pop()) // returns 1
	assert.False(queue.Empty())  // returns false
}
