package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Conditional struct {
	Base
	Condition Expression
	First     Expression
	Second    Expression
}

func (e *Conditional) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
