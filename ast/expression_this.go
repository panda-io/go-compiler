package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type This struct {
	ExpressionBase
}

func (t *This) Type(c *Context) ir.Type {
	return nil
}

func (t *This) GenerateIR(c *Context) ir.Value {
	return nil
}
