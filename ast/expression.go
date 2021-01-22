package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Expression interface {
	Node
	Type(*Context) ir.Type
	//IsConstant() bool //TO-DO
	GenerateIR(*Context) ir.Value
}

type ExpressionBase struct {
	NodeBase
}
