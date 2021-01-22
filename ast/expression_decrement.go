package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Decrement struct {
	ExpressionBase
	Expression Expression
}

func (d *Decrement) Type(c *Context) ir.Type {
	if ir.IsNumber(d.Expression.Type(c)) {
		return d.Expression.Type(c)
	}
	c.Error(d.Position, "invalid type for decrement expression")
	return nil
}

func (d *Decrement) GenerateIR(c *Context) ir.Value {
	typ := i.Expression.Type(c)
	if ir.IsInt(typ) {
		add := ir.NewSub(i.Expression.GenerateIR(c), ir.NewInt(typ.(*ir.IntType), 1))
		c.Block.AddInstruction(add)
		return add
	} else if ir.IsFloat(typ) {
		add := ir.NewFSub(i.Expression.GenerateIR(c), ir.NewFloat(typ.(*ir.FloatType), 1))
		c.Block.AddInstruction(add)
		return add
	}
	c.Error(i.Position, "invalid type for decrement")
	return nil
}
