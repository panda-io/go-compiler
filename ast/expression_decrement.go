package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Decrement struct {
	ExpressionBase
	Expression Expression
}

func (d *Decrement) Type(c *Context, expected ir.Type) ir.Type {
	t := d.Expression.Type(c, expected)
	if ir.IsNumber(t) {
		return t
	}
	c.Program.Error(d.Position, "invalid type for decrement expression")
	return nil
}

func (d *Decrement) GenerateIR(c *Context, expected ir.Type) ir.Value {
	t := d.Expression.Type(c, expected)
	if ir.IsNumber(t) {
		e := d.Expression.GenerateIR(c, expected)
		operand := c.AutoLoad(e)
		if ir.IsInt(t) {
			sub := ir.NewSub(operand, ir.NewInt(t.(*ir.IntType), 1))
			c.Block.AddInstruction(sub)
			c.Block.AddInstruction(ir.NewStore(sub, e))
			return sub
		} else if ir.IsFloat(t) {
			sub := ir.NewFSub(operand, ir.NewFloat(t.(*ir.FloatType), 1))
			c.Block.AddInstruction(sub)
			c.Block.AddInstruction(ir.NewStore(sub, e))
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
