package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Return struct {
	StatementBase
	Expression Expression
}

func (r *Return) GenerateIR(c *Context) {
	c.Block.Term = ir.NewRet(r.Expression.GenerateIR(c))
}
