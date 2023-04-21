package main

/* 
	对称的二叉树
	请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recur(root.Left, root.Right)
}

func recur(left *TreeNode, right *TreeNode) bool {
	// 同时到底则相同
	if left == nil && right == nil {
		return true
	}
	// 不同时间到底或者值不同则不对称
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	// 递归比较
	return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}