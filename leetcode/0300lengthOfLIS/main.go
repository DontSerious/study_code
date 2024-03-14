package main

import "fmt"

func main() {
	// ans = 7, 8, 9 -> 1, 8, 9 -> 1, 2, 9 -> 1, 2, 3, 4, 5
	fmt.Printf("lengthOfLIS([]int{7, 8, 9, 1, 2, 3, 3, 5}): %v\n", lengthOfLIS([]int{7, 8, 9, 1, 2, 3, 3, 5}))
	fmt.Printf("lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}): %v\n", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	ans := make([]int, len(nums))
	res := 0

	for idx := range nums {

		// binary search, until ans[i] smaller or equal to nums[i]
		i, j := 0, res
		for i < j {
			mid := (i + j) / 2
			if ans[mid] < nums[idx] {
				i = mid + 1
			} else {
				j = mid
			}
		}

		ans[i] = nums[idx]
		if res == i {
			res++
		}
	}

	return res
}
