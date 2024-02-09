package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	start, end := 0, 0
	sum, ans := 0, math.MaxInt32

	for end < n {
		sum += nums[end]

		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}

		end++
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
