package ast

type Foreach struct {
	StatementBase
	Key      Statement
	Item     Statement
	Iterator Expression
	Body     Statement
}

func (*Foreach) GenerateIR(*Context) {
	//TO-DO
}
