package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Return struct {
	StatementBase
	Expression Expression
}

func (r *Return) GenerateIR(c *Context) {
	value := r.Expression.GenerateIR(c)
	var t ir.Type = ir.Void
	if c.Function.ReturnType != nil {
		t = c.Function.ReturnType.Type(c)
	}
	if value.Type().Equal(t) {
		c.Block.AddInstruction(ir.NewStore(value, c.Function.Return))
	} else if p, ok := value.Type().(*ir.PointerType); ok && p.ElemType.Equal(t) {
		load := ir.NewLoad(p.ElemType, value)
		c.Block.AddInstruction(load)
		c.Block.AddInstruction(ir.NewStore(load, c.Function.Return))
	} else {
		c.Error(r.Position, "return type mismatch with function define")
	}
	c.Block.Term = ir.NewBr(c.Function.Exit)
}
