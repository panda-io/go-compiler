package ast

type Continue struct {
	StatementBase
}

func (*Continue) GenerateIR(*Context) bool {
	//TO-DO
	return true
}
