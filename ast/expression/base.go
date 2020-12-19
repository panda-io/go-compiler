package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Super struct {
	Base
}

func (e *Super) GenerateIR(c *node.Context) {
	//TO-DO
}
