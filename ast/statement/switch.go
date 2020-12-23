package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/token"
)

type Switch struct {
	Base
	Initialization Statement
	Operand        Statement
	Body           []*Case
}

type Case struct {
	Base
	Token token.Token
	Case  expression.Expression
	Body  Statement
}

func (*Switch) GenerateIR(*node.Context) {
	//TO-DO
}
