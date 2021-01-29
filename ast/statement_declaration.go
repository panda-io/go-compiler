package ast

type DeclarationStatement struct {
	StatementBase
	Name  *Identifier
	Type  Type
	Value Expression
}

func (*DeclarationStatement) GenerateIR(*Context) {
	//TO-DO // zero initialize
	//TO-DO class type conversion with auto pointer
	//Builtin, TypeName
}
