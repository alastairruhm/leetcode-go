package p234

import (
	"github.com/alastairruhm/leetcode-go/datastructure"
)

// ListNode Definition for singly-linked list.
type ListNode = datastructure.ListNode

func isPalindrome(head *ListNode) bool {
	// 边界条件
	if head == nil || head.Next == nil {
		return true
	}

	var prev *ListNode
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		// 快指针每次走2格，
		fast = fast.Next.Next
		// 慢指针每次走1格
		next := slow.Next
		// 翻转
		slow.Next = prev
		prev = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true
}
