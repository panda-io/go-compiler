package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type This struct {
	Base
}

func (t *This) Type(c *node.Context) ir.Type {
	return nil
}

func (t *This) GenerateIR(c *node.Context) ir.Value {
	return nil
}
