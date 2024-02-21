package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	ans := make([]int, n)
	ans[0] = 1
	tmp := 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	for i := n - 2; i > -1; i-- {
		tmp *= nums[i+1]
		ans[i] *= tmp
	}

	return ans
}

// -1 1 0 -3 3
// 1 -1 0 -3 3
// 1 0 -1 -3 3
// 1 0 -3 -1 3
// 1 0 -3 3 -1
