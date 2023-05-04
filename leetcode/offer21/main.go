package main

/* 
	调整数组顺序使奇数位于偶数前面
	输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
	使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
*/

func exchange(nums []int) []int {
	i, j := 0, len(nums) - 1

	for i < j {
		if nums[i] % 2 != 0 {
			i++
			continue
		}

		if nums[j] % 2 == 0 {
			j--
			continue
		}

		if nums[i] % 2 == 0 && nums[j] % 2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++; j--;
		}
	}
	
	return nums
}