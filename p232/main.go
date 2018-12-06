package p232

import "fmt"

type MyQueue struct {
	StackNew *Stack
	StackOld *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{&Stack{}, &Stack{}}
}

func (m MyQueue) String() string {
	return fmt.Sprintf("t1:{%v}, t2:{%v}", m.StackNew.items, m.StackOld.items)
}

func (this *MyQueue) turnover() {
	if this.StackOld.Empty() { // 只有空的时候才要倒腾数据
		for !this.StackNew.Empty() {
			this.StackOld.Push(this.StackNew.Pop())
		}
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.StackNew.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.turnover()
	return this.StackOld.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.turnover()
	return this.StackOld.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	this.turnover()
	return this.StackOld.Empty()
}

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() int {
	top := s.Peek()
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
