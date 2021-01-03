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
		functions := c.FindDelaration(t.Name)
		if len(functions) == 1 {
			if f, ok := functions[0].(*ir.Func); ok {
				function = f
			} else {
				c.Error(t.Position, fmt.Sprintf("%s is not a function", t.Name))
			}
		} else if len(functions) == 0 {
			c.Error(t.Position, fmt.Sprintf("%s undefined", t.Name))
		} else {
			c.Error(t.Position, fmt.Sprintf("ambiguous variable %s", t.Name))
		}
	}

	if function != nil {
		args := i.Arguments.GenerateIR(c)
		value := ir.NewCall(function, args...)
		c.Block.AddInstruction(value)
		return value
	}
	return nil
}
