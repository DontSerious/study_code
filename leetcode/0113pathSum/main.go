package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	res := []int{}
	var find func(*TreeNode, int)
	find = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil && sum+root.Val == targetSum {
			res = append(res, root.Val)
			// 这里res的底层数组是同一个，因为ans保存的是地址，他们地址相同，如果在后面重复append或进行其他改动，那么会改动同地址的数组，所以在append到ans里时需要新建一个数组添加进去
			ans = append(ans, append([]int{}, res...))
			res = res[:len(res)-1]
			return
		}

		res = append(res, root.Val)
		find(root.Left, sum+root.Val)
		res = res[:len(res)-1]

		res = append(res, root.Val)
		find(root.Right, sum+root.Val)
		res = res[:len(res)-1]
	}
	find(root, 0)
	return ans
}
