package expression

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Invocation struct {
	Base
	Function  Expression
	Arguments *Arguments
}

type Arguments struct {
	Base
	Arguments []Expression
	Ellipsis  int
}

func (args *Arguments) GenerateIR(c *node.Context, parent ir.Value, function *ir.Func) []ir.Value {
	arguments := []ir.Value{}
	//TO-DO if parent != nil, add and convert "this"
	if args == nil {
		return arguments
	}
	//TO-DO if parent != nil, arguments number +1
	length := len(args.Arguments)
	if parent != nil {
		length++
	}
	if length > len(function.Params) {
		c.Error(args.Position, "arguments mismatch with function paraments. too many arguments.")
		return arguments
	} else if length < len(function.Params) {
		c.Error(args.Position, "arguments mismatch with function paraments. too few arguments.")
		return arguments
	}
	for _, arg := range args.Arguments {
		arguments = append(arguments, arg.GenerateIR(c))
	}
	for i, arg := range arguments {
		if !arg.Type().Equal(function.Params[i].Typ) {
			newArg := Cast(c, arg, function.Params[i].Typ)
			if newArg == nil {
				index := i
				if parent != nil {
					index--
				}
				c.Error(args.Arguments[index].GetPosition(), fmt.Sprintf("cannot convert %s to %s", arg.Type().String(), function.Params[i].Typ.String()))
			} else {
				arguments[i] = newArg
			}
		}
	}
	return arguments
}

func (i *Invocation) Type(c *node.Context) ir.Type {
	t := i.Function.Type(c)
	if f, ok := t.(*ir.FuncType); ok {
		if f.RetType == nil {
			c.Error(i.Position, "no value is returned")
		}
		return f.RetType
	}
	c.Error(i.Position, "not a function")
	return nil
}

func (i *Invocation) GenerateIR(c *node.Context) ir.Value {
	var parent ir.Value
	var function *ir.Func
	switch t := i.Function.(type) {
	case *MemberAccess:
		d := t.GenerateIR(c)
		// TO-DO generate "this"
		if f, ok := d.(*ir.Func); ok {
			function = f
		} else {
			c.Error(t.Position, fmt.Sprintf("%s is not a function", t.Member.Name))
		}
		parent = t.GenerateParentIR(c)

	case *Identifier:
		//TO-DO if class, search member function and parent function
		d := c.FindObject(t.Name)
		if d == nil {
			c.Error(t.Position, fmt.Sprintf("%s undefined", t.Name))
		} else {
			if f, ok := d.(*ir.Func); ok {
				function = f
			} else {
				c.Error(t.Position, fmt.Sprintf("%s is not a function", t.Name))
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
