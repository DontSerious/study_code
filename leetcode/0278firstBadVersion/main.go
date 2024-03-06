package main

import "sort"

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}

// func firstBadVersion(n int) int {
// 	head, tail := 0, n

// 	for head < tail {
// 		mid := head + (tail-head)/2
// 		if isBadVersion(mid) {
// 			tail = mid
// 		} else {
// 			head = mid + 1
// 		}
// 	}

// 	return head
// }
