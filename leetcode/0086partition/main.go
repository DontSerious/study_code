package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func partition(head *ListNode, x int) *ListNode {
	var headNode, min, bigHead, bigTail *ListNode

	for head != nil {
		if head.Val < x {
			if headNode == nil {
				headNode = head
			}
			if min == nil {
				min = head
			} else {
				min.Next = head
				min = head
			}
		} else {
			if bigHead == nil {
				bigHead = head
				bigTail = head
			} else {
				bigTail.Next = head
				bigTail = head
			}
		}

		head = head.Next
	}

	if headNode == nil {
		return bigHead
	}

	min.Next = bigHead
	if bigTail != nil {
		bigTail.Next = nil
	}
	return headNode
}
