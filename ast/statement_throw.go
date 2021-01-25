package ast

type Throw struct {
	StatementBase
	Expression Expression
}

func (*Throw) GenerateIR(*Context) bool {
	//TO-DO
	return true // ? TO-DO
}
