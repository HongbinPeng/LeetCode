package main

import (
	"leetcode/common"
)

func flatten(root *common.TreeNode) {
	st := make([]*common.TreeNode, 0)
	queue := make([]*common.TreeNode, 0)
	cur := root
	if root == nil {
		return
	}
	for cur != nil || len(st) > 0 {
		if cur == nil {
			cur = st[len(st)-1]
			st = st[0 : len(st)-1]
			cur = cur.Right
			continue
		}
		queue = append(queue, cur)
		st = append(st, cur)
		cur = cur.Left
	}
	var prefix *common.TreeNode
	for idx, node := range queue {
		if idx == 0 {
			node.Left = prefix
			prefix = node
			continue
		} else {
			prefix.Right = node
			node.Left = nil
			prefix = node
		}
	}
	prefix.Left = nil
	prefix.Right = nil
}
