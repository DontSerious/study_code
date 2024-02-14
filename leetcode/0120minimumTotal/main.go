package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	println(minimumTotal(triangle))
}

// 空间优化
func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	dp := make([]int, h+1)

	for i := h - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	fmt.Println(dp)

	return dp[0]
}

// func minimumTotal(triangle [][]int) int {
// 	h := len(triangle)
// 	dp := make([][]int, h)
// 	for i := range dp {
// 		dp[i] = make([]int, len(triangle[i]))
// 	}

// 	for i := h - 1; i >= 0; i-- {
// 		for j := 0; j < len(triangle[i]); j++ {
// 			if i == h-1 {
// 				dp[i][j] = triangle[i][j]
// 			} else {
// 				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
// 			}
// 		}
// 	}

// 	fmt.Println(dp)

// 	return dp[0][0]
// }

func min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}
