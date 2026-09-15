package main

import (
	"leetcode/common"
)

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	firstNode := new(common.ListNode)
	firstNode.Next = head
	slow, quick := firstNode, firstNode
	for i := 0; i < n; i++ {
		quick = quick.Next
	}
	for quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
	}
	slow.Next = slow.Next.Next
	return firstNode.Next
}
