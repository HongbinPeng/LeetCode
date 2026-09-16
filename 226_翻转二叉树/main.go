package main

import "leetcode/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	} else {
		left := invertTree(root.Left)
		right := invertTree(root.Right)
		temp := left
		root.Left = right
		root.Right = temp
		return root
	}
}
