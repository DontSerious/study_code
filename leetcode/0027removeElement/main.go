package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	i := removeElement(nums, 2)
	fmt.Println(i)
}

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	fmt.Println(nums)
	return len(nums)
}

// func removeElement(nums []int, val int) int {
// 	tail := len(nums) - 1

// 	for i := 0; i <= tail; i++ {
// 		if nums[i] == val {
// 			swap(nums, i, tail)
// 			tail--
// 			i--
// 		}
// 	}

// 	fmt.Println(nums)

// 	return tail + 1
// }

// func swap(nums []int, i int, j int) {
// 	tmp := nums[i]
// 	nums[i] = nums[j]
// 	nums[j] = tmp
// }
