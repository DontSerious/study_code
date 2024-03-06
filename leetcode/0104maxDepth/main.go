package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	res := 0
// 	queue := []*TreeNode{root}
// 	tmp_q := []*TreeNode{}
// 	for len(queue) != 0 {
// 		tmp_n := queue[0]
// 		if tmp_n.Left != nil {
// 			tmp_q = append(tmp_q, tmp_n.Left)
// 		}
// 		if tmp_n.Right != nil {
// 			tmp_q = append(tmp_q, tmp_n.Right)
// 		}
// 		queue = queue[1:]

// 		if len(queue) == 0 {
// 			res++
// 			queue = tmp_q
// 			tmp_q = []*TreeNode{}
// 		}
// 	}

// 	return res
// }
