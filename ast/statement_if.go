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
	ctx := c.NewContext()
	ctx.Block = c.Block
	if i.Initialization != nil {
		i.Initialization.GenerateIR(ctx)
	}

	leaveBlock := c.Function.IRFunction.NewBlock("")

	bodyContext := ctx.NewContext()
	bodyContext.Block = c.Function.IRFunction.NewBlock("")
	i.Body.GenerateIR(bodyContext)
	if bodyContext.Returned {
		ctx.Returned = true
	} else if !bodyContext.Block.Terminated {
		bodyContext.Block.AddInstruction(ir.NewBr(leaveBlock))
	}

	elseContext := ctx.NewContext()
	elseContext.Block = leaveBlock
	if i.Else == nil {
		ctx.Returned = false
	} else {
		elseContext.Block = c.Function.IRFunction.NewBlock("")
		i.Else.GenerateIR(elseContext)
		ctx.Returned = elseContext.Returned
		if !elseContext.Block.Terminated {
			elseContext.Block.AddInstruction(ir.NewBr(leaveBlock))
		}
	}

	c.Block.AddInstruction(ir.NewCondBr(i.Condition.GenerateIR(c), bodyContext.Block, elseContext.Block))
	c.Block = leaveBlock
	c.Returned = ctx.Returned
}
