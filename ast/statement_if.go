package ast

type If struct {
	StatementBase
	Initialization Statement
	Condition      Statement
	Body           Statement
	Else           Statement
}

func (*If) GenerateIR(*Context) {
	//TO-DO
}
