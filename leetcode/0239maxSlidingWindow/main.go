package main

import "fmt"

func main() {
	// fmt.Printf("maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3): %v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Printf("maxSlidingWindow([]int{1, -1}, 1): %v\n", maxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Printf("maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5): %v\n", maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5))
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	q := []int{} // 单调队列，存放下标

	for i, v := range nums {
		// in
		for len(q) > 0 && nums[q[len(q)-1]] <= v {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// out
		if i-q[0] >= k {
			q = q[1:]
		}

		// record
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
