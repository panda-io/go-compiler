package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Increment struct {
	Base
	Expression Expression
}

func (e *Increment) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
