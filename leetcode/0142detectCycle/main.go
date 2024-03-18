package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	fast = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

// func detectCycle(head *ListNode) *ListNode {
// 	m := make(map[*ListNode]int)

// 	for head != nil {
// 		if m[head] != 0 {
// 			return head
// 		}
// 		m[head]++
// 		head = head.Next
// 	}

// 	return nil
// }
