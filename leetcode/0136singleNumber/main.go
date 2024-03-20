package main

import "fmt"

func main() {
	fmt.Printf("singleNumber([]int{4, 1, 2, 1, 2}): %v\n", singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}

	return ans
}

// func singleNumber(nums []int) int {
// 	m := make(map[int]int)
// 	var res int
// 	for i := range nums {
// 		if m[nums[i]] == 1 {
// 			delete(m, nums[i])
// 		} else {
// 			m[nums[i]]++
// 		}
// 	}
// 	for k := range m {
// 		res = k
// 	}

// 	return res
// }
