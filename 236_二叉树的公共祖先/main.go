package main

import "leetcode/common"

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	} else {
		left := lowestCommonAncestor(root.Left, p, q)
		right := lowestCommonAncestor(root.Right, p, q)
		if left == nil && right == nil {
			return nil
		} else if left == nil && right != nil {
			return right
		} else if left != nil && right != nil {
			return root
		}
		return left
	}
}

// lowestCommonAncestorIter 非递归版：父指针 + 向上回溯。
// 把每个节点到 root 的祖先链看作一条"链表"，找 p、q 两条链第一次汇合的节点（相交链表思路）。
func lowestCommonAncestorIter(root, p, q *common.TreeNode) *common.TreeNode {
	parent := map[*common.TreeNode]*common.TreeNode{root: nil}
	cur := root
	treestack := make([]*common.TreeNode, 0)
	treestack = append(treestack, cur)
	for len(treestack) > 0 {
		node := treestack[len(treestack)-1]
		treestack = treestack[:len(treestack)-1]
		if node.Left != nil {
			parent[node.Left] = node
			treestack = append(treestack, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			treestack = append(treestack, node.Right)
		}
	}
	mpp := map[*common.TreeNode]struct{}{}
	temp := p
	for temp != nil {
		mpp[temp] = struct{}{}
		temp = parent[temp]
	}
	tempq := q
	for tempq != nil {
		if _, ok := mpp[tempq]; ok {
			return tempq
		}
		tempq = parent[tempq]
	}
	return nil
}
