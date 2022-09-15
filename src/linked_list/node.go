package linked_list

type TestValue struct {
	Value int
}

type Node struct {
	Value *TestValue
	//value *particle.Particle

	Prev *Node
	Next *Node
}
