package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Parentheses struct {
	Base
	Expression Expression
}

func (p *Parentheses) Type(c *node.Context) ir.Type {
	return p.Expression.Type(c)
}

func (p *Parentheses) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
