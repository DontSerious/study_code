package main

import "fmt"

func main() {
	fmt.Printf("search([]int{-1, 0, 3, 5, 9, 12}, 9): %v\n", search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
