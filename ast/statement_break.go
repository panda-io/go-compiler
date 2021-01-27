package ast

import "github.com/panda-foundation/go-compiler/ir"

type Break struct {
	StatementBase
}

func (b *Break) GenerateIR(c *Context) {
	if c.LeaveBlock == nil {
		c.Program.Error(b.Position, "invalid break")
	}
	c.Block.AddInstruction(ir.NewBr(c.LeaveBlock))
}
