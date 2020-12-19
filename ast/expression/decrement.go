package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Decrement struct {
	Base
	Expression Expression
}

func (e *Decrement) GenerateIR(c *node.Context) {
	//TO-DO
}
