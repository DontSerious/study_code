package main

/*
	连续子数组的最大和
	输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

	要求时间复杂度为O(n)。

	输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出: 6
	解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

// 动态规划
// func maxSubArray(nums []int) int {
// 	max := nums[0]
	
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] + nums[i - 1] > nums[1] {
// 			nums[i] += nums[i - 1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}

// 	return max
// }

// 分治
func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum
}

/* 
	lSum 表示 [l,r] 内以 l 为左端点的最大子段和
		它要么等于「左子区间」的 lSum，要么等于「左子区间」的 iSum 加上「右子区间」的 lSum，二者取大。
	rSum 表示 [l,r] 内以 r 为右端点的最大子段和
		它要么等于「右子区间」的 rSum，要么等于「右子区间」的 iSum 加上「左子区间」的 rSum，二者取大。
	mSum 表示 [l,r] 内的最大子段和
	iSum 表示 [l,r] 的区间和
*/
func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)

	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}

	// m = (l + r) / 2
	m := (l + r) >> 1
	// 进入递归 拆分成最小粒子之后进行pushUp
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)

	return pushUp(lSub, rSub)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

func main() {
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	println(maxSubArray(arr))
}