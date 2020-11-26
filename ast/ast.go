package ast

// Node is base struct of all ast
type Node struct {
	position int
}

// Position return global position of this node
func (n *Node) Position() int {
	return n.position
}

// Metadata type
type Metadata struct {
	Node
	Name   string
	Text   string
	Values map[string]*Literal
}
