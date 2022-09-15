package linked_list

import "github.com/PlebusSupremus1234/FluidSim/src/particle"

type Node struct {
	Value *particle.Particle

	Prev *Node
	Next *Node
}

func NewNode(p *particle.Particle) *Node {
	return &Node{
		Value: p,
	}
}
