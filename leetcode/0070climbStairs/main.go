package main

func main() {
	println(climbStairs(10))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := []int{1, 2}
	for i := 3; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}

	return dp[1]
}
