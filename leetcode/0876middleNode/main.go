package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle := head
	flag := false

	for head != nil {
		if flag == true {
			middle = middle.Next
		}
		head = head.Next
		flag = !flag
	}

	return middle
}
