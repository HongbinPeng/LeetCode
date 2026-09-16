package main

import "leetcode/common"

func rightSideView(root *common.TreeNode) []int {
	ans := make([]int, 0)
	cur, flag := 0, 0
	queue := make([]*common.TreeNode, 0)
	if root == nil {
		return ans
	} else {
		queue = append(queue, root)
		for cur <= len(queue)-1 {
			if queue[cur].Left != nil {
				queue = append(queue, queue[cur].Left)
			}
			if queue[cur].Right != nil {
				queue = append(queue, queue[cur].Right)
			}
			if cur == flag {
				ans = append(ans, queue[cur].Val)
				flag = len(queue) - 1
			}
			cur++
		}
	}
	return ans
}
