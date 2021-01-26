package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Expression interface {
	Node
	Type(*Context) ir.Type
	IsConstant() bool
	GenerateIR(*Context) ir.Value
	GenerateConstIR(c Context, exprect ir.Type) ir.Value
}

type ExpressionBase struct {
	NodeBase
}
