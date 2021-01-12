package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Decrement struct {
	Base
	Expression Expression
}

func (d *Decrement) Type(c *node.Context) ir.Type {
	if ir.IsNumber(d.Expression.Type(c)) {
		return d.Expression.Type(c)
	}
	c.Error(d.Position, "invalid type for decrement expression")
	return nil
}

func (d *Decrement) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
