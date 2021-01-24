package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

//TO-DO validate function declaration and actual return type

type Function struct {
	DeclarationBase
	TypeParameters *TypeParameters
	Parameters     *Parameters
	ReturnType     Type
	Body           *Block

	Class *Class

	IRParams   []*ir.Param
	IRFunction *ir.Func
	Entry      *ir.Block
	Exit       *ir.Block
	Return     ir.Value
}

func (f *Function) GenerateIRDeclaration(c *Context) *ir.Func {
	if f.ObjectName != "" && f.Name.Name != Constructor {
		param := ir.NewParam(c.StructPointer(c.Module.Namespace + "." + f.ObjectName))
		param.LocalName = ClassThis
		f.IRParams = append(f.IRParams, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(parameter.Type.Type(c))
			param.LocalName = parameter.Name
			f.IRParams = append(f.IRParams, param)
		}
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = f.ReturnType.Type(c)
	}
	f.IRFunction = c.Program.Module.NewFunc(f.Qualified(c.Module.Namespace), t, f.IRParams...)
	if f.ObjectName == "" {
		err := c.AddObject(f.Qualified(c.Module.Namespace), f.IRFunction)
		if err != nil {
			c.Error(f.Position, err.Error())
		}
	}
	return f.IRFunction
}

func (f *Function) GenerateIR(c *Context) {
	if f.Body != nil {
		ctx := c.NewContext()
		ctx.Function = f
		f.Entry = f.IRFunction.NewBlock(FunctionEntry)
		f.Exit = f.IRFunction.NewBlock(FunctionExit)
		ctx.Block = f.Entry

		// prepare params
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

		// prepare return value
		if f.ReturnType == nil {
			f.Exit.Term = ir.NewRet(nil)
		} else {
			alloca := ir.NewAlloca(f.ReturnType.Type(c))
			f.Entry.AddInstruction(alloca)
			f.Return = alloca
		}

		// generate constructor
		if f.ObjectName != "" && f.Name.Name == Constructor {
			// malloc struct and set 0
			ptr := ir.NewGetElementPtr(f.Class.IRStruct, ir.NewNull(ir.NewPointerType(f.Class.IRStruct)), ir.NewInt(ir.I32, 1))
			ctx.Block.AddInstruction(ptr)
			size := ir.NewPtrToInt(ptr, ir.I32)
			ctx.Block.AddInstruction(size)
			address := ir.NewCall(malloc, size)
			ctx.Block.AddInstruction(address)
			ctx.Block.AddInstruction(ir.NewCall(memset, address, ir.NewInt(ir.I32, 0), size))

			// set vtable
			instance := ir.NewBitCast(address, ir.NewPointerType(f.Class.IRStruct))
			ctx.Block.AddInstruction(instance)
			vtable := ir.NewGetElementPtr(f.Class.IRStruct, instance, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			ctx.Block.AddInstruction(vtable)
			ctx.Block.AddInstruction(ir.NewStore(f.Class.IRVTableData, vtable))

			// set default values
			current := f.Class
			for current != nil {
				for _, v := range current.Variables {
					if v.Value != nil {
						value := v.Value.GenerateIR(c) // constant expr
						if !ir.IsConstant(value) {
							ctx.Error(v.Position, "initialize value must be const expression")
						}
						index := f.Class.VariableIndexes[v.Name.Name]
						offset := ir.NewGetElementPtr(f.Class.IRStruct, instance, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
						ctx.Block.AddInstruction(offset)
						ctx.Block.AddInstruction(ir.NewStore(value, offset))
					}
				}
				current = current.Parent
			}
			ctx.Block.AddInstruction(ir.NewStore(instance, ctx.Function.Return))
		}

		f.Body.GenerateIR(ctx)
		ctx.Block.Term = ir.NewBr(f.Exit)

		// TO-DO clean up function variables in exit block

		// generate destructor
		if f.ObjectName != "" && f.Name.Name == Destructor {
			// TO-DO call parent destructor
			// TO-DO clean up members
			// TO-DO check if free memory
			address := ir.NewBitCast(f.IRParams[0], ir.NewPointerType(ir.I8))
			f.Exit.AddInstruction(address)
			f.Exit.AddInstruction(ir.NewCall(free, address))
		}

		// return
		if f.ReturnType != nil {
			load := ir.NewLoad(f.ReturnType.Type(c), f.Return)
			f.Exit.AddInstruction(load)
			f.Exit.Term = ir.NewRet(load)
		}
	}
}

type Parameters struct {
	NodeBase
	Parameters []*Parameter
	Ellipsis   bool
}

type Parameter struct {
	NodeBase
	Name string
	Type Type
}

type Arguments struct {
	NodeBase
	Arguments []Expression
	Ellipsis  int
}

// TO-DO not finished
func (args *Arguments) GenerateIR(c *Context, parent ir.Value, function *ir.Func) []ir.Value {
	arguments := []ir.Value{}
	//TO-DO if parent != nil, add and convert "this"
	if args == nil {
		return arguments
	}
	//TO-DO if parent != nil, arguments number +1
	length := len(args.Arguments)
	if parent != nil {
		length++
	}
	if length > len(function.Params) {
		c.Error(args.Position, "arguments mismatch with function paraments. too many arguments.")
		return arguments
	} else if length < len(function.Params) {
		c.Error(args.Position, "arguments mismatch with function paraments. too few arguments.")
		return arguments
	}
	for _, arg := range args.Arguments {
		arguments = append(arguments, arg.GenerateIR(c))
	}
	for i, arg := range arguments {
		if !arg.Type().Equal(function.Params[i].Typ) {
			newArg := Cast(c, arg, function.Params[i].Typ)
			if newArg == nil {
				index := i
				if parent != nil {
					index--
				}
				c.Error(args.Arguments[index].GetPosition(), fmt.Sprintf("cannot convert %s to %s", arg.Type().String(), function.Params[i].Typ.String()))
			} else {
				arguments[i] = newArg
			}
		}
	}
	return arguments
}
