package main

/* 
	0～n-1中缺失的数字
	一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
	在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
*/

func missingNumber(nums []int) int {
	/* for i:=0;i<len(nums);i++{
        if nums[i] != i{
            return i
        }
    }
    return len(nums) */

	if nums[0] == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
		if i + 1 == len(nums) {
			return i + 1
		}
	}
	return -1
}