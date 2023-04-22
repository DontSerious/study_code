package main

/* 
	青蛙跳台阶问题
	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

// 递归剪枝
var hash map[int]int = make(map[int]int)
const mod int = 1e9 + 7

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}

	if hash[n] != 0 {
		return hash[n]
	} else {
		num := (numWays(n - 1) + numWays(n - 2)) % mod
		hash[n] = num
		return num
	}
}

func main() {
	println(numWays(2))
}