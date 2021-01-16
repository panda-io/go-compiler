package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Type interface {
	node.Node
}

type Base struct {
	node.Base
}
