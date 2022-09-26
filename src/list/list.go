package list

import "fmt"

// List is a doubly linked list
type List struct {
	Head *Node
	Tail *Node
}

func NewList() *List {
	return &List{
		Head: nil,
		Tail: nil,
	}
}

func (l *List) Print() {
	// Display the linked list

	list := l.Head

	for list != nil {
		fmt.Printf("%+v ->", list.Value)
		list = list.Next
	}

	fmt.Println()
}

func (l *List) Length() int {
	// Return the length of the list

	count := 0
	node := l.Head

	for node != nil {
		count++
		node = node.Next
	}

	return count
}

func (l *List) Add(node *Node) {
	// Add the node to the end of the list

	if l.Head == nil {
		// Zero element list
		l.Head = node
		l.Tail = node

		node.Next = nil
		node.Prev = nil
	} else {
		// Non-zero element list
		l.Tail.Next = node

		node.Next = nil
		node.Prev = l.Tail

		l.Tail = node
	}
}

func (l *List) Delete(node *Node) {
	// Delete the given node from the list

	if node == l.Head {
		// Node is the head
		if node.Next == nil {
			// Node is the only element
			l.Head = nil
			l.Tail = nil
		} else {
			// Node is the head but not the only element
			l.Head = node.Next
			node.Next.Prev = nil
		}
	} else {
		// Node is not the head

		if node.Next == nil {
			// Node is the tail
			l.Tail = node.Prev
			node.Prev.Next = nil
		} else {
			// Node is in the middle
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
		}
	}
}
