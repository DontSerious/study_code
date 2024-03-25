package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("reverseWords(\"  hel   lo    w o   rld  \"): %v\n", reverseWords("  hel   lo    w o   rld  "))
}

func reverseWords(s string) string {
	if len(s) == 1 {
		return s
	}
	res := []string{}
	// go from end to front
	for i, j := len(s)-1, len(s)-1; i >= 0; {
		// skip words
		for i >= 0 && s[i] != ' ' {
			i--
		}
		// append the final word to the res
		res = append(res, s[i+1:j+1])
		// skip spaces
		for i >= 0 && s[i] == ' ' {
			i--
		}
		// pointer point to new word's end
		j = i
	}

	if res[0] == "" {
		res = res[1:]
	}

	return strings.Join(res, " ")
}
