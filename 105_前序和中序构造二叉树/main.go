package main

import "leetcode/common"

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	inmap := make(map[int]int)
	for idx, val := range inorder {
		inmap[val] = idx
	}
	var buildBtree func(preorder []int, inorder []int, pl, pr, il, ir int) *common.TreeNode
	buildBtree = func(preorder, inorder []int, pl, pr, il, ir int) *common.TreeNode {
		if pl > pr {
			return nil
		}
		rootVal := preorder[pl]
		root := &common.TreeNode{Val: preorder[pl]}
		left := buildBtree(preorder, inorder, pl+1, pl+inmap[rootVal]-il, il, inmap[rootVal]-1)
		right := buildBtree(preorder, inorder, pl+inmap[rootVal]-il+1, pr, inmap[rootVal]+1, ir)
		root.Left = left
		root.Right = right
		return root
	}
	return buildBtree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}
