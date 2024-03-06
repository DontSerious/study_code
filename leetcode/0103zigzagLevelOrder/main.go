package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	flag := true
	queue := []*TreeNode{root}
	tmp_q := []*TreeNode{}
	tmp_res := []int{}

	for len(queue) != 0 {
		tmp_node := queue[0]
		tmp_res = append(tmp_res, tmp_node.Val)
		if tmp_node.Left != nil {
			tmp_q = append(tmp_q, tmp_node.Left)
		}
		if tmp_node.Right != nil {
			tmp_q = append(tmp_q, tmp_node.Right)
		}

		queue = queue[1:]
		if len(queue) == 0 {
			queue = tmp_q
			tmp_q = []*TreeNode{}
			if flag {
				res = append(res, tmp_res)
			} else {
				res = append(res, reverseArr(tmp_res))
			}
			flag = !flag
			tmp_res = []int{}
		}
	}

	return res
}

func reverseArr(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
