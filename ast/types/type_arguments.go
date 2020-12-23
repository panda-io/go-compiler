package types

type TypeArguments struct {
	Base
	Arguments []Type
	Ellipsis  int
}
