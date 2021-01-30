package ast

import "github.com/panda-foundation/go-compiler/ir"

type For struct {
	StatementBase
	Initialization Statement
	Condition      Expression
	Post           Statement
	Body           Statement
}

func (f *For) GenerateIR(c *Context) {
	ctx := c.NewContext()
	ctx.Block = c.Block
	if f.Initialization != nil {
		f.Initialization.GenerateIR(ctx)
	}

	leaveBlock := c.Function.IRFunction.NewBlock("")
	ctx.LeaveBlock = leaveBlock

	conditionBlock := c.Function.IRFunction.NewBlock("")
	conditionContext := ctx.NewContext()
	conditionContext.Block = conditionBlock

	postBlock := c.Function.IRFunction.NewBlock("")
	ctx.LoopBlock = postBlock
	postContext := ctx.NewContext()
	postContext.Block = postBlock
	if f.Post != nil {
		f.Post.GenerateIR(postContext)
	}
	postContext.Block.AddInstruction(ir.NewBr(conditionBlock))

	bodyBlock := c.Function.IRFunction.NewBlock("")
	bodyContext := ctx.NewContext()
	bodyContext.Block = bodyBlock
	f.Body.GenerateIR(bodyContext)
	if bodyContext.Returned {
		ctx.Returned = true
	} else if !bodyContext.Block.Terminated {
		bodyContext.Block.AddInstruction(ir.NewBr(postBlock))
	}

	var condition ir.Value
	if f.Condition.IsConstant(c.Program) {
		condition = f.Condition.GenerateConstIR(conditionContext.Program, ir.I1)
	} else {
		condition = f.Condition.GenerateIR(conditionContext, nil)
	}
	conditionContext.Block.AddInstruction(ir.NewCondBr(condition, bodyBlock, leaveBlock))
	ctx.Block.AddInstruction(ir.NewBr(conditionBlock))
	c.Block = leaveBlock
	c.Returned = ctx.Returned
}
