package list

import "github.com/PlebusSupremus1234/FluidSim/src/particle"

type Node struct {
	Value *particle.Particle // Particle the node holds

	Prev *Node // Previous node
	Next *Node // Next node
}

func NewNode(p *particle.Particle) *Node {
	return &Node{
		Value: p,
	}
}
