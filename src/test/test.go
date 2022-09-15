package test

import (
	"github.com/PlebusSupremus1234/FluidSim/src/linked_list"
)

func Test() {
	a := &linked_list.Node{
		Value: &linked_list.TestValue{
			Value: 1,
		},
	}
	b := &linked_list.Node{
		Value: &linked_list.TestValue{
			Value: 2,
		},
	}
	c := &linked_list.Node{
		Value: &linked_list.TestValue{
			Value: 3,
		},
	}
	d := &linked_list.Node{
		Value: &linked_list.TestValue{
			Value: 4,
		},
	}

	l := linked_list.NewList(a)

	l.Print()

	l.Insert(b)
	l.Insert(c)
	l.Insert(d)

	l.Print()

	l.Delete(c)

	l.Print()

	l.Insert(c)
	l.Delete(b)

	l.Print()

	l.Delete(a)
	l.Delete(c)

	l.Print()

	l.Delete(d)

	l.Print()
}
