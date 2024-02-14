package levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	tmp_q := []*TreeNode{}
	tmp_res := []int{}
	for len(queue) > 0 {
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
			res = append(res, tmp_res)
			tmp_res = []int{}
		}
	}

	return res
}
