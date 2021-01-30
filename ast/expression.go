package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Expression interface {
	Node
	Type(c *Context, expected ir.Type) ir.Type
	IsConstant(p *Program) bool
	GenerateIR(c *Context, expected ir.Type) ir.Value
	GenerateConstIR(p *Program, expected ir.Type) ir.Constant
}

type ExpressionBase struct {
	NodeBase
}
