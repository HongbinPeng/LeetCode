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
	// 1. 用栈 DFS 遍历，记录每个节点的父节点，直到同时找到 p、q
	parent := map[*common.TreeNode]*common.TreeNode{}
	parent[root] = nil

	stack := []*common.TreeNode{root}
	var foundP, foundQ bool
	for !foundP || !foundQ {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == p {
			foundP = true
		}
		if node == q {
			foundQ = true
		}
		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
	}

	// 2. 把 p 的整条祖先链存入集合
	ancestors := map[*common.TreeNode]bool{}
	for cur := p; cur != nil; cur = parent[cur] {
		ancestors[cur] = true
	}

	// 3. 从 q 向上走，第一个出现在 p 祖先链里的节点就是最近公共祖先
	for cur := q; cur != nil; cur = parent[cur] {
		if ancestors[cur] {
			return cur
		}
	}
	return nil
}
