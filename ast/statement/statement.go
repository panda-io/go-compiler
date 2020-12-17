package statement

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Statement interface {
	node.Node
	GenerateIR() ir.Value
}

type Base struct {
	node.Base
}
