package main

import "fmt"

func main() {
	fmt.Printf("isPowerOfTwo(16): %v\n", isPowerOfTwo(16))
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
