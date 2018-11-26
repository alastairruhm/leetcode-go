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

	slow := head
	fast := head.Next

	// 确定中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	cur := slow.Next
	slow.Next = nil
	var temp *ListNode

	// 反转后半链表
	for cur != nil {
		q := cur.Next
		// 反转链表
		cur.Next = temp

		temp = cur
		// 向前遍历
		cur = q
	}

	// 前半链表和后半链表比较
	for temp != nil && head != nil { // 不需要管链表元素是奇数还是偶数
		if temp.Val != head.Val {
			return false
		}
		temp = temp.Next
		head = head.Next
	}

	return true
}

// FindMiddle 找到链表中点
// 思路：利用快慢指针
// func FindMiddle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	slow := head
// 	fast := head

// 	for fast != nil && fast.Next != nil {
// 		slow := slow.Next
// 		fast := fast.Next.Next
// 	}

// 	return slow
// }
