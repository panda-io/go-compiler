package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Parentheses struct {
	Base
	Expression Expression
}

func (e *Parentheses) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
