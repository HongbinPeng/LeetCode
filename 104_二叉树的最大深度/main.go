package main

import (
	common "leetcode/common"
)

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	} else {
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		return max(left, right) + 1
	}
}
