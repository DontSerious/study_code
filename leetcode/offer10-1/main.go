package main

/*
	斐波那契数列
	写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

	F(0) = 0,   F(1) = 1
	F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

	斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

// 滚动数组
// func fib(n int) int {
// 	const mod int = 1e9 + 7

// 	if n < 2 {
// 		return n
// 	}

// 	// n = 3时的滚动数组
// 	p, q, r := 0, 0, 1
// 	for i := 2; i <= n; i++ {
// 		p = q
// 		q = r
// 		r = (q + p) % mod
// 	}

// 	return r
// }

// 递归剪枝
var hash map[int]int = make(map[int]int)
const mod int = 1e9 + 7

func fib(n int) int {

	if (n < 2) {
		return n
	}

	if hash[n] != 0 {
		return hash[n]
	} else {
		fib_n := (fib(n - 1) + fib(n - 2)) % mod
		hash[n] = fib_n
		return fib_n
	}
}

func main() {
	println(fib(4))
}