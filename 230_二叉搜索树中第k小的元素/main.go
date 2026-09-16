package main

import "leetcode/common"

func kthSmallest(root *common.TreeNode, k int) int {
	var inorder func(root *common.TreeNode)
	ans := 0
	flag := false
	inorder = func(root *common.TreeNode) {
		if root == nil || flag {
			return
		}
		inorder(root.Left)
		k--
		if k == 0 {
			ans = root.Val
			flag = true
			return
		}
		inorder(root.Right)
	}
	inorder(root)
	return ans
}
