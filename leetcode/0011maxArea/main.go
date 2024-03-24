package main

import "fmt"

func main() {
	fmt.Printf("maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}): %v\n", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	res, left, right, width := 0, 0, len(height)-1, len(height)-1

	for left < right {
		num := min(height[left], height[right]) * width
		if res < num {
			res = num
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
		width--
	}

	return res
}
