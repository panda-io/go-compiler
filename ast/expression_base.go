package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Base struct {
	ExpressionBase
}

func (*Base) Type(c *Context) ir.Type {
	return nil
}

func (*Base) GenerateIR(c *Context) ir.Value {
	return nil
}

func (*Base) IsConstant() bool {
	return false
}

func (*Base) GenerateConstIR(c Context, exprect ir.Type) ir.Value {
	return nil
}
