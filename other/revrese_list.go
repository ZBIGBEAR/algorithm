package other

// 反转链表
type Node struct {
	V    int64
	Next *Node
}

func RevreseList(head *Node) *Node {
	if head == nil {
		return nil
	}

	p1 := head
	p2 := head.Next
	head.Next = nil
	if p2 == nil {
		return head
	}

	p3 := p2.Next

	for {
		p2.Next = p1
		if p3 == nil {
			return p2
		}

		p1 = p2
		p2 = p3
		p3 = p3.Next
	}
	return nil
}
