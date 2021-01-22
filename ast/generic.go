package ast

type TypeArguments struct {
	NodeBase
	Arguments []Type
	Ellipsis  int
}

type TypeParameters struct {
	NodeBase
	Parameters []*TypeParameter
	Ellipsis   bool
}

type TypeParameter struct {
	NodeBase
	Name string
	Type Type
}
