package main

import "leetcode/common"

func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	var reversefunc func(head *common.ListNode, n int) *common.ListNode
	reversefunc = func(head *common.ListNode, n int) *common.ListNode {
		if n == 1 {
			return head.Next
		}
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
	for ; nodenum >= k; nodenum -= k {
		first = reversefunc(first, k)
	}
	return ans.Next
}
