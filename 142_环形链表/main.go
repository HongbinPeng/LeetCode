package main

import "leetcode/common"

func detectCycle(head *common.ListNode) *common.ListNode {
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if slow == quick {
			break
		}
	}
	if quick == nil || quick.Next == nil {
		return nil
	}
	cur := head
	for cur != slow {
		cur = cur.Next
		slow = slow.Next
	}
	return cur
}
