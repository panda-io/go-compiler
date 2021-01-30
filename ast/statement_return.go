package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Return struct {
	StatementBase
	Expression Expression
}

func (r *Return) GenerateIR(c *Context) {
	var value ir.Value
	if r.Expression.IsConstant(c.Program) {
		value = r.Expression.GenerateConstIR(c.Program, c.Function.ReturnType.Type(c.Program))
	} else {
		value = r.Expression.GenerateIR(c, nil)
	}
	var t ir.Type = ir.Void
	if c.Function.ReturnType != nil {
		t = c.Function.ReturnType.Type(c.Program)
	}
	if value.Type().Equal(t) {
		c.Block.AddInstruction(ir.NewStore(value, c.Function.IRReturn))
	} else if p, ok := value.Type().(*ir.PointerType); ok && p.ElemType.Equal(t) {
		load := ir.NewLoad(p.ElemType, value)
		c.Block.AddInstruction(load)
		c.Block.AddInstruction(ir.NewStore(load, c.Function.IRReturn))
	} else {
		c.Program.Error(r.Position, "return type mismatch with function define")
	}
	c.Returned = true
	c.Block.AddInstruction(ir.NewBr(c.Function.IRExit))
}
