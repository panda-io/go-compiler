package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Decrement struct {
	Base
	Expression Expression
}

func (e *Decrement) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
