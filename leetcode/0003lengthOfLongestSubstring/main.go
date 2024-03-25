package main

import "fmt"

func main() {
	fmt.Printf("lengthOfLongestSubstring(\"pwwkew\"): %v\n", lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	m, res, bor := make(map[rune]int), 0, -1

	for i, v := range s {
		if _, ok := m[v]; ok {
			bor = max(bor, m[v])
		}
		m[v] = i
		res = max(res, i-bor)
	}

	return res
}
