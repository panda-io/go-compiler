package ast

type TypeArguments struct {
	NodeBase
	Arguments []Type
	Ellipsis  int
}
