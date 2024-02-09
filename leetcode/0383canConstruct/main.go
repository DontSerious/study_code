package main

import "fmt"

func main() {
	fmt.Printf("canConstruct(\"aa\", \"aab\"): %v\n", canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	wordCount := make([]int, 26)
	for _, x := range magazine {
		wordCount[x-'a']++
	}
	for _, x := range ransomNote {
		i := int(x - 'a')
		if wordCount[i] <= 0 {
			return false
		}
		wordCount[i]--
	}

	return true
}

// func canConstruct(ransomNote string, magazine string) bool {
// 	if len(ransomNote) == 0 {
// 		return true
// 	}
// 	if len(magazine) == 0 {
// 		return false
// 	}

// 	mag := make(map[rune]int)
// 	for _, x := range magazine {
// 		mag[x]++
// 	}

// 	for _, x := range ransomNote {
// 		mag[x]--
// 		if mag[x] < 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
