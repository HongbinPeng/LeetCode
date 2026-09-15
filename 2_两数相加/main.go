package main

import "leetcode/common"

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	temp := 0
	head := new(common.ListNode)
	cur := head
	for l1 != nil && l2 != nil {
		curnum := (l1.Val + l2.Val + temp) % 10
		tempnode := &common.ListNode{Val: curnum, Next: nil}
		cur.Next = tempnode
		cur = tempnode
		temp = (l1.Val + l2.Val + temp) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		curnum := (l1.Val + temp) % 10
		tempnode := &common.ListNode{Val: curnum, Next: nil}
		cur.Next = tempnode
		cur = tempnode
		temp = (l1.Val + temp) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		curnum := (l2.Val + temp) % 10
		tempnode := &common.ListNode{Val: curnum, Next: nil}
		cur.Next = tempnode
		cur = tempnode
		temp = (l2.Val + temp) / 10
		l2 = l2.Next
	}
	if temp != 0 {
		cur.Next = &common.ListNode{Val: temp, Next: nil}
	}
	return head.Next

}
