package LinkedList

type Node struct {
	Next  *Node
	Prev  *Node
	Value int
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList(data int) *LinkedList {
	node := &Node{nil, nil, data}
	head := node
	tail := node
	ll := &LinkedList{head, tail}
	return ll
}

func (l *LinkedList) add(data int) {
	if l.Head == nil {
		l.Head = &Node{nil, nil, data}
		return
	}

	l.Tail.Next = &Node{nil, l.Tail, data}
	l.Tail = l.Tail.Next
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

func (l *LinkedList) count() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
