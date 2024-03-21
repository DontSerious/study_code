package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return recur(root) >= 0
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	if abs(left-right) < 2 {
		return max(left, right) + 1
	} else {
		return -1
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
