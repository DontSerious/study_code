package main

import "fmt"

func main() {
	fmt.Printf("myPow(2, 10): %v\n", myPow(2, 10))
	fmt.Printf("myPow(2, -2): %v\n", myPow(2, -2))
	fmt.Printf("myPow(2.1, 3): %v\n", myPow(2.1, 3))
	fmt.Printf("myPow(2, -2147483648): %v\n", myPow(2, -2147483648))
}

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}

	var b int64
	var res float64

	b = int64(n)
	res = 1
	if b < 0 {
		b = -b
		x = 1 / x
	}

	for b > 0 {
		if (b & 1) == 1 {
			res *= x
		}
		b >>= 1
		x *= x
	}

	return res
}
