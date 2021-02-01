package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Invocation struct {
	ExpressionBase
	Function      Expression
	Arguments     *Arguments
	TypeArguments *TypeArguments
}

func (i *Invocation) Type(c *Context, expected ir.Type) ir.Type {
	t := i.Function.Type(c, expected)
	if ir.IsFunc(t) {
		return t.(*ir.FuncType).RetType
	}
	c.Program.Error(i.Position, "not a function type")
	return nil
}

func (i *Invocation) GenerateIR(c *Context, expected ir.Type) ir.Value {
	if IsCompilerFunction(c, i.Function) {
		return InvokeCompilerFunction(c, i)
	}
	value := i.Function.GenerateIR(c, nil)
	if value != nil {
		if call, ok := value.(*ir.InstCall); ok {
			i.Arguments.GenerateIR(c, call)
			c.Block.AddInstruction(call)
			return value
		}
	}
	c.Program.Error(i.Position, "invalid function call")
	return nil
}

func (*Invocation) IsConstant(p *Program) bool {
	return false
}

func (*Invocation) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
