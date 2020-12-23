package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type New struct {
	Base
	Type      types.Type
	Arguments *types.Arguments
}

func (e *New) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
