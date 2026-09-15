package main

import (
	"leetcode/common"
)

func sortList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	quick, slow := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}
	cur := slow.Next
	slow.Next = nil
	right := sortList(cur)
	left := sortList(head)
	sortnode := merge(left, right)
	return sortnode
}
func merge(l1, l2 *common.ListNode) *common.ListNode {
	temp := &common.ListNode{}
	ans := temp
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return ans.Next
}
