package LinkedList

import "testing"

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList(1)
	if ll.Head.Value != 1 {
		t.Errorf("Head set to %d, expected 1", ll.Head.Value)
	}
}

func TestAppend(t *testing.T) {
	ll := NewLinkedList(1)
	ll.append(4)

	current := ll.Head
	if current.Next.Value != 4 {
		t.Errorf("Next value is %d, when 4 is expected", current.Next.Value)
	}
}

func TestRemove(t *testing.T) {
	ll := NewLinkedList(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.remove(3)

	current := ll.Head
	for current.Next != nil {
		if current.Value == 3 {
			t.Error("3 not removed as expected")
		}
		current = current.Next
	}

}
