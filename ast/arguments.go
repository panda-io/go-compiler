package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type Arguments struct {
	NodeBase
	Arguments []Expression
	Ellipsis  int
}

func (args *Arguments) GenerateIR(c *Context, parent ir.Value, function *ir.Func) []ir.Value {
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
