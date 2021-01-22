package ast

import (
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Function struct {
	DeclarationBase
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Block
	Class          *Class

	IRParams   []*ir.Param
	IRFunction *ir.Func
}

func (f *Function) GenerateIR(c *Context) {
	if f.Body != nil {
		ctx := c.NewContext()
		ctx.Block = f.IRFunction.NewBlock(FunctionEntry)
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
			if f.Name.Name == Constructor {
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
			if f.Name.Name == Destructor {
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

func (f *Function) GenerateDeclaration(c *Context, declarations map[string]Declaration) *ir.Func {
	if f.ObjectName != "" && f.Name.Name != Constructor {
		t := ir.NewStructType()
		t.TypeName = c.Namespace + "." + f.ObjectName
		param := ir.NewParam(ir.NewPointerType(t))
		param.LocalName = ClassThis
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
	c.AddObject(f.Qualified(c.Namespace), f.IRFunction)
	return f.IRFunction
}

//TO-DO refactor
func CompareMemberFunction(f1 *ir.FuncType, f2 *ir.FuncType) bool {
	if !f1.RetType.Equal(f2.RetType) {
		return false
	}
	if len(f1.Params) != len(f2.Params) {
		return false
	}
	for i := 1; i < len(f1.Params); i++ {
		if !f1.Params[i].Equal(f2.Params[i]) {
			return false
		}
	}
	return f1.Variadic == f2.Variadic
}
