package main

/* 
	合并两个排序的链表
	输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
*/

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		tmp := l1
		l1 = l2
		l2 = tmp
	}

	pre := l1
	cur := l1.Next

	for {
		if cur == nil {
			pre.Next = l2
			return l1
		} else if l2 == nil {
			return l1
		}

		if cur.Val < l2.Val {
			pre = cur
			cur = cur.Next
			continue
		}

		tmp := l2.Next
		pre.Next = l2
		l2.Next = cur
		pre = pre.Next
		l2 = tmp
	}
}