package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("permuteUnique([]int{1, 1, 3}): %v\n", permuteUnique([]int{1, 1, 3}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}
	n := len(nums)
	path := make([]int, n)
	used := make([]bool, n)

	var dfs func(int)
	dfs = func(depth int) {

		// 到达最低，记录答案
		if depth == n {
			res = append(res, append([]int{}, path...))
			return
		}

		for i, flag := range used {
			if !flag {

				// 剪枝
				if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
					continue
				}

				path[depth] = nums[i]
				used[i] = true
				dfs(depth + 1)
				used[i] = false
			}
		}
	}

	dfs(0)

	return res
}
