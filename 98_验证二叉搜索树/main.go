package main

import (
	"leetcode/common"
)

func isValidBST(root *common.TreeNode) bool {
	var isBST func(root *common.TreeNode, min, max *int) bool
	isBST = func(root *common.TreeNode, min, max *int) bool {
		if root == nil {
			return true
		} else {
			if min != nil && root.Val <= *min {
				return false
			}
			if max != nil && root.Val >= *max {
				return false
			}
			return isBST(root.Left, min, &root.Val) && isBST(root.Right, &root.Val, max)
		}
	}
	return isBST(root, nil, nil)
}
