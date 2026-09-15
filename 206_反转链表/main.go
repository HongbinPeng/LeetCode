package main

import "leetcode/common"

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	var first *common.ListNode = new(common.ListNode)
	first.Next = head
	pre := head
	cur := head.Next
	for cur != nil {
		pre.Next = cur.Next
		cur.Next = first.Next
		first.Next = cur
		cur = pre.Next
	}
	return first.Next

}
