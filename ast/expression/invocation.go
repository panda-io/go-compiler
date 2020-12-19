package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Invocation struct {
	Base
	Function  Expression
	Arguments *types.Arguments
}

func (e *Invocation) GenerateIR(c *node.Context) {
	/*
		args := []ir.Value{}
		for _, arg := range e.Arguments.Arguments {
			args = append(args, arg.(types.Type).GenerateIR())
		}
		return ir.NewCall(e.Function.GenerateIR(), args...)*/
}
