package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("", ""))
}

func isSubsequence(s string, t string) bool {
	si := 0
	l1, l2 := len(s), len(t)

	for i := 0; i < l2; i++ {
		if si == l1 {
			break
		}

		if t[i] == s[si] {
			si++
		}
	}

	return si == l1
}

// func isSubsequence(s string, t string) bool {
// 	if s == t || len(s) == 0 {
// 		return true
// 	}
// 	l1, l2 := len(s), len(t)

// 	if l1 > l2 {
// 		return false
// 	}

// 	for p1, p2 := 0, 0; p2 < l2; p2++ {
// 		if s[p1] == t[p2] {
// 			p1++
// 			if p1 == l1 {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }
