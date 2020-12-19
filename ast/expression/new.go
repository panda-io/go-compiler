package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type New struct {
	Base
	Type      types.Type
	Arguments *types.Arguments
}

func (e *New) GenerateIR(c *node.Context) {
	//TO-DO
}
