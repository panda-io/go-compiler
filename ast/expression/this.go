package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type This struct {
	Base
}

func (e *This) GenerateIR(c *node.Context) {
	//TO-DO
}
