package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Function struct {
	Base
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Block
	ClassName      string
}

func (f *Function) GenerateIR(c *node.Context) {
	function := c.Module.NewFunc(f.Name.Name, types.TypeOf(f.ReturnType), f.Parameters.GenerateIR(c)...)
	if f.Body != nil {
		c.Block = function.NewBlock("")
		f.Body.GenerateIR(c)
	}
}

func (f *Function) GenerateIRDeclaration() ir.Value {
	params := []*ir.Param{}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(types.TypeOf(parameter.Type))
			params = append(params, param)
		}
	}
	return ir.NewFunc(f.Name.Name, types.TypeOf(f.ReturnType), params...)
}
