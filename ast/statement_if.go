package ast

import "github.com/panda-foundation/go-compiler/ir"

type If struct {
	StatementBase
	Initialization Statement
	Condition      Expression
	Body           Statement
	Else           Statement
}

func (i *If) GenerateIR(c *Context) {
	if i.Initialization != nil {
		i.Initialization.GenerateIR(c)
	}

	leaveBlock := c.Function.IRFunction.NewBlock("")
	leaveContext := c.NewContext()
	leaveContext.Block = leaveBlock

	bodyBlock := c.Function.IRFunction.NewBlock("")
	bodyContext := c.NewContext()
	bodyContext.Block = bodyBlock
	i.Body.GenerateIR(bodyContext)
	if bodyBlock.Term == nil {
		bodyBlock.Term = ir.NewBr(leaveBlock)
	}

	elseBlock := leaveBlock
	if i.Else != nil {
		elseBlock = c.Function.IRFunction.NewBlock("")
		elseContext := c.NewContext()
		elseContext.Block = elseBlock
		i.Else.GenerateIR(elseContext)
		if elseBlock.Term == nil {
			elseBlock.Term = ir.NewBr(leaveBlock)
		}
	}

	c.Block.Term = ir.NewCondBr(i.Condition.GenerateIR(c), bodyBlock, elseBlock)
	c.Block = leaveBlock
}
