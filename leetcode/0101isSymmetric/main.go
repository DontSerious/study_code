package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || recur(root.Left, root.Right)
}

func recur(L *TreeNode, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}

	if L == nil || R == nil || L.Val != R.Val {
		return false
	}

	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}
