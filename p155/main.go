package p155

type MinStack struct {
	Items []Item
}

type Item struct {
	min   int
	value int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]Item{}}
}

// Push ...
func (this *MinStack) Push(x int) {
	min := x

	if len(this.Items) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.Items = append(this.Items, Item{min: min, value: x})
}

// Pop ...
func (this *MinStack) Pop() {
	if len(this.Items) == 0 {
		panic("stack is empty")
	}
	this.Items = this.Items[:len(this.Items)-1]
}

// Top ...
func (this *MinStack) Top() int {
	if len(this.Items) == 0 {
		panic("stack is empty")
	}
	return this.Items[len(this.Items)-1].value
}

// GetMin ...
func (this *MinStack) GetMin() int {
	return this.Items[len(this.Items)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
