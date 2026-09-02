/**
 * Definition for singly-linked list.
 */
package main

import "leetcode/common"

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	headA_ := headA
	headB_ := headB
	for headA != headB {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
		if headA == nil && headB == nil {
			return nil
		}
		if headA == nil {
			headA = headB_
		}
		if headB == nil {
			headB = headA_
		}
	}
	return headA
}
