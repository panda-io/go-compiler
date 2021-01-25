package ast

type Foreach struct {
	StatementBase
	Key      Statement
	Item     Statement
	Iterator Expression
	Body     Statement
}

func (*Foreach) GenerateIR(*Context) bool {
	//TO-DO
	return false // TO-DO check if child terminated
}
