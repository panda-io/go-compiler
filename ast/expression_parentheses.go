package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Parentheses struct {
	ExpressionBase
	Expression Expression
}

func (p *Parentheses) Type(c *Context) ir.Type {
	return p.Expression.Type(c)
}

func (p *Parentheses) GenerateIR(c *Context) ir.Value {
	//TO-DO
	return nil
}
