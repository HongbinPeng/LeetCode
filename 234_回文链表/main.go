package main

import "leetcode/common"

func isPalindrome(head *common.ListNode) bool {
	firstNode := head
	var recurecive func(node *common.ListNode) bool
	recurecive = func(node *common.ListNode) bool {
		if node != nil {
			if !recurecive(node.Next) {
				return false
			}
			if firstNode.Val != node.Val {
				return false
			}
			firstNode = firstNode.Next
		}
		return true
	}
	return recurecive(head)
}
