package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Return struct {
	Base
	Expression expression.Expression
}

func (r *Return) GenerateIR(c *node.Context) {
	c.Block.Term = ir.NewRet(r.Expression.GenerateIR(c))
}
