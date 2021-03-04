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

	nextBlock := c.Function.IRFunction.NewBlock("")
	bodyBlock := c.Function.IRFunction.NewBlock("")
	elseBlock := nextBlock

	bodyContext := ctx.NewContext()
	bodyContext.Block = bodyBlock
	i.Body.GenerateIR(bodyContext)
	if bodyContext.Returned {
		ctx.Returned = true
	} else if !bodyContext.Block.Terminated {
		bodyContext.Block.AddInstruction(ir.NewBr(nextBlock))
	}

	elseContext := ctx.NewContext()
	if i.Else == nil {
		ctx.Returned = false
	} else {
		elseBlock = c.Function.IRFunction.NewBlock("")
		elseContext.Block = elseBlock
		i.Else.GenerateIR(elseContext)
		ctx.Returned = elseContext.Returned
		if !elseContext.Block.Terminated {
			elseContext.Block.AddInstruction(ir.NewBr(nextBlock))
		}
	}

	var condition ir.Value
	if i.Condition.IsConstant(c.Program) {
		condition = i.Condition.GenerateConstIR(ctx.Program, ir.I1)
	} else {
		condition = i.Condition.GenerateIR(ctx, nil)
	}
	ctx.Block.AddInstruction(ir.NewCondBr(condition, bodyBlock, elseBlock))
	c.Block = nextBlock
	c.Returned = ctx.Returned
}
