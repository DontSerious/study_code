package main

import "fmt"

func main() {
	fmt.Printf("permute([]int{1, 2, 3}): %v\n", permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, path...))
			return
		}

		for j, flag := range onPath {
			if !flag {
				path[i] = nums[j]

				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}

	dfs(0)

	return res
}
