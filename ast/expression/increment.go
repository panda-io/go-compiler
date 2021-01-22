package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Increment struct {
	Base
	Expression Expression
}

func (i *Increment) Type(c *node.Context) ir.Type {
	if ir.IsNumber(i.Expression.Type(c)) {
		return i.Expression.Type(c)
	}
	c.Error(i.Position, "invalid type for increment expression")
	return nil
}

func (i *Increment) GenerateIR(c *node.Context) ir.Value {
	typ := i.Expression.Type(c)
	if ir.IsInt(typ) {
		// cannot be bool
		if typ.(*ir.IntType).BitSize > 1 {
			add := ir.NewAdd(i.Expression.GenerateIR(c), ir.NewInt(typ.(*ir.IntType), 1))
			c.Block.AddInstruction(add)
			return add
		} else {
			c.Error(i.Position, "bool type is not valid for increment")
		}
	} else if ir.IsFloat(typ) {
		add := ir.NewFAdd(i.Expression.GenerateIR(c), ir.NewFloat(typ.(*ir.FloatType), 1))
		c.Block.AddInstruction(add)
		return add
	}
	return nil
}
