package types

type Parameters struct {
	Base
	Parameters []*Parameter
	Ellipsis   bool
}

type Parameter struct {
	Base
	Name string
	Type Type
}
