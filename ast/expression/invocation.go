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

func (args *Arguments) GenerateIR(c *node.Context) []ir.Value {
	arguments := []ir.Value{}
	if args == nil {
		return arguments
	}
	for _, arg := range args.Arguments {
		arguments = append(arguments, arg.GenerateIR(c))
	}
	return arguments
}

func (i *Invocation) GenerateIR(c *node.Context) ir.Value {
	var function *ir.Func
	switch t := i.Function.(type) {
	case *MemberAccess:
		//TO-DO
		panic("not implement")
		// search import

	case *Identifier:
		//TO-DO if class, search member function and parent function
		d := c.FindDelaration(t.Name)
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

	//TO-DO compare sig
	if function != nil {
		args := i.Arguments.GenerateIR(c)
		value := ir.NewCall(function, args...)
		c.Block.AddInstruction(value)
		return value
	}
	return nil
}
