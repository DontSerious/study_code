package main

import "fmt"

func main() {
	fmt.Printf("pivotIndex([]int{1, 7, 3, 6, 5, 6}): %v\n", pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Printf("pivotIndex([]int{2, 1, -1}): %v\n", pivotIndex([]int{2, 1, -1}))
}

func pivotIndex(nums []int) int {
	lSum, rSum := 0, 0
	for i := range nums {
		rSum += nums[i]
	}

	for i := range nums {
		rSum -= nums[i]
		if lSum == rSum {
			return i
		}
		lSum += nums[i]
	}

	return -1
}
