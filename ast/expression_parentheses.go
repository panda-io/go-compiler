package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Parentheses struct {
	ExpressionBase
	Expression Expression
}

func (p *Parentheses) Type(c *Context, expected ir.Type) ir.Type {
	return p.Expression.Type(c, expected)
}

func (p *Parentheses) GenerateIR(c *Context, expected ir.Type) ir.Value {
	return p.Expression.GenerateIR(c, expected)
}

func (p *Parentheses) IsConstant(program *Program) bool {
	return p.Expression.IsConstant(program)
}

func (p *Parentheses) GenerateConstIR(program *Program, expected ir.Type) ir.Constant {
	return p.Expression.GenerateConstIR(program, expected)
}
