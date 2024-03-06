package main

import "fmt"

func main() {
	fmt.Printf("rob([]int{1, 2, 3, 1}): %v\n", rob([]int{1, 2, 3, 1}))
}

// dynamic programming
func rob(nums []int) int {
	pre, cur := 0, 0

	for _, v := range nums {
		tmp := max(cur, pre+v)
		pre = cur
		cur = tmp
	}

	return cur
}
