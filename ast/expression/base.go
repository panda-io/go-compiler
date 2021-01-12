package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Super struct {
	Base
}

func (e *Super) Type(c *node.Context) ir.Type {
	return nil
}

func (e *Super) GenerateIR(c *node.Context) ir.Value {
	return nil
}
