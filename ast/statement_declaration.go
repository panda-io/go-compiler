package ast

type Declaration struct {
	StatementBase
	Name  *Identifier
	Type  Type
	Value Expression
}

func (*Declaration) GenerateIR(*Context) {
	//TO-DO
}
