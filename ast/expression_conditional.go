package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Conditional struct {
	ExpressionBase
	Condition Expression
	First     Expression
	Second    Expression
}

func (c *Conditional) Type(ctx *Context) ir.Type {
	if ir.IsBool(c.Condition.Type(ctx)) {
		if c.First.Type(ctx).Equal(c.Second.Type(ctx)) {
			return c.First.Type(ctx)
		} else if ir.IsNumber(c.First.Type(ctx)) && ir.IsNumber(c.Second.Type(ctx)) {
			return PromoteNumberType(ctx, c.First, c.Second)
		}
	}
	ctx.Program.Error(c.Position, "invalid type of condition expression")
	return nil
}

func (c *Conditional) GenerateIR(ctx *Context) ir.Value {
	//TO-DO
	return nil
}
