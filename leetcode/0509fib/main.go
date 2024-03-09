package main

import "fmt"

func main() {
	fmt.Printf("fib(4): %v\n", fib(4))
	fmt.Printf("fib(5): %v\n", fib(5))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	res := []int{0, 1}
	for i := 2; i <= n; i++ {
		tmp := res[0] + res[1]
		res[0] = res[1]
		res[1] = tmp
	}

	return res[1]
}
