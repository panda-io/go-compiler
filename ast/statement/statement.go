package statement

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Statement interface {
	node.Node
	GenerateIR(*node.Context)
}

type Base struct {
	node.Base
}

func (*Base) GenerateIR(*node.Context) {
}
