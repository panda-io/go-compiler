package ast

import (
	"fmt"

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
	var parent ir.Value
	var function *ir.Func
	switch t := i.Function.(type) {
	case *MemberAccess:
		d := t.GenerateIR(c, expected)
		// TO-DO generate "this"
		if f, ok := d.(*ir.Func); ok {
			function = f
		} else {
			c.Program.Error(t.Position, fmt.Sprintf("%s is not a function", t.Member.Name))
		}
		parent = t.GenerateParentIR(c)

	case *Identifier:
		//TO-DO if class, search member function and parent function
		d := c.FindObject(t.Name)
		if d == nil {
			c.Program.Error(t.Position, fmt.Sprintf("%s undefined", t.Name))
		} else {
			if f, ok := d.(*ir.Func); ok {
				function = f
			} else {
				c.Program.Error(t.Position, fmt.Sprintf("%s is not a function", t.Name))
			}
		}
	}

	if function != nil {
		args := i.Arguments.GenerateIR(c, parent, function)
		value := ir.NewCall(function, args...)
		c.Block.AddInstruction(value)
		return value
	}
	c.Program.Error(i.Position, "function %s not found")
	return nil
}

func (*Invocation) IsConstant(p *Program) bool {
	return false
}

func (*Invocation) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
