package linked_list

import "fmt"

type List struct {
	Head *Node
	Tail *Node
}

func NewList(node *Node) *List {
	l := &List{
		Head: node,
		Tail: node,
	}

	l.Head.Next = nil

	return l
}

func (l *List) Print() {
	list := l.Head

	for list != nil {
		fmt.Printf("%+v ->", list.Value)
		list = list.Next
	}

	fmt.Println()
}

func (l *List) Insert(node *Node) {
	// Insert the node into the back of the list
	l.Tail.Next = node
	node.Next = nil
	node.Prev = l.Tail

	l.Tail = node
}

func (l *List) Delete(node *Node) {
	// Delete the given node from the list
	if node == l.Head {
		l.Head = node.Next

		if node.Next != nil {
			l.Head.Prev = nil
		}
	} else {
		node.Prev.Next = node.Next
	}
}
