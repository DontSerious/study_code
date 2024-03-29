package main

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		profit += max(0, prices[i]-prices[i-1])
	}

	return profit
}
