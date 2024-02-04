package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i := removeDuplicates(nums)
	fmt.Println(i)
}

func removeDuplicates(nums []int) int {
	tmp := nums[0]
	count := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == tmp {
			continue
		}

		tmp = nums[i]
		nums[count] = nums[i]
		count++
	}

	fmt.Println(nums)

	return count
}
