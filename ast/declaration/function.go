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

/*
define void @add(i32, i32) #0 {
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  store i32 %0, i32* %3, align 4
  store i32 %1, i32* %4, align 4
  store i32 1, i32* %3, align 4
  ret void
}
*/
func (f *Function) GenerateIR(c *node.Context) {
	if f.Body != nil {
		ctx := c.NewContext()
		ctx.Block = f.IRFunction.NewBlock("entry")
		for _, param := range f.IRParams {
			var value ir.Value
			if ir.IsPointer(param.Typ) {
				value = param
			} else {
				alloc := ir.NewAlloca(param.Typ)
				ctx.Block.AddInstruction(alloc)
				store := ir.NewStore(param, alloc)
				ctx.Block.AddInstruction(store)
				value = alloc
			}
			err := ctx.AddObject(param.LocalName, value)
			if err != nil {
				c.Error(f.Position, err.Error())
			}
		}
		f.Body.GenerateIR(ctx)
		if ctx.Block.Term == nil {
			ctx.Block.Term = ir.NewRet(nil)
		}
	}
}

func (f *Function) GenerateDeclaration(c *node.Context, declarations map[string]Declaration) *ir.Func {
	if f.ObjectName != "" && f.Name.Name != node.Constructor {
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
