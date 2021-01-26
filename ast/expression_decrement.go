package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Decrement struct {
	ExpressionBase
	Expression Expression
}

func (d *Decrement) Type(c *Context) ir.Type {
	t := d.Expression.Type(c)
	if ir.IsPointer(t) {
		e := t.(*ir.PointerType).ElemType
		if ir.IsNumber(e) {
			return e
		}
	}
	c.Program.Error(d.Position, "invalid type for decrement expression")
	return nil
}

func (d *Decrement) GenerateIR(c *Context) ir.Value {
	t := d.Expression.Type(c)
	if ir.IsPointer(t) {
		p := d.Expression.GenerateIR(c)
		e := t.(*ir.PointerType).ElemType
		load := ir.NewLoad(e, p)
		c.Block.AddInstruction(load)
		if ir.IsInt(e) {
			sub := ir.NewSub(load, ir.NewInt(e.(*ir.IntType), 1))
			c.Block.AddInstruction(sub)
			c.Block.AddInstruction(ir.NewStore(sub, p))
			return sub
		} else if ir.IsFloat(e) {
			sub := ir.NewFSub(load, ir.NewFloat(e.(*ir.FloatType), 1))
			c.Block.AddInstruction(sub)
			c.Block.AddInstruction(ir.NewStore(sub, p))
			return sub
		}
	}
	c.Program.Error(d.Position, "invalid type for decrement")
	return nil
}

func (*Decrement) IsConstant() bool {
	//TO-DO
	return false
}

func (*Decrement) GenerateConstIR(p *Program, exprect ir.Type) ir.Value {
	//TO-DO
	return nil
}
