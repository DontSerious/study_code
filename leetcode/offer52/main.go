package main

/* 
	两个链表的第一个公共节点
	输入两个链表，找出它们的第一个公共节点。
*/

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	node1, node2 := headA, headB

	for {
		if node1 == node2 {
			return node1
		}

		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}

		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
}