package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	i := majorityElement(nums)
	fmt.Println(i)
}

// Boyer-Moore 算法 candidate
func majorityElement(nums []int) int {
	var cd int
	count := 0

	for _, x := range nums {
		if count == 0 {
			cd = x
		}

		if cd == x {
			count++
		} else {
			count--
		}
	}

	return cd
}
