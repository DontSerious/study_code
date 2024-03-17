package main

import "fmt"

func main() {
	fmt.Printf("rotateString(\"abcde\", \"cdeab\"): %v\n", rotateString("abcde", "cdeab"))
	// fmt.Printf("rotateString(\"aa\", \"a\"): %v\n", rotateString("aa", "a"))
	// fmt.Printf("rotateString(\"clrwmpkwru\", \"wmpkwruclr\"): %v\n", rotateString("clrwmpkwru", "wmpkwruclr"))
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}

	for i := range s {
		if s[i] == goal[0] {
			tmp := s[i:] + s[:i]
			if tmp == goal {
				return true
			} else {
				continue
			}
		}
	}

	return false
}

// func rotateString(s string, goal string) bool {
// 	if len(s) != len(goal) {
// 		return false
// 	}

// 	var flag bool

// 	verify := func(index int) {
// 		for i := range goal {
// 			if index == len(s) {
// 				index = 0
// 			}
// 			if s[index] != goal[i] {
// 				flag = false
// 				return
// 			}
// 			index++
// 		}
// 		flag = true
// 	}

// 	for i := range s {
// 		if s[i] == goal[0] {
// 			verify(i)
// 			if flag {
// 				break
// 			} else {
// 				continue
// 			}
// 		}
// 	}

// 	return flag
// }
