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

func (s *Switch) GenerateIR(c *Context) {
	ctx := c.NewContext()
	ctx.Block = c.Block
	ctx.Returned = true
	if s.Initialization != nil {
		s.Initialization.GenerateIR(ctx)
	}
	var operand ir.Value
	if s.Operand.IsConstant(c.Program) {
		operand = s.Operand.GenerateConstIR(c.Program, ir.I1)
	} else {
		operand = c.AutoLoad(s.Operand.GenerateIR(c, nil))
	}
	t := operand.Type()
	if !ir.IsInt(t) {
		c.Program.Error(s.Operand.GetPosition(), "switch operand must be integer type")
		return
	}

	nextBlock := c.Function.IRFunction.NewBlock("")
	ctx.LeaveBlock = nextBlock

	defaultContext := ctx.NewContext()
	defaultBlock := c.Function.IRFunction.NewBlock("")
	defaultContext.Block = defaultBlock
	if s.Default != nil {
		s.Default.Body.GenerateIR(defaultContext)
	}
	if !defaultContext.Block.Terminated {
		defaultContext.Block.AddInstruction(ir.NewBr(nextBlock))
	}
	if !defaultContext.Returned {
		ctx.Returned = false
	}

	var caseBlocks []*ir.Case
	for _, cc := range s.Cases {
		caseContext := ctx.NewContext()
		caseBlock := c.Function.IRFunction.NewBlock("")
		caseContext.Block = caseBlock
		cc.Body.GenerateIR(caseContext)
		if !caseContext.Returned {
			ctx.Returned = false
		}
		if !cc.Case.IsConstant(c.Program) {
			c.Program.Error(cc.Position, "expect constant int expression")
		}
		caseBlocks = append(caseBlocks, ir.NewCase(cc.Case.GenerateConstIR(c.Program, t.(*ir.IntType)), caseBlock))
	}

	for i, cc := range caseBlocks {
		b := cc.Target.(*ir.Block)
		if !b.Terminated {
			if i == len(caseBlocks)-1 {
				// last one
				b.AddInstruction(ir.NewBr(defaultBlock))
			} else {
				b.AddInstruction(ir.NewBr(caseBlocks[i+1].Target))
			}
		}
	}

	ctx.Block.AddInstruction(ir.NewSwitch(operand, defaultBlock, caseBlocks...))
	c.Block = nextBlock
	c.Returned = ctx.Returned
}
