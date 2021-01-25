package ast

type Try struct {
	StatementBase
	Try     Statement
	Operand *Parameters
	Catch   Statement
	Finally Statement
}

func (*Try) GenerateIR(*Context) bool {
	//TO-DO
	return false //Check children
}
