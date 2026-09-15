package main

import "leetcode/common"

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	head := new(common.ListNode)
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			head = head.Next
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
			head = head.Next
		}
	}
	for list1 != nil {
		head.Next = list1
		list1 = list1.Next
		head = head.Next
	}
	for list2 != nil {
		head.Next = list2
		list2 = list2.Next
		head = head.Next
	}
	return cur.Next
}
