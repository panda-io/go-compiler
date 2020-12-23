package types

type TypeParameters struct {
	Base
	Parameters []*TypeParameter
	Ellipsis   bool
}

type TypeParameter struct {
	Base
	Name string
	Type Type
}
