package main

import "fmt"

/*
	旋转数组的最小数字
	把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。

	给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
	请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。

	注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
*/

func minArray(numbers []int) int {
	var mid int
	left, right := 0, len(numbers) - 1
	for left < right {
		mid = (left + right) / 2
		if numbers[mid] > numbers[right] {			// 这种情况下 mid 一定在左排序序列中
			left = mid + 1
		} else if numbers[mid] < numbers[right] {	// 这种情况下 mid 一定在右排序序列中
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}

func main() {
	arr := []int{2,2,2,0,1}
	fmt.Printf("minArray(arr): %v\n", minArray(arr))
}