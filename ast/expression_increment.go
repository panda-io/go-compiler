package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Increment struct {
	ExpressionBase
	Expression Expression
}

func (i *Increment) Type(c *Context) ir.Type {
	t := i.Expression.Type(c)
	if ir.IsPointer(t) {
		e := t.(*ir.PointerType).ElemType
		if ir.IsNumber(e) {
			return e
		}
	}
	c.Program.Error(i.Position, "invalid type for increment expression")
	return nil
}

func (i *Increment) GenerateIR(c *Context) ir.Value {
	t := i.Expression.Type(c)
	if ir.IsPointer(t) {
		p := i.Expression.GenerateIR(c)
		e := t.(*ir.PointerType).ElemType
		load := ir.NewLoad(e, p)
		c.Block.AddInstruction(load)
		if ir.IsInt(e) {
			add := ir.NewAdd(load, ir.NewInt(e.(*ir.IntType), 1))
			c.Block.AddInstruction(add)
			c.Block.AddInstruction(ir.NewStore(add, p))
			return add
		} else if ir.IsFloat(e) {
			add := ir.NewFAdd(load, ir.NewFloat(e.(*ir.FloatType), 1))
			c.Block.AddInstruction(add)
			c.Block.AddInstruction(ir.NewStore(add, p))
			return add
		}
	}
	c.Program.Error(i.Position, "invalid type for increment")
	return nil
}

func (*Increment) IsConstant() bool {
	//TO-DO
	return false
}

func (*Increment) GenerateConstIR(c Context, exprect ir.Type) ir.Value {
	//TO-DO
	return nil
}
