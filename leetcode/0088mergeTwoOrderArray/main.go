package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; n > 0; i-- {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			m--
			nums1[i] = nums1[m]
		} else {
			n--
			nums1[i] = nums2[n]
		}
	}
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	sorted := make([]int, 0, m+n)
// 	p1, p2 := 0, 0

// 	for {
// 		if p1 == m {
// 			sorted = append(sorted, nums2[p2:]...)
// 			break
// 		}

// 		if p2 == n {
// 			sorted = append(sorted, nums1[p2:]...)
// 			break
// 		}

// 		if nums1[p1] < nums2[p2] {
// 			sorted = append(sorted, nums1[p1])
// 			p1++
// 		} else {
// 			sorted = append(sorted, nums2[p2])
// 			p2++
// 		}
// 	}

// 	copy(nums1, sorted)
// }
