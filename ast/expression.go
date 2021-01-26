package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Expression interface {
	Node
	Type(*Context) ir.Type
	IsConstant(p *Program) bool
	GenerateIR(*Context) ir.Value
	GenerateConstIR(p *Program, expected ir.Type) ir.Constant
}

type ExpressionBase struct {
	NodeBase
}
