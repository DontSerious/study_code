package main

import "fmt"

func main() {
	fmt.Printf("integerBreak(10): %v\n", integerBreak(10))
}

func integerBreak(n int) int {
	dp := make(map[int]int)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}

	return dp[n]
}
