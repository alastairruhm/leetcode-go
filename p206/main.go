package p206

import (
	ds "github.com/alastairruhm/leetcode-go/datastructure"
)

// ListNode ...
type ListNode = ds.ListNode

// reverseListIteratively reverse list iteratively
func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversedHead := head
	head = head.Next
	// 处理头节点
	reversedHead.Next = nil

	for head != nil {
		// 后一个节点链到前一个节点上
		head.Next = reversedHead

	}

	return reversedHead
}

//  LeetCode No.206 反转链表 | GolangCaff - 高品质的 Golang 开发者论坛: https://golangcaff.com/articles/270
func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
