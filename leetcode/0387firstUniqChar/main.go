package main

import "fmt"

func main() {
	fmt.Printf("firstUniqChar(\"loveleetcode\"): %v\n", firstUniqChar("loveleetcode"))
	fmt.Printf("firstUniqChar(\"aadadaad\"): %v\n", firstUniqChar("aadadaad"))
	fmt.Printf("firstUniqChar(\"acaadcad\"): %v\n", firstUniqChar("acaadcad"))
}
func firstUniqChar(s string) int {
	m := make([]int, 26)

	for _, v := range s {
		m[v-'a']++
	}

	for i, v := range s {
		if m[v-'a'] == 1 {
			return i
		}
	}

	return -1
}
