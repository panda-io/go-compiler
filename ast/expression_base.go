package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Base struct {
	ExpressionBase
}

func (*Base) Type(c *Context, expected ir.Type) ir.Type {
	return nil
}

func (*Base) GenerateIR(c *Context, expected ir.Type) ir.Value {
	return nil
}

func (*Base) IsConstant(p *Program) bool {
	return false
}

func (*Base) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
