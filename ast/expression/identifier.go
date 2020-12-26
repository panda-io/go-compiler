package expression

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Identifier struct {
	Base
	Name string
}

func (i *Identifier) GenerateIR(c *node.Context) ir.Value {
	v := c.FindVariable(i.Name)
	if v == nil {
		c.Error(i.Position, fmt.Sprintf("undefined %s", i.Name))
		return nil
	}
	switch t := v.(type) {
	case *ir.Global:
		load := ir.NewLoad(t.Type(), t)
		c.Block.AddInstruction(load)
		return load

	case *ir.Param, *ir.InstAlloca:
		return t
	}
	c.Error(i.Position, fmt.Sprintf("invalid variable: %s", i.Name))
	return nil
}
