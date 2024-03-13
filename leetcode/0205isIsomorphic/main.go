package main

import (
	"fmt"
)

func main() {
	fmt.Printf("isIsomorphic(\"egg\", \"add\"): %v\n", isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	m1, m2 := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if c, ok := m1[x]; ok {
			if c != y {
				return false
			}
		}
		if c, ok := m2[y]; ok {
			if c != x {
				return false
			}
		}

		m1[x] = y
		m2[y] = x
	}

	return true
}
