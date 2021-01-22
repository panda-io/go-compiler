package ast

type Throw struct {
	StatementBase
	Expression Expression
}

func (*Throw) GenerateIR(*Context) {
	//TO-DO
}
