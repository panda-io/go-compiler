package ast

type For struct {
	StatementBase
	Initialization Statement
	Condition      Statement
	Post           Statement
	Body           Statement
}

func (*For) GenerateIR(*Context) {
	//TO-DO
}
