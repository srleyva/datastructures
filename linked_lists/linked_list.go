package LinkedList

type Node struct {
	Next  *Node
	Value int
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(data int) *LinkedList {
	head := &Node{nil, data}
	ll := &LinkedList{head}
	return ll
}

func (l *LinkedList) append(data int) {
	if l.Head == nil {
		l.Head = &Node{nil, data}
		return
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{nil, data}
}

func (l *LinkedList) remove(data int) {
	if l.Head.Value == data && l.Head.Next != nil {
		l.Head = l.Head.Next
		return
	}

	current := l.Head

	for current.Next != nil {
		if current.Next.Value == data {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
}
