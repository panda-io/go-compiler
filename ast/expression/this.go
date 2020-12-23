package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type This struct {
	Base
}

func (e *This) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
