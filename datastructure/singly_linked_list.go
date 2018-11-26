package datastructure

import (
	"strconv"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// GenSinglylinkedList generates singly-linked list
func GenSinglylinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return &ListNode{}
	}

	head := &ListNode{
		Val:  vals[len(vals)-1],
		Next: nil,
	}

	for i := len(vals) - 2; i >= 0; i-- {
		n := &ListNode{vals[i], nil}
		n.Next = head
		head = n
	}
	return head
}

func (head *ListNode) String() string {
	vals := []string{}
	for head != nil {
		vals = append(vals, strconv.Itoa(head.Val))
		head = head.Next
	}
	return strings.Join(vals, "->")
}
