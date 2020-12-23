package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Expression struct {
	Base
	Expression expression.Expression
}

func (e *Expression) GenerateIR(c *node.Context) {
	e.Expression.GenerateIR(c)
}
