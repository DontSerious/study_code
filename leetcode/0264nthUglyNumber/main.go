package main

import "fmt"

func main() {
	fmt.Printf("nthUglyNumber(10): %v\n", nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1
	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		j, q, k := res[a]*2, res[b]*3, res[c]*5
		res[i] = min(j, q, k)
		if res[i] == j {
			a++
		}
		if res[i] == q {
			b++
		}
		if res[i] == k {
			c++
		}
	}

	return res[n-1]
}
