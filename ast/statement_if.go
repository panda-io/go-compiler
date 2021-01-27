package ast

import "github.com/panda-foundation/go-compiler/ir"

type If struct {
	StatementBase
	Initialization Statement
	Condition      Expression
	Body           Statement
	Else           Statement
}

func (i *If) GenerateIR(c *Context) bool {
	ctx := c.NewContext()
	ctx.Block = c.Block
	if i.Initialization != nil {
		i.Initialization.GenerateIR(ctx)
	}

	leaveBlock := c.Function.IRFunction.NewBlock("")

	bodyBlock := c.Function.IRFunction.NewBlock("")
	bodyContext := ctx.NewContext()
	bodyContext.Block = bodyBlock
	i.Body.GenerateIR(bodyContext)
	if bodyContext.Terminated {
		ctx.Terminated = true
	} else {
		bodyContext.Block.Term = ir.NewBr(leaveBlock)
	}

	elseBlock := leaveBlock
	if i.Else == nil {
		ctx.Terminated = false
	} else {
		elseBlock = c.Function.IRFunction.NewBlock("")
		elseContext := ctx.NewContext()
		elseContext.Block = elseBlock
		i.Else.GenerateIR(elseContext)
		if !elseContext.Terminated {
			ctx.Terminated = false
			elseBlock.Term = ir.NewBr(leaveBlock)
		}
	}

	c.Block.Term = ir.NewCondBr(i.Condition.GenerateIR(c), bodyBlock, elseBlock)
	c.Block = leaveBlock
	c.Terminated = ctx.Terminated

	return ctx.Terminated
}
