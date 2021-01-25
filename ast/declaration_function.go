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
	IREntry    *ir.Block
	IRBody     *ir.Block
	IRExit     *ir.Block
	IRReturn   ir.Value
}

func (f *Function) GenerateIRDeclaration(p *Program) *ir.Func {
	if f.ObjectName != "" && f.Name.Name != Constructor {
		param := ir.NewParam(CreateStructPointerType(p.Module.Namespace + "." + f.ObjectName))
		param.LocalName = ClassThis
		f.IRParams = append(f.IRParams, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(parameter.Type.Type(p))
			param.LocalName = parameter.Name
			f.IRParams = append(f.IRParams, param)
		}
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = f.ReturnType.Type(p)
	}
	f.IRFunction = p.IRModule.NewFunc(f.Qualified(p.Module.Namespace), t, f.IRParams...)
	return f.IRFunction
}

func (f *Function) GenerateIR(p *Program) {
	if f.Body != nil {
		c := NewContext(p)
		c.Function = f
		f.IREntry = f.IRFunction.NewBlock(FunctionEntry)
		f.IRBody = f.IRFunction.NewBlock(FunctionBody)
		f.IRExit = f.IRFunction.NewBlock(FunctionExit)
		c.Block = f.IREntry

		// prepare params
		for _, param := range f.IRParams {
			var value ir.Value
			if ir.IsPointer(param.Typ) {
				value = param
			} else {
				alloc := ir.NewAlloca(param.Typ)
				f.IREntry.AddInstruction(alloc)
				store := ir.NewStore(param, alloc)
				f.IREntry.AddInstruction(store)
				value = alloc
			}
			err := c.AddObject(param.LocalName, value)
			if err != nil {
				p.Error(f.Position, err.Error())
			}
		}
		f.IREntry.Term = ir.NewBr(f.IRBody)

		// prepare return value
		if f.ReturnType == nil {
			f.IRExit.Term = ir.NewRet(nil)
		} else {
			alloca := ir.NewAlloca(f.ReturnType.Type(p))
			f.IREntry.AddInstruction(alloca)
			f.IRReturn = alloca
		}

		// generate constructor
		if f.ObjectName != "" && f.Name.Name == Constructor {
			// malloc struct and set 0
			ptr := ir.NewGetElementPtr(f.Class.IRStruct, ir.NewNull(ir.NewPointerType(f.Class.IRStruct)), ir.NewInt(ir.I32, 1))
			f.IREntry.AddInstruction(ptr)
			size := ir.NewPtrToInt(ptr, ir.I32)
			f.IREntry.AddInstruction(size)
			address := ir.NewCall(malloc, size)
			f.IREntry.AddInstruction(address)
			f.IREntry.AddInstruction(ir.NewCall(memset, address, ir.NewInt(ir.I32, 0), size))

			// set vtable
			instance := ir.NewBitCast(address, ir.NewPointerType(f.Class.IRStruct))
			f.IREntry.AddInstruction(instance)
			vtable := ir.NewGetElementPtr(f.Class.IRStruct, instance, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			f.IREntry.AddInstruction(vtable)
			f.IREntry.AddInstruction(ir.NewStore(f.Class.IRVTableData, vtable))

			// set default values
			current := f.Class
			for current != nil {
				for _, v := range current.Variables {
					if v.Value != nil {
						value := v.Value.GenerateIR(c) // constant expr //TO-DO use generate const expr instead
						if !ir.IsConstant(value) {
							p.Error(v.Position, "initialize value must be const expression")
						}
						index := f.Class.VariableIndexes[v.Name.Name]
						offset := ir.NewGetElementPtr(f.Class.IRStruct, instance, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
						f.IREntry.AddInstruction(offset)
						f.IREntry.AddInstruction(ir.NewStore(value, offset))
					}
				}
				current = current.Parent
			}
			f.IREntry.AddInstruction(ir.NewStore(instance, f.IRReturn))
		}

		c.Block = f.IRBody
		f.Body.GenerateIR(c)

		if f.ObjectName != "" && f.Name.Name == Constructor {
			c.Terminated = true
		}
		if f.ReturnType != nil && !c.Terminated {
			c.Program.Error(f.Position, "missing return")
		}
		c.Block.Term = ir.NewBr(f.IRExit)

		// TO-DO clean up function variables in exit block

		// generate destructor
		if f.ObjectName != "" && f.Name.Name == Destructor {
			// TO-DO call parent destructor
			// TO-DO clean up members
			// TO-DO check if free memory
			address := ir.NewBitCast(f.IRParams[0], ir.NewPointerType(ir.I8))
			f.IRExit.AddInstruction(address)
			f.IRExit.AddInstruction(ir.NewCall(free, address))
		}

		// return
		if f.ReturnType != nil {
			load := ir.NewLoad(f.ReturnType.Type(p), f.IRReturn)
			f.IRExit.AddInstruction(load)
			f.IRExit.Term = ir.NewRet(load)
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
func (args *Arguments) GenerateIR(c *Context, this ir.Value, function *ir.Func) []ir.Value {
	arguments := []ir.Value{}
	//TO-DO if parent != nil, add and convert "this"
	if args == nil {
		return arguments
	}
	//TO-DO if parent != nil, arguments number +1
	length := len(args.Arguments)
	if this != nil {
		length++
	}
	if length > len(function.Params) {
		c.Program.Error(args.Position, "arguments mismatch with function paraments. too many arguments.")
		return arguments
	} else if length < len(function.Params) {
		c.Program.Error(args.Position, "arguments mismatch with function paraments. too few arguments.")
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
				if this != nil {
					index--
				}
				c.Program.Error(args.Arguments[index].GetPosition(), fmt.Sprintf("cannot convert %s to %s", arg.Type().String(), function.Params[i].Typ.String()))
			} else {
				arguments[i] = newArg
			}
		}
	}
	return arguments
}
