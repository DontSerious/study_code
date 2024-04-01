package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("threeSum([]int{-1, 0, 1, 2, -1, -4}): %v\n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("threeSum2([]int{-1, 0, 1, 2, -1, -4}): %v\n", threeSum2([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				res = append(res, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}

	return res
}
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i, j := k+1, len(nums)-1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if sum > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}

	return res
}
