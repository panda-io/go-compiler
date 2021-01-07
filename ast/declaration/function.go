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
}

func (f *Function) GenerateIR(c *node.Context) {
	function := c.Program.Module.NewFunc(f.Qualified(c.Namespace), types.TypeOf(f.ReturnType), f.Parameters.GenerateIR(c)...)
	if f.Body != nil {
		c.Block = function.NewBlock("")
		f.Body.GenerateIR(c)
	}
}

func (f *Function) GenerateDeclaration(namespace string) *ir.Func {
	params := []*ir.Param{}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(types.TypeOf(parameter.Type))
			params = append(params, param)
		}
	}
	return ir.NewFunc(f.Qualified(namespace), types.TypeOf(f.ReturnType), params...)
}
