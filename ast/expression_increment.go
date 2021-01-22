package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Increment struct {
	ExpressionBase
	Expression Expression
}

func (i *Increment) Type(c *Context) ir.Type {
	if ir.IsNumber(i.Expression.Type(c)) {
		return i.Expression.Type(c)
	}
	c.Error(i.Position, "invalid type for increment expression")
	return nil
}

func (i *Increment) GenerateIR(c *Context) ir.Value {
	typ := i.Expression.Type(c)
	if ir.IsInt(typ) {
		add := ir.NewAdd(i.Expression.GenerateIR(c), ir.NewInt(typ.(*ir.IntType), 1))
		c.Block.AddInstruction(add)
		return add
	} else if ir.IsFloat(typ) {
		add := ir.NewFAdd(i.Expression.GenerateIR(c), ir.NewFloat(typ.(*ir.FloatType), 1))
		c.Block.AddInstruction(add)
		return add
	}
	c.Error(i.Position, "invalid type for increment")
	return nil
}
