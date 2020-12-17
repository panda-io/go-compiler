package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Parentheses struct {
	Base
	Expression Expression
}

func (e *Parentheses) GenerateIR() ir.Value {
	return e.Expression.GenerateIR()
}

func (e *Parentheses) IsConstant() bool {
	return e.Expression.IsConstant()
}
