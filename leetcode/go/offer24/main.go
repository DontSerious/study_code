package main

type ListNode struct {
    Val int
    Next *ListNode
}

/*
	递归实现反转链表
	定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
	示例:
		输入: 1->2->3->4->5->NULL
		输出: 5->4->3->2->1->NULL
*/
func reverseList(head *ListNode) *ListNode {
	return toTheEnd(head, nil)
}

func toTheEnd(cur *ListNode, pre *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	// 直接让最后一个节点当作反转链表的表头，res 存的就是这个链头的地址
	res := toTheEnd(cur.Next, cur)

	// 操作这一步相当于也操作了反转链表，因为它们的链头地址相同
	cur.Next = pre

	return res
}