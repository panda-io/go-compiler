package ast

type Statement interface {
	Node
	// return type is if ternated in block
	GenerateIR(*Context) bool
}

type StatementBase struct {
	NodeBase
}
