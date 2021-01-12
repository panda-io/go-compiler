package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Parameters struct {
	Base
	Parameters []*Parameter
	Ellipsis   bool
}

func (p *Parameters) GenerateIR(c *node.Context, structType *ir.StructType) []*ir.Param {
	params := []*ir.Param{}
	if structType != nil {
		param := ir.NewParam(ir.NewPointerType(structType))
		param.LocalName = node.This
		params = append(params, param)
	}
	if p == nil {
		return params
	}
	for _, parameter := range p.Parameters {
		param := ir.NewParam(TypeOf(parameter.Type))
		param.LocalName = parameter.Name
		params = append(params, param)
	}
	return params
}

type Parameter struct {
	Base
	Name string
	Type Type
}
