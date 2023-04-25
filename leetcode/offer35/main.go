package main

/*
	复杂链表的复制
	请实现 copyRandomList 函数，复制一个复杂链表。
	在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
	还有一个 random 指针指向链表中的任意节点或者 null。
*/

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
    if head == nil {
		return nil
	}

	hash := make(map[*Node]*Node)

	// 复制各节点，并建立 “原节点 -> 新节点” 的 Map 映射
	cur := head
	for cur != nil {
		var p Node
		p.Val = cur.Val
		hash[cur] = &p
		cur = cur.Next
	}

	// 构建新链表的 next 和 random 指向
	cur = head
	for cur != nil  {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur = cur.Next
	}

	return hash[head]
}

func main() {
	n0 := Node{Val: 3}
	n1 := Node{Val: 3}
	n2 := Node{Val: 3}

	n0.Next = &n1
	n1.Next = &n2
	n1.Random = &n0

	copyRandomList(&n0)
}