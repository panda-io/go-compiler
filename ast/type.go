package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Type interface {
	Node
	Type(*Context) ir.Type
}

type TypeBase struct {
	NodeBase
}
