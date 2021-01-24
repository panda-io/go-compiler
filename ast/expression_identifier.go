package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type Identifier struct {
	ExpressionBase
	Name string
}

func (i *Identifier) Type(c *Context) ir.Type {
	t := c.ObjectType(i.Name)
	if t == nil {
		c.Program.Error(i.Position, fmt.Sprintf("%s undefined", i.Name))
	}
	return t
}

func (i *Identifier) GenerateIR(c *Context) ir.Value {
	v := c.FindObject(i.Name)
	if v == nil {
		c.Program.Error(i.Position, fmt.Sprintf("undefined %s", i.Name))
		return nil
	}
	switch t := v.(type) {
	case *ir.Global:
		load := ir.NewLoad(t.Type(), t)
		c.Block.AddInstruction(load)
		return load

	case *ir.InstAlloca:
		load := ir.NewLoad(t.ElemType, t)
		c.Block.AddInstruction(load)
		return load

	case *ir.InstLoad:
		return t

	case *ir.InstGetElementPtr:
		return t
	}
	c.Program.Error(i.Position, fmt.Sprintf("invalid variable: %s", i.Name))
	return nil
}
