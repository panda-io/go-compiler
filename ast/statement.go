package ast

type Statement interface {
	Node
	GenerateIR(*Context)
}

type StatementBase struct {
	NodeBase
}
