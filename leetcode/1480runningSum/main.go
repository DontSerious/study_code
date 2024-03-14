package main

import "fmt"

func main() {
	fmt.Printf("runningSum([]int{1, 2, 3, 4}): %v\n", runningSum([]int{1, 2, 3, 4}))
}

func runningSum(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}
