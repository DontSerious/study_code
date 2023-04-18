package main

/*
	在排序数组中查找数字
	统计一个数字在排序数组中出现的次数。
*/

func search(nums []int, target int) int {
	res := 0
	
	for i := range nums {
		if nums[i] > target {
			break
		}
		if nums[i] == target {
			res++
		}
	}
	return res
}