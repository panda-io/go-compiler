package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Switch struct {
	StatementBase
	Initialization Statement
	Operand        Expression
	Cases          []*Case
	Default        *Case
}

type Case struct {
	StatementBase
	Token token.Token
	Case  Expression
	Body  Statement
}

func (s *Switch) GenerateIR(c *Context) bool {
	ctx := c.NewContext()
	ctx.Block = c.Block
	ctx.Terminated = true
	if s.Initialization != nil {
		s.Initialization.GenerateIR(ctx)
	}

	leaveBlock := c.Function.IRFunction.NewBlock("")
	defaultBlock := c.Function.IRFunction.NewBlock("")
	//var caseBlocks []*ir.Block

	ctx.LeaveBlock = leaveBlock
	defaultContext := ctx.NewContext()
	defaultContext.Block = defaultBlock
	if s.Default != nil {
		s.Default.Body.GenerateIR(defaultContext)
	}
	if !defaultContext.Terminated {
		ctx.Terminated = false
		if defaultBlock.Term == nil {
			defaultBlock.Term = ir.NewBr(leaveBlock)
		}
	}

	//ir.NewSwitch()
	//c.Block.Term = ir.NewCondBr(i.Condition.GenerateIR(c), bodyBlock, elseBlock)
	c.Block = leaveBlock
	c.Terminated = ctx.Terminated

	//TO-DO
	return false //TO-DO Check children
}
