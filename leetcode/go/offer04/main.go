package main

/* 
	二维数组中的查找
	在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。
	请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

// 暴力解法
// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 {
// 		return false
// 	} else if len(matrix[0]) == 0 {
// 		return false
// 	}

// 	a := len(matrix) - 1

// 	for i := 0; i <= a; i++ {
// 		if matrix[i][0] == target {
// 			return true
// 		}
// 		if matrix[i][0] > target {
// 			a = i - 1
// 			break
// 		}
// 	}

// 	for ; a >= 0; a-- {
// 		for _, v := range matrix[a] {
// 			if v == target {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func findNumberIn2DArray(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0

	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

func main() {
	arr := [][]int{{1,2,3,4,5}, {6,7,8,9,10}, {11,12,13,14,15}, {16,17,18,19,20}, {21,22,23,24,25}}
	println(findNumberIn2DArray(arr, 19))
}