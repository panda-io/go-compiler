package ast

type Statement interface {
	Node
	// return type is if ternated in block
	GenerateIR(*Context)
}

type StatementBase struct {
	NodeBase
}
