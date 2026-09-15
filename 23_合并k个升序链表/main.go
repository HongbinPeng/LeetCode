package main

import (
	"leetcode/common"
)

func mergeKLists(lists []*common.ListNode) *common.ListNode {
	var recurcivesort func(lists []*common.ListNode, left, right int) *common.ListNode
	mergenode := func(l1, l2 *common.ListNode) *common.ListNode {
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
	recurcivesort = func(lists []*common.ListNode, left, right int) *common.ListNode {
		if len(lists) == 0 {
			return nil
		}
		if left == right {
			return lists[left]
		}
		mid := (left + right) / 2
		leftnode := recurcivesort(lists, left, mid)
		rightnode := recurcivesort(lists, mid+1, right)
		return mergenode(leftnode, rightnode)
	}
	return recurcivesort(lists, 0, len(lists)-1)
}
