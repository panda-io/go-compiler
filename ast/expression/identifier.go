package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Identifier struct {
	Base
	Name string
}

func (e *Identifier) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
