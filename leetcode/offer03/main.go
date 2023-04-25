package main

/*
	数组中重复的数字
	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
	请找出数组中任意一个重复的数字。
*/

func findRepeatNumber(nums []int) int {
	/* 
	// 哈希表法
	hash := make(map[int]int)

	for _, num := range nums {
		if hash[num] == 1 {
			return num
		}
		hash[num] = 1
	} 
	*/

	/* 
		原地交换：
		先对数组进行一定程度的排序，
		将前面几个未重复的数字放到与数字值相等的下标上，
		如果数字值已经和下标值相等，则跳过，
		如果在后续遍历中，数字值指向了已排序好的下标则代表重复，直接返回
	*/
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	return -1
}