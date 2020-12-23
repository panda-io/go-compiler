package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Arguments struct {
	Base
	Arguments []node.Node
	Ellipsis  int
}
