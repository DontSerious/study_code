package main

import "fmt"

func main() {
	fmt.Printf("twoSum([]int{2, 7, 11, 15}, 9): %v\n", twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		num := numbers[i] + numbers[j]
		if num == target {
			return []int{i + 1, j + 1}
		} else if num > target {
			j--
		} else if num < target {
			i++
		}
	}

	return []int{-1, -1}
}
