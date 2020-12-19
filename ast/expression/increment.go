package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Increment struct {
	Base
	Expression Expression
}

func (e *Increment) GenerateIR(c *node.Context) {
	//TO-DO
}
