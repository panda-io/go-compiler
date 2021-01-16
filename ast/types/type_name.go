package types

type TypeName struct {
	Base
	Name          string
	Selector      string
	TypeArguments *TypeArguments
}
