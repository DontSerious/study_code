package main

func main() {
	println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-'a']++
	}

	for _, r := range t {
		record[r-'a']--
	}

	return record == [26]int{}
}

// func isAnagram(s string, t string) bool {
// 	m := make(map[byte]int)
// 	l1, l2 := len(s), len(t)
// 	if l1 != l2 {
// 		return false
// 	}

// 	for i := 0; i < l1; i++ {
// 		m[s[i]]++
// 	}

// 	for i := 0; i < l2; i++ {
// 		m[t[i]]--
// 		if m[t[i]] < 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
