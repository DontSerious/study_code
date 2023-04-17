package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

type ListNode struct {
	Val int
	Next *ListNode
}

// 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	var res []int
	toTheEnd(head, &res)
	return res
}

func toTheEnd(node *ListNode, res *[]int) {
	if node == nil {
		return
	}

	toTheEnd(node.Next, res)
	*res = append(*res, node.Val)
}