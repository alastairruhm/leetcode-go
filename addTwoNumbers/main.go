package addTwoNumbers

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Compare ListNode
func Compare(l1 *ListNode, l2 *ListNode) bool {
	c1, c2 := scan(l1), scan(l2)
	for {
		// ok is false if there are no more values to receive and the channel is closed.
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	p := l1
	q := l2
	curr := dummyHead
	carry := 0
	for p != nil || q != nil {
		var x int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		var y int
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

// func addTwoNumbersParallel(l1 *ListNode, l2 *ListNode) *ListNode {
// 	c1, c2 := scan(l1), scan(l2)
// 	var ns []int
// 	for {
// 		// ok is false if there are no more values to receive and the channel is closed.
// 		v1, ok1 := <-c1
// 		v2, ok2 := <-c2
// 		if ok1 && ok2 {
// 			ns = append(ns, v1+v2)
// 		}
// 		if !ok1 && !ok2 {
// 			break
// 		}
// 	}
// 	size := len(ns)

// 	for index, digit := range ns {
// 		if digit >= 10 {
// 			if index+1 == size {
// 				ns[index+1] = 1
// 			} else {
// 				ns[index+1] = ns[index+1] + 1
// 			}
// 			ns[index] = ns[index] - 10
// 		}
// 	}

// 	headNode := &ListNode{Val: ns[0]}
// 	ch := getLinkNode(ns[1:])
// 	go link(headNode, ch)

// 	return headNode
// }

// func getLinkNode(ns []int) chan *ListNode {
// 	c := make(chan *ListNode)
// 	go func() {
// 		var wg sync.WaitGroup
// 		wg.Add(len(ns))
// 		for _, v := range ns {
// 			go func(v int) {
// 				defer wg.Done()
// 				c <- &ListNode{Val: v}
// 			}(v)
// 		}
// 		wg.Wait()
// 		close(c)
// 	}()
// 	return c
// }

func scanListNode(l *ListNode, ch chan int) {
	if l == nil {
		return
	}
	ch <- l.Val
	if l.Next != nil {
		scanListNode(l.Next, ch)
	}
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func scan(l *ListNode) <-chan int {
	ch := make(chan int)
	go func() {
		scanListNode(l, ch)
		close(ch)
	}()
	return ch
}

// func link(l *ListNode, ch chan *ListNode) {
// 	lnNext, ok := <-ch
// 	if ok {
// 		l.Next = lnNext
// 		link(lnNext, ch)
// 	}
// }
