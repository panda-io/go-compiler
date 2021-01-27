package ast

import "github.com/panda-foundation/go-compiler/ir"

type Break struct {
	StatementBase
}

func (b *Break) GenerateIR(c *Context) bool {
	if c.LeaveBlock == nil {
		c.Program.Error(b.Position, "invalid break")
	}
	c.Block.Term = ir.NewBr(c.LeaveBlock)
	return false
}
