package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	var ans []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0

	for {
		// read top
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		if top++; top > bottom {
			break
		}

		// read right
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if right--; right < left {
			break
		}

		// read bottom
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		if bottom--; bottom < top {
			break
		}

		// read left
		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}

	return ans
}
