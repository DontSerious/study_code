package main

import (
	"fmt"
)

func main() {
	fmt.Printf("rob([]int{1, 2, 3, 1}): %v\n", rob([]int{1}))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// rob-1
	_rob := func(nums []int) int {
		pre, cur := 0, 0

		for _, v := range nums {
			tmp := max(cur, pre+v)
			pre = cur
			cur = tmp
		}

		return cur
	}

	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}
