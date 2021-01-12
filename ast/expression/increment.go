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
	//TO-DO
	return nil
}
