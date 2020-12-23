package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Unary struct {
	Base
	Operator   token.Token
	Expression Expression
}

func (e *Unary) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
