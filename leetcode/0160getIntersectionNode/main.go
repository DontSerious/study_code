package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpA, tmpB := headA, headB
	for tmpA != tmpB {
		if tmpA != nil {
			tmpA = tmpA.Next
		} else {
			tmpA = headB
		}

		if tmpB != nil {
			tmpB = tmpB.Next
		} else {
			tmpB = headA
		}
	}

	return tmpA
}
