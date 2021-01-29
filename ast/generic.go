package ast

type TypeArguments struct {
	NodeBase
	Arguments []Type
}

type TypeParameters struct {
	NodeBase
	Parameters []*TypeParameter
}

type TypeParameter struct {
	NodeBase
	Name string
	Type Type
}
