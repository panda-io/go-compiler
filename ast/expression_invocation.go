package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type Invocation struct {
	ExpressionBase
	Function  Expression
	Arguments *Arguments
}

func (i *Invocation) Type(c *Context) ir.Type {
	t := i.Function.Type(c)
	if f, ok := t.(*ir.FuncType); ok {
		return f.RetType
	}
	c.Program.Error(i.Position, fmt.Sprintf("%s is not a function", t.Name()))
	return nil
}

func (i *Invocation) GenerateIR(c *Context) ir.Value {
	var parent ir.Value
	var function *ir.Func
	switch t := i.Function.(type) {
	case *MemberAccess:
		d := t.GenerateIR(c)
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
		//TO-DO compare sig
		//TO-DO automatically convert type if possible
		args := i.Arguments.GenerateIR(c, parent, function)
		value := ir.NewCall(function, args...)
		c.Block.AddInstruction(value)
		return value
	}
	return nil
}
