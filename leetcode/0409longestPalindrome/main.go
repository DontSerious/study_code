package main

import "fmt"

func main() {
	fmt.Printf("longestPalindrome(\"abccccdd\"): %v\n", longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	res, odd := 0, 0

	for _, v := range s {
		m[byte(v)]++
	}

	for _, v := range m {
		rem := v % 2
		res += v - rem
		if rem == 1 {
			odd = 1
		}
	}

	return res + odd
}
