package main

/* 
	从上到下打印二叉树 2
	从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// bfs队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := []int{}
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}