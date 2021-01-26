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

func (d *Decrement) IsConstant(p *Program) bool {
	return d.Expression.IsConstant(p)
}

func (d *Decrement) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	expr := d.Expression.GenerateConstIR(p, nil)
	if expr != nil {
		t := expr.Type()
		if ir.IsInt(t) {
			return ir.NewExprSub(expr, ir.NewInt(t.(*ir.IntType), 1))
		} else if ir.IsFloat(t) {
			return ir.NewExprFSub(expr, ir.NewFloat(t.(*ir.FloatType), 1))
		}
	}
	p.Error(d.Position, "invalid decrement")
	return nil
}
