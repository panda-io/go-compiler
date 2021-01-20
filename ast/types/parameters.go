package types

//TO-DO move to declaration/function.go, since parameter itself is not a pure type
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
