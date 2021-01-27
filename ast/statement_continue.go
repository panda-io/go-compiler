package ast

import "github.com/panda-foundation/go-compiler/ir"

type Continue struct {
	StatementBase
}

func (con *Continue) GenerateIR(c *Context) bool {
	if c.LoopBlock == nil {
		c.Program.Error(con.Position, "invalid continue")
	}
	c.Block.Term = ir.NewBr(c.LoopBlock)
	return false
}
