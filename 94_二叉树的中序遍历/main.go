package main

import "leetcode/common"

func inorderTraversal(root *common.TreeNode) []int {
	st := make([]*common.TreeNode, 0)
	cur := root
	ans := make([]int, 0)
	for cur != nil || len(st) > 0 {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
		} else {
			cur = st[len(st)-1]
			ans = append(ans, cur.Val)
			st = st[0 : len(st)-1]
			cur = cur.Right
		}
	}
	return ans
}
