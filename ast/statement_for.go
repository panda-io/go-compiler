package ast

type For struct {
	StatementBase
	Initialization Statement
	Condition      Statement
	Post           Statement
	Body           Statement
}

func (*For) GenerateIR(*Context) bool {
	//TO-DO
	return false // TO-DO check if child terminated
}
