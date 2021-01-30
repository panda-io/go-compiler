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

	leaveBlock := c.Function.IRFunction.NewBlock("")
	ctx.LeaveBlock = leaveBlock

	defaultContext := ctx.NewContext()
	defaultContext.Block = c.Function.IRFunction.NewBlock("")
	if s.Default != nil {
		s.Default.Body.GenerateIR(defaultContext)
	}
	if !defaultContext.Block.Terminated {
		defaultContext.Block.AddInstruction(ir.NewBr(leaveBlock))
	}
	if !defaultContext.Returned {
		ctx.Returned = false
	}

	var caseBlocks []*ir.Case
	for _, cc := range s.Cases {
		caseContext := ctx.NewContext()
		caseContext.Block = c.Function.IRFunction.NewBlock("")
		cc.Body.GenerateIR(caseContext)
		if !caseContext.Returned {
			ctx.Returned = false
		}
		if !cc.Case.IsConstant(c.Program) {
			c.Program.Error(cc.Position, "expect constant int expression")
		}
		caseBlocks = append(caseBlocks, ir.NewCase(cc.Case.GenerateConstIR(c.Program, t.(*ir.IntType)), caseContext.Block))
	}

	for i, cc := range caseBlocks {
		b := cc.Target.(*ir.Block)
		if !b.Terminated {
			if i == len(caseBlocks)-1 {
				// last one
				b.AddInstruction(ir.NewBr(defaultContext.Block))
			} else {
				b.AddInstruction(ir.NewBr(caseBlocks[i+1].Target))
			}
		}
	}

	c.Block.AddInstruction(ir.NewSwitch(operand, defaultContext.Block, caseBlocks...))
	c.Block = leaveBlock
	c.Returned = ctx.Returned
}
