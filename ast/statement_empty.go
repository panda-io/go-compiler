package ast

type Empty struct {
	StatementBase
}

func (*Empty) GenerateIR(*Context) {
}
