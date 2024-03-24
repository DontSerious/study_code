package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("largestNumber([]int{3, 30, 34, 5, 9}): %v\n", largestNumber([]int{3, 30, 34, 5, 9}))
}

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}

	quickSort(ss)

	res := strings.Join(ss, "")
	if res[0] == '0' {
		return "0"
	}

	return res
}

func match(a, b string) bool {
	return a+b >= b+a
}

func quickSort(s []string) {
	if len(s) <= 1 {
		return
	}

	left, right := 0, len(s)-1
	base := 0

	for left != right {
		for left < right && match(s[base], s[right]) {
			right--
		}
		for left < right && match(s[left], s[base]) {
			left++
		}
		s[left], s[right] = s[right], s[left]
	}

	s[left], s[base] = s[base], s[left]
	quickSort(s[:left])
	quickSort(s[left+1:])
}
