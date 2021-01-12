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
	var s *ir.StructType
	if f.ObjectName != "" {
		s = ir.NewStructType()
		s.TypeName = c.Namespace + "." + f.ObjectName
	}
	params := f.Parameters.GenerateIR(c, s)
	function := c.Program.Module.NewFunc(f.Qualified(c.Namespace), types.TypeOf(f.ReturnType), params...)
	if f.Body != nil {
		ctx := c.NewContext()
		ctx.Block = function.NewBlock("entry")
		for _, param := range params {
			load := ir.NewLoad(param.Type(), param)
			ctx.Block.AddInstruction(load)
			err := ctx.AddObject(param.LocalName, load)
			if err != nil {
				c.Error(f.Position, err.Error())
			}
		}
		f.Body.GenerateIR(ctx)
		if ctx.Block.Term == nil {
			ctx.Block.Term = ir.NewRet(ir.NewNull(nil))
		}
	}
}

func (f *Function) GenerateDeclaration(c *node.Context) {
	params := []*ir.Param{}
	if f.ObjectName != "" {
		t := ir.NewStructType()
		t.TypeName = c.Namespace + "." + f.ObjectName
		param := ir.NewParam(ir.NewPointerType(t))
		params = append(params, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(types.TypeOf(parameter.Type))
			params = append(params, param)
		}
	}
	d := ir.NewFunc(f.Qualified(c.Namespace), types.TypeOf(f.ReturnType), params...)
	err := c.AddObject(f.Name.Name, d)
	if err != nil {
		c.Error(f.Position, err.Error())
	}
}
