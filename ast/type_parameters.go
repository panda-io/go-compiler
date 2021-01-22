package ast

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
