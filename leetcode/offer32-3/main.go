package main

/* 
	从上到下打印二叉树 III
	请实现一个函数按照之字形顺序打印二叉树，
	即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，
	第三行再按照从左到右的顺序打印，其他行以此类推。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	flag := true

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := []int{}
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]

			if flag {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		flag = !flag
		res = append(res, tmp)
	}

	return res
}