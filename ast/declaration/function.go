package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

var (
	malloc  = ir.NewFunc("malloc", ir.NewPointerType(ir.I8), ir.NewParam(ir.I32))
	free    = ir.NewFunc("free", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	memcpy  = ir.NewFunc("memcpy", ir.NewPointerType(ir.I8), ir.NewParam(ir.NewPointerType(ir.I8)), ir.NewParam(ir.I32))
	memset  = ir.NewFunc("memset", ir.Void, ir.NewParam(ir.I32), ir.NewParam(ir.I32))
	counter = ir.NewPointerType(&ir.StructType{TypeName: "global.counter"})
)

type Function struct {
	Base
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Block
	Class          *Class

	IRParams   []*ir.Param
	IRFunction *ir.Func
}

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
		if f.ObjectName != "" {
			// generate constructor
			if f.Name.Name == node.Constructor {
				// malloc struct and set 0
				ptr := ir.NewGetElementPtr(f.Class.IRStruct.Type, ir.NewNull(ir.NewPointerType(f.Class.IRStruct.Type)), ir.NewInt(ir.I32, 1))
				ctx.Block.AddInstruction(ptr)
				size := ir.NewPtrToInt(ptr, ir.I32)
				ctx.Block.AddInstruction(size)
				address := ir.NewCall(malloc, size)
				ctx.Block.AddInstruction(address)
				ctx.Block.AddInstruction(ir.NewCall(memset, address, ir.NewInt(ir.I32, 0), size))

				// set vtable
				instance := ir.NewBitCast(address, ir.NewPointerType(f.Class.IRStruct.Type))
				ctx.Block.AddInstruction(instance)
				vtable := ir.NewGetElementPtr(f.Class.IRStruct.Type, instance, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
				ctx.Block.AddInstruction(vtable)
				ctx.Block.AddInstruction(ir.NewStore(f.Class.IRVTable.Data, vtable))

				//TO-DO set default values
				ctx.Block.Term = ir.NewRet(instance)
			}
		}
		f.Body.GenerateIR(ctx)
		if f.ObjectName != "" {
			// generate destructor
			if f.Name.Name == node.Destructor {
				//TO-DO call parent destructor

				// free struct
				address := ir.NewBitCast(f.IRParams[0], ir.NewPointerType(ir.I8))
				ctx.Block.AddInstruction(address)
				ctx.Block.AddInstruction(ir.NewCall(free, address))
			}
		}
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
	return f.IRFunction
}
