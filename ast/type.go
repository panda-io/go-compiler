package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Type interface {
	Node
	Type() ir.Type
}

type TypeBase struct {
	NodeBase
}
