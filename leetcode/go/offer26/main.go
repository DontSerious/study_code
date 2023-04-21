package main

/*
	树的子结构
	输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

	B是A的子结构， 即 A中有出现和B相同的结构和节点值。
	例如:
	给定的树 A:

		 3

		/ \

	   4   5

	 / \

	1   2
	给定的树 B：

	    4

	  /

	1
	返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	// 递归左子树和右子树进行比较，但凡一个对的上就true
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A *TreeNode, B *TreeNode) bool {
	// 当节点 B 为空：说明树 B 已匹配完成（越过叶子节点）
	if B == nil {
		return true
	}

	// A 为空时说明已经越过A叶子节点，即匹配失败
	// AB值不同即匹配失败
	if A == nil || A.Val != B.Val {
		return false
	}

	// 开启递归检查左右子树
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}