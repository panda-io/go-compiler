package expression

import (
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Invocation struct {
	Base
	Function  Expression
	Arguments *types.Arguments
}

func (e *Invocation) GenerateIR() ir.Value {
	args := []ir.Value{}
	for _, arg := range e.Arguments.Arguments {
		args = append(args, arg.(types.Type).GenerateIR())
	}
	return ir.NewCall(e.Function.GenerateIR(), args...)
}

func (e *Invocation) IsConstant() bool {
	return false
}
