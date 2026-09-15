package main

import "leetcode/common"

func hasCycle(head *common.ListNode) bool {
	slow, quick := head, head
	hsCyc := false
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			hsCyc = true
			break
		}
	}
	return hsCyc
}
