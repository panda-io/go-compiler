package node

type Node interface {
	GetPosition() int
}

type Base struct {
	Position int
}

func (b *Base) GetPosition() int {
	return b.Position
}
