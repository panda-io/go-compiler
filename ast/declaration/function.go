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

	IRParams   []*ir.Param
	IRFunction *ir.Func
}

func (f *Function) GenerateIR(c *node.Context) {
	if f.Body != nil {
		ctx := c.NewContext()
		ctx.Block = f.IRFunction.NewBlock("entry")
		for _, param := range f.IRParams {
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

func (f *Function) GenerateDeclaration(c *node.Context, declarations map[string]Declaration) *ir.Func {
	if f.ObjectName != "" {
		t := ir.NewStructType()
		t.TypeName = c.Namespace + "." + f.ObjectName
		param := ir.NewParam(ir.NewPointerType(t))
		param.LocalName = node.This
		f.IRParams = append(f.IRParams, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(TypeOf(c, declarations, parameter.Type))
			param.LocalName = parameter.Name
			f.IRParams = append(f.IRParams, param)
		}
	}
	f.IRFunction = c.Program.Module.NewFunc(f.Qualified(c.Namespace), TypeOf(c, declarations, f.ReturnType), f.IRParams...)
	n := f.Name.Name
	if f.ObjectName != "" {
		n = f.ObjectName + "." + n
	}
	err := c.AddObject(n, f.IRFunction)
	if err != nil {
		c.Error(f.Position, err.Error())
	}
	return f.IRFunction
}
