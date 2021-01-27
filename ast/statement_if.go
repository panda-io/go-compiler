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
	if bodyContext.Returned {
		ctx.Returned = true
	} else {
		bodyContext.Block.AddInstruction(ir.NewBr(leaveBlock))
	}

	elseBlock := leaveBlock
	if i.Else == nil {
		ctx.Returned = false
	} else {
		elseBlock = c.Function.IRFunction.NewBlock("")
		elseContext := ctx.NewContext()
		elseContext.Block = elseBlock
		i.Else.GenerateIR(elseContext)
		ctx.Returned = elseContext.Returned
		if !elseContext.Returned {
			elseBlock.AddInstruction(ir.NewBr(leaveBlock))
		}
	}

	c.Block.AddInstruction(ir.NewCondBr(i.Condition.GenerateIR(c), bodyBlock, elseBlock))
	c.Block = leaveBlock
	c.Returned = ctx.Returned

	return ctx.Returned
}
