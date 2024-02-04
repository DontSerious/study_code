package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	i := removeDuplicates(nums)
	fmt.Println(i)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	tmp := nums[0]
	flag := true

	for i := 1; i < n; i++ {
		if nums[i] == tmp && flag {
			flag = false
			continue
		} else if nums[i] != tmp {
			tmp = nums[i]
			flag = true
			continue
		}
		nums = append(nums[:i], nums[i+1:]...)
		i--
		n--
	}

	fmt.Println(nums)

	return n
}
