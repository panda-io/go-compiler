package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Unary struct {
	ExpressionBase
	Operator   token.Token
	Expression Expression
}

func (u *Unary) Type(c *Context, expected ir.Type) ir.Type {
	t := u.Expression.Type(c, expected)

	switch u.Operator {
	case token.Plus, token.Minus:
		if ir.IsNumber(t) {
			return t
		}

	case token.Not:
		if ir.IsBool(t) {
			return ir.I1
		}

	case token.Complement:
		if ir.IsInt(t) {
			return t
		}
	}
	return nil
}

func (u *Unary) GenerateIR(c *Context, expected ir.Type) ir.Value {
	t := u.Expression.Type(c, expected)
	v := u.Expression.GenerateIR(c, expected)
	var inst ir.Instruction

	switch u.Operator {
	case token.Plus:
		if ir.IsNumber(t) {
			return v
		}

	case token.Minus:
		if ir.IsFloat(t) {
			inst = ir.NewFSub(ir.NewFloat(t.(*ir.FloatType), 0), v)
		} else if ir.IsInt(t) {
			inst = ir.NewSub(ir.NewInt(t.(*ir.IntType), 0), v)
		}

	case token.Not:
		if ir.IsBool(t) {
			inst = ir.NewXor(v, ir.True)
		}

	case token.Complement:
		if ir.IsInt(t) {
			inst = ir.NewXor(v, ir.NewInt(t.(*ir.IntType), -1))
		}
	}

	if inst == nil {
		c.Program.Error(u.Position, "invalid type for unary expression")
		return nil
	}
	c.Block.AddInstruction(inst)
	return inst.(ir.Value)
}

func (u *Unary) IsConstant(p *Program) bool {
	return u.Expression.IsConstant(p)
}

func (u *Unary) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	v := u.Expression.GenerateConstIR(p, expected)
	t := v.Type()

	switch u.Operator {
	case token.Plus:
		if ir.IsNumber(t) {
			return v
		}

	case token.Minus:
		if ir.IsFloat(t) {
			return ir.NewExprFSub(ir.NewFloat(t.(*ir.FloatType), 0), v)
		} else if ir.IsInt(t) {
			return ir.NewExprSub(ir.NewInt(t.(*ir.IntType), 0), v)
		}

	case token.Not:
		if ir.IsBool(t) {
			return ir.NewExprXor(v, ir.True)
		}

	case token.Complement:
		if ir.IsInt(t) {
			return ir.NewExprXor(v, ir.NewInt(t.(*ir.IntType), -1))
		}
	}

	p.Error(u.Position, "invalid type for unary expression")
	return nil
}
