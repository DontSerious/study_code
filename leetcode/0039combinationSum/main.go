package main

import (
	"fmt"
	"slices"
)

func main() {
	// fmt.Printf("combinationSum([]int{2, 3, 6, 7}, 7): %v\n", combinationSum([]int{2, 3, 6, 7}, 7))
	// fmt.Printf("combinationSum([]int{2, 3, 5}, 8): %v\n", combinationSum([]int{2, 3, 5}, 8))
	fmt.Printf("combinationSum([]int{8,7,4,3}, 11)): %v\n", combinationSum([]int{8, 7, 4, 3}, 11))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	cur := []int{}
	// make sure candidates order esc
	slices.Sort(candidates)

	var dfs func(int, int)
	dfs = func(remain int, start int) {
		if remain == 0 {
			arr := make([]int, len(cur))
			copy(arr, cur)
			res = append(res, arr)
			return
		}
		// start form startIndex, avoid repeat subset
		for i := start; i < len(candidates); i++ {
			tmp := remain - candidates[i]
			// if the minimum of subset is bigger than remain, cut the rest loop.
			if tmp < 0 {
				break
			}
			cur = append(cur, candidates[i])
			dfs(tmp, i)
			// rollback
			cur = cur[:len(cur)-1]
		}
	}

	dfs(target, 0)

	return res
}

// func combinationSum(candidates []int, target int) [][]int {
// 	res := [][]int{}
// 	tmp := []int{}
// 	n := len(candidates)

// 	var dep func(int, int)
// 	dep = func(target int, index int) {
// 		if index == n {
// 			return
// 		}
// 		if target == 0 {
// 			res = append(res, append([]int(nil), tmp...))
// 			return
// 		}
// 		dep(target, index+1)
// 		if target-candidates[index] >= 0 {
// 			tmp = append(tmp, candidates[index])
// 			dep(target-candidates[index], index)
// 			tmp = tmp[:len(tmp)-1]
// 		}
// 	}

// 	dep(target, 0)

// 	return res
// }
