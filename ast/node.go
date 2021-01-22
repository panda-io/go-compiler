package ast

type Node interface {
	GetPosition() int
}

type NodeBase struct {
	Position int
}

func (b *NodeBase) GetPosition() int {
	return b.Position
}
