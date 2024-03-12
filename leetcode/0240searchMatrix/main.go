package main

import "fmt"

func main() {
	fmt.Printf("searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5): %v\n",
		searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Printf("searchMatrix([][]int{{5}, {6}}, 6): %v\n", searchMatrix([][]int{{5}, {6}}, 6))
}

func searchMatrix(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
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

// func searchMatrix(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return false
// 	}
// 	h, hm := 0, len(matrix)

// 	for i := 0; i < len(matrix[0]); i++ {
// 		num := matrix[h][i]
// 		if num == target {
// 			return true
// 		}
// 		if matrix[h][i] > target || i+1 == len(matrix[0]) {
// 			if h+1 < hm && matrix[h+1][0] <= target {
// 				i = -1
// 				h++
// 			} else {
// 				return false
// 			}
// 		}
// 	}

// 	return false
// }
