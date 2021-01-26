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
		// TO-DO check pointer element type
		if c.First.Type(ctx).Equal(c.Second.Type(ctx)) {
			return c.First.Type(ctx)
		} else if ir.IsNumber(c.First.Type(ctx)) && ir.IsNumber(c.Second.Type(ctx)) {
			t, err := PromoteNumberType(ctx, c.First.Type(ctx), c.Second.Type(ctx))
			if err != nil {
				ctx.Program.Error(c.Position, err.Error())
			}
			return t
		}
	}
	ctx.Program.Error(c.Position, "invalid type of condition expression")
	return nil
}

func (c *Conditional) GenerateIR(ctx *Context) ir.Value {
	//TO-DO
	return nil
}
