package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Expression interface {
	node.Node
	Type(*node.Context) ir.Type
	//TypeName() string //TO-DO
	//IsConstant() bool
	GenerateIR(*node.Context) ir.Value
}

type Base struct {
	node.Base
}
