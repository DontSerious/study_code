package main

/* 
	从上到下打印二叉树
	从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// bfs队列
func levelOrder(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		res = append(res, node.Val)
	}

	return
}
