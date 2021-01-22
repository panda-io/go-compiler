package ast

import (
	"github.com/panda-foundation/go-compiler/token"
)

type Switch struct {
	StatementBase
	Initialization Statement
	Operand        Statement
	Body           []*Case
}

type Case struct {
	StatementBase
	Token token.Token
	Case  Expression
	Body  Statement
}

func (*Switch) GenerateIR(*Context) {
	//TO-DO
}
