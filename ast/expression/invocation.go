package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Invocation struct {
	Base
	Function  Expression
	Arguments *types.Arguments
}

func (e *Invocation) GenerateIR(c *node.Context) ir.Value {
	//TO-DO check full name of function
	/*
		args := []*ir.Param{}
		for _, arg := range e.Arguments.Arguments {
			args = append(args, arg.(types.Type).GenerateIR())
		}
		return ir.NewCall(e.Function.GenerateIR(), args...)*/
	return nil
}
