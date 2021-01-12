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
		params = append(params, ir.NewParam(ir.NewPointerType(structType)))
	}
	if p == nil {
		return params
	}
	for _, parameter := range p.Parameters {
		param := ir.NewParam(TypeOf(parameter.Type))
		params = append(params, param)
		err := c.AddVariable(parameter.Name, param)
		if err != nil {
			c.Error(parameter.Position, err.Error())
		}
	}
	return params
}

type Parameter struct {
	Base
	Name string
	Type Type
}
