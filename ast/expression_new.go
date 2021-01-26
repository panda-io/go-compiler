package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type New struct {
	ExpressionBase
	Typ       Type
	Arguments *Arguments
}

func (n *New) Type(c *Context) ir.Type {
	//TO-DO
	return nil
}

func (n *New) GenerateIR(c *Context) ir.Value {
	//TO-DO
	return nil
}

func (*New) IsConstant(p *Program) bool {
	return false
}

func (*New) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
