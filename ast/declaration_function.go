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

	AutoReleasePool    []ir.Value
	BuiltinReleasePool []ir.Value
}

func (f *Function) GenerateIRDeclaration(p *Program) *ir.Func {
	if IsCompilerFunction(f.Qualified(p.Module.Namespace)) {
		return nil
	}
	if f.ObjectName != "" && f.Name.Name != Constructor {
		param := ir.NewParam(pointerType)
		param.LocalName = ClassThis
		f.IRParams = append(f.IRParams, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(parameter.Type.Type(p))
			param.LocalName = parameter.Name
			if t, ok := parameter.Type.(*TypeName); ok {
				userData, _ := p.FindDeclaration(t)
				SetUserData(param, userData)
			}
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
		for _, param := range f.IRParams {
			var v ir.Value
			if param.Type().Equal(pointerType) {
				//TO-DO add shared ref //TO-DO string
				v = param
			} else {
				alloc := ir.NewAlloca(param.Typ)
				CopyUserData(param, alloc)
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
			instance := CastFromPointer(f.IREntry, address, ir.NewPointerType(f.Class.IRStruct))
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
			f.IREntry.AddInstruction(ir.NewStore(address, f.IRReturn))
		}

		f.IREntry.AddInstruction(ir.NewBr(f.IRBody))

		if f.ObjectName != "" && f.Name.Name == Destructor {
			// call parent destructor
			if f.Class.Parent != nil {
				destructor := f.Class.Parent.IRFunctions[1]
				call := ir.NewCall(destructor, f.IRParams[0])
				f.IRBody.AddInstruction(call)
			}
		}
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

		// generate destructor
		if f.ObjectName != "" && f.Name.Name == Destructor {
			// TO-DO clean up members // none-builtin-member
		}

		// auto release pool
		for _, obj := range f.BuiltinReleasePool {
			qualified := ""
			switch t := obj.(type) {
			case *ir.InstCall:
				qualified = GetUserData(t)

			case *ir.InstAlloca:
				qualified = GetUserData(t)
				load := ir.NewLoad(t.ElemType, t)
				f.IRExit.AddInstruction(load)
				obj = load
			}
			class := c.Program.FindQualified(qualified).(*Class)
			class.DestroyInstance(f.IRExit, obj)
		}
		for _, obj := range f.AutoReleasePool {
			obj = AutoLoad(obj, f.IRExit)
			call := ir.NewCall(releaseShared, obj)
			f.IRExit.AddInstruction(call)
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
