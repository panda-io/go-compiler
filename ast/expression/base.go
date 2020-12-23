package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Super struct {
	Base
}

func (e *Super) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
