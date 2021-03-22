package ast

type TypeName struct {
	TypeBase
	Name          string
	Selector      string
	TypeArguments *TypeArguments
}
