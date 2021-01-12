package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Conditional struct {
	Base
	Condition Expression
	First     Expression
	Second    Expression
}

func (c *Conditional) Type(ctx *node.Context) ir.Type {
	if c.Condition.Type(ctx) == ir.I1 {
		if c.First.Type(ctx).Equal(c.Second.Type(ctx)) {
			return c.First.Type(ctx)
		} else if ir.IsNumber(c.First.Type(ctx)) && ir.IsNumber(c.Second.Type(ctx)) {
			return PromoteNumberType(ctx, c.First, c.Second)
		}
	}
	ctx.Error(c.Position, "invalid type of condition expression")
	return nil
}

func (c *Conditional) GenerateIR(ctx *node.Context) ir.Value {
	//TO-DO
	return nil
}
