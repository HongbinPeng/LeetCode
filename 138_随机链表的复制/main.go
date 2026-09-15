package main

import "leetcode/common"

func copyRandomList(head *common.Node) *common.Node {
	first := &common.Node{}
	ans := first
	jilumap := make(map[*common.Node]*common.Node)
	cur := head
	for cur != nil {
		temp := &common.Node{Val: cur.Val}
		first.Next = temp
		jilumap[cur] = temp
		cur = cur.Next
		first = first.Next
	}
	first = ans.Next
	cur = head
	for first != nil {
		first.Random = jilumap[cur.Random]
		first = first.Next
		cur = cur.Next
	}
	return ans.Next
}
