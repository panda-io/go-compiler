package ast

type Statement interface {
	Node
	GenerateIR(*Context) bool
}

type StatementBase struct {
	NodeBase
}
