package main

import "leetcode/common"

func swapPairs(head *common.ListNode) *common.ListNode {
	var reversefunc func(head *common.ListNode, n int) *common.ListNode
	reversefunc = func(head *common.ListNode, n int) *common.ListNode {
		n--
		cur := head.Next.Next
		pre := head.Next
		for n > 0 {
			pre.Next = cur.Next
			cur.Next = head.Next
			head.Next = cur
			cur = pre.Next
			n--
		}
		return pre
	}
	first := &common.ListNode{}
	first.Next = head
	nodenum := 0
	cur := head
	for cur != nil {
		nodenum++
		cur = cur.Next
	}
	ans := first
	for ; nodenum >= 2; nodenum -= 2 {
		first = reversefunc(first, 2)
	}
	return ans.Next
}
