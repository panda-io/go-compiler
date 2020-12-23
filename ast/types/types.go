package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Type interface {
	node.Node
}

type Base struct {
	node.Base
}

func TypeOf(t Type) ir.Type {
	switch typ := t.(type) {
	case *BuitinType:
		return typ.GenerateIR()

	case *TypeName:
		return typ.GenerateIR()

	default:
		panic("inviad type define")
	}
}
