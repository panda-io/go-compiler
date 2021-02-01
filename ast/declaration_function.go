package ast

import (
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

	AutoReleasePool []ir.Value
	//TO-DO string ...
}

func (f *Function) GenerateIRDeclaration(p *Program) *ir.Func {
	if f.HasAttribute(Compiler) {
		return nil
	}
	if f.ObjectName != "" && f.Name.Name != Constructor {
		param := ir.NewParam(ir.NewPointerType(ir.I8))
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
	if f.HasAttribute(Extern) {
		l := f.GetAttributeValue(Extern, Variadic)
		if l != nil {
			v := l.GetValue()
			if variadic, ok := v.(bool); ok {
				if variadic {
					f.IRFunction.Sig.Variadic = true
				}
			}
		}
	}
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
		for i, param := range f.IRParams {
			var v ir.Value
			if i == 0 && f.ObjectName != "" && f.Name.Name != Constructor {
				cast := ir.NewBitCast(param, CreateStructPointer(p.Module.Namespace+"."+f.ObjectName))
				f.IREntry.AddInstruction(cast)
				v = cast
			} else {
				//TO-DO add weak ref
				alloc := ir.NewAlloca(param.Typ)
				f.IREntry.AddInstruction(alloc)
				store := ir.NewStore(param, alloc)
				f.IREntry.AddInstruction(store)
				v = alloc
			}
			err := c.AddObject(param.LocalName, v)
			if err != nil {
				p.Error(f.Position, err.Error())
			}
		}

		// prepare return value
		if f.ReturnType != nil {
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
				for i, v := range current.Variables {
					if v.Value != nil {
						index := f.Class.VariableIndexes[v.Name.Name]
						offset := ir.NewGetElementPtr(f.Class.IRStruct, instance, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
						f.IREntry.AddInstruction(offset)
						f.IREntry.AddInstruction(ir.NewStore(current.IRValues[i], offset))
					}
				}
				current = current.Parent
			}
			f.IREntry.AddInstruction(ir.NewStore(instance, f.IRReturn))
		}

		f.IREntry.AddInstruction(ir.NewBr(f.IRBody))

		c.Block = f.IRBody
		f.Body.GenerateIR(c)

		if f.ObjectName != "" && f.Name.Name == Constructor {
			c.Returned = true
			c.Block.AddInstruction(ir.NewBr(f.IRExit))
		}
		if !c.Block.Terminated {
			if c.Returned || f.ReturnType == nil {
				c.Block.AddInstruction(ir.NewBr(f.IRExit))
			} else {
				c.Program.Error(f.Position, "missing return")
			}
		}

		// TO-DO clean up function variables in exit block
		// exit

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
		if f.ReturnType == nil {
			f.IRExit.AddInstruction(ir.NewRet(nil))
		} else {
			load := ir.NewLoad(f.ReturnType.Type(p), f.IRReturn)
			f.IRExit.AddInstruction(load)
			f.IRExit.AddInstruction(ir.NewRet(load))
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
}

func (args *Arguments) GenerateIR(c *Context, call *ir.InstCall) {
	function := call.Callee.Type().(*ir.PointerType).ElemType.(*ir.FuncType)
	if args == nil {
		return
	}
	length := len(args.Arguments) + len(call.Args)
	if length < len(function.Params) {
		c.Program.Error(args.Position, "too few arguments")
		return
	} else if length > len(function.Params) && !function.Variadic {
		c.Program.Error(args.Position, "too many arguments")
		return
	}
	for _, arg := range args.Arguments {
		i := len(call.Args)
		var v ir.Value
		if i < len(function.Params) {
			v = arg.GenerateIR(c, function.Params[i])
		} else {
			v = arg.GenerateIR(c, nil)
		}
		if v == nil {
			c.Program.Error(arg.GetPosition(), "invalid expression")
		} else {
			call.Args = append(call.Args, c.AutoLoad(arg.GenerateIR(c, nil)))
		}
	}
}
