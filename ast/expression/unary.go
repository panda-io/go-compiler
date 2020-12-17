package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Unary struct {
	Base
	Operator   token.Token
	Expression Expression
}

func (e *Unary) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Unary) IsConstant() bool {
	return e.Expression.IsConstant()
}
