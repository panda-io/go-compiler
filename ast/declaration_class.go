package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Class struct {
	DeclarationBase
	TypeParameters *TypeParameters
	Parents        []*TypeName
	Functions      []*Function
	Variables      []*Variable

	Parent     *Class
	Interfaces []*Interface

	IRStruct *ir.StructType
	// only variables in current class, same content with "Variables"
	IRVariables          []ir.Type
	IRValues             []ir.Value
	AccumulatedVariables []ir.Type
	// index to the total varaibles defined in IRStruct (also store in AccumulatedVariables)
	VariableIndexes map[string]int

	IRVTable *ir.StructType
	// only functions in current class, same content with "Functions"
	IRFunctions          []*ir.Func
	IRVTableData         *ir.Global
	AccumulatedFunctions []*ir.Func
	// index to the total functions defined in IRVTable (also store in AccumulatedFunctions)
	FunctionIndexes map[string]int
}

func (c *Class) FindVariableFuncType(name string) *ir.Func {
	for _, v := range c.Variables {
		if v.Name.Name == name {
			return v.FuncType
		}
	}
	if c.Parent != nil {
		return c.Parent.FindVariableFuncType(name)
	}
	return nil
}

func (c *Class) AddVariable(v *Variable) error {
	for _, variable := range c.Variables {
		if v.Name.Name == variable.Name.Name {
			return fmt.Errorf("%s redeclared", v.Name.Name)
		}
	}
	for _, function := range c.Functions {
		if v.Name.Name == function.Name.Name {
			return fmt.Errorf("%s redeclared", v.Name.Name)
		}
	}
	c.Variables = append(c.Variables, v)
	return nil
}

func (c *Class) AddFunction(f *Function) error {
	for _, variable := range c.Variables {
		if f.Name.Name == variable.Name.Name {
			return fmt.Errorf("%s redeclared", f.Name.Name)
		}
	}
	for _, function := range c.Functions {
		if f.Name.Name == function.Name.Name {
			return fmt.Errorf("%s redeclared", f.Name.Name)
		}
	}
	c.Functions = append(c.Functions, f)
	return nil
}

func (c *Class) GenerateIRDeclaration(p *Program) {
	for _, v := range c.Variables {
		if _, ok := v.Type.(*TypeFunction); ok {
			v.FuncType = v.Type.(*TypeFunction).Define(p)
		}
		c.IRVariables = append(c.IRVariables, GetIRType(v.Type, p, false))
		if v.Value == nil {
			c.IRValues = append(c.IRValues, nil)
		} else {
			c.IRValues = append(c.IRValues, v.Value.GenerateConstIR(p, GetIRType(v.Type, p, false)))
		}
	}
	for _, f := range c.Functions {
		c.IRFunctions = append(c.IRFunctions, f.GenerateIRDeclaration(p))
	}
}

func (c *Class) GenerateIRStruct(p *Program) {
	c.VariableIndexes = make(map[string]int)

	c.AccumulatedVariables = []ir.Type{CreateStructPointer(p.Module.Namespace + "." + c.Name.Name + ".vtable.type")}
	classes := []*Class{c}
	current := c
	for current.Parent != nil {
		classes = append(classes, current.Parent)
		current = current.Parent
	}
	index := 1
	for i := len(classes) - 1; i > -1; i-- {
		current = classes[i]
		for j, v := range current.Variables {
			c.AccumulatedVariables = append(c.AccumulatedVariables, current.IRVariables[j])
			if _, ok := c.VariableIndexes[v.Name.Name]; ok {
				p.Error(v.Position, fmt.Sprintf("duplicate class member: %s", v.Name.Name))
			} else {
				c.VariableIndexes[v.Name.Name] = index
			}
			index++
		}
	}

	qualified := c.Qualified(p.Module.Namespace)
	c.IRStruct = ir.NewStructType(c.AccumulatedVariables...)
	p.IRModule.NewTypeDef(qualified, c.IRStruct)
}

func (c *Class) GenerateIRVTable(p *Program) {
	c.FunctionIndexes = make(map[string]int)

	c.AccumulatedFunctions = []*ir.Func{}
	classes := []*Class{c}
	current := c
	for current.Parent != nil {
		classes = append(classes, current.Parent)
		current = current.Parent
	}
	index := 0
	for i := len(classes) - 1; i > -1; i-- {
		current = classes[i]
		for j, f := range current.Functions {
			if existing, ok := c.FunctionIndexes[f.Name.Name]; ok {
				// existing function
				function := c.AccumulatedFunctions[existing]
				if !function.Sig.Equal(current.IRFunctions[j].Sig) {
					p.Error(f.Position, fmt.Sprintf("member function %s does not match its parent class", f.Name.Name))
					//TO-DO print more params details here
				} else {
					c.AccumulatedFunctions[existing] = current.IRFunctions[j]
				}
			} else {
				// new function
				c.AccumulatedFunctions = append(c.AccumulatedFunctions, current.IRFunctions[j])
				c.FunctionIndexes[f.Name.Name] = index
				index++
			}
		}
	}

	var types []ir.Type
	var constants []ir.Constant
	for _, f := range c.AccumulatedFunctions {
		types = append(types, ir.NewPointerType(f.Sig))
		constants = append(constants, f)
	}
	c.IRVTable = ir.NewStructType(types...)
	p.IRModule.NewTypeDef(c.Qualified(p.Module.Namespace)+".vtable.type", c.IRVTable)

	data := ir.NewStruct(CreateStruct(c.Qualified(p.Module.Namespace)+".vtable.type"), constants...)
	c.IRVTableData = p.IRModule.NewGlobalDef(c.Qualified(p.Module.Namespace)+".vtable.data", data)
}

func (c *Class) GenerateIR(p *Program) {
	for _, v := range c.Functions {
		v.GenerateIR(p)
	}
}

func (c *Class) PreProcess(*Program) {
	// first is constructor, second is destructor
	functions := []*Function{nil, nil}
	for _, f := range c.Functions {
		f.Class = c
		if f.Name.Name == Constructor {
			functions[0] = f
		} else if f.Name.Name == Destructor {
			functions[1] = f
		} else {
			functions = append(functions, f)
		}
	}
	c.Functions = functions
	if c.Functions[0] == nil {
		c.Functions[0] = c.CreateEmptyFunction(Constructor)
	}
	c.Functions[0].ReturnType = &BuitinType{
		Token: token.Pointer,
	}
	if c.Functions[1] == nil {
		c.Functions[1] = c.CreateEmptyFunction(Destructor)
	}
}

func (c *Class) CreateEmptyFunction(name string) *Function {
	f := &Function{}
	f.ObjectName = c.Name.Name
	f.Name = &Identifier{
		Name: name,
	}
	f.Body = &Block{}
	f.Class = c
	return f
}

func (c *Class) ResolveParents(p *Program) {
	for _, parent := range c.Parents {
		_, d := p.FindDeclaration(parent)
		if d == nil {
			p.Error(parent.Position, fmt.Sprintf("%s undefined", parent.Name))
		} else {
			switch t := d.(type) {
			case *Class:
				if c.Parent == nil {
					c.Parent = t
					//TO-DO check, cannot self inherit, cycle inherit
				} else {
					p.Error(parent.Position, "class can only inherit 1 other class")
				}

			case *Interface:
				c.Interfaces = append(c.Interfaces, t)

			default:
				p.Error(parent.Position, fmt.Sprintf("invalid parent type: %s", t.Identifier()))
			}
		}
	}
}

func (c *Class) HasMember(member string) bool {
	_, ok := c.VariableIndexes[member]
	if !ok {
		_, ok = c.FunctionIndexes[member]
	}
	return ok
}

func (c *Class) MemberType(member string) ir.Type {
	if index, ok := c.VariableIndexes[member]; ok {
		return c.AccumulatedVariables[index]
	} else if index, ok := c.FunctionIndexes[member]; ok {
		return c.AccumulatedFunctions[index].Type()
	}
	return nil
}

func (c *Class) GetMember(ctx *Context, this ir.Value, member string, useVTable bool) *ObjectInfo {
	result := &ObjectInfo{
		Parent: this,
	}

	if index, ok := c.VariableIndexes[member]; ok {
		class := CastToClass(ctx.Block, this, ir.NewPointerType(c.IRStruct))
		v := ir.NewGetElementPtr(c.IRStruct, class, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
		ctx.Block.AddInstruction(v)
		result.Object = v
		value := c.AccumulatedVariables[c.VariableIndexes[member]]
		if _, ok := value.(*ir.FuncType); ok {
			result.FunctionDefine = c.FindVariableFuncType(member)
		}
		return result

	} else if index, ok := c.FunctionIndexes[member]; ok {
		if useVTable {
			class := CastToClass(ctx.Block, this, ir.NewPointerType(c.IRStruct))
			vtable := ir.NewGetElementPtr(c.IRStruct, class, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			ctx.Block.AddInstruction(vtable)
			value := ctx.AutoLoad(vtable)
			f := ir.NewGetElementPtr(c.IRVTable, value, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
			ctx.Block.AddInstruction(f)
			result.Object = f
			result.IsMemberFunction = true
			result.FunctionDefine = c.AccumulatedFunctions[c.FunctionIndexes[member]]
			return result

		} else {
			result.Object = c.AccumulatedFunctions[c.FunctionIndexes[member]]
			result.IsMemberFunction = true
			result.FunctionDefine = c.AccumulatedFunctions[c.FunctionIndexes[member]]
			return result
		}
	}

	return nil
}

func (c *Class) GetClass(ctx *Context, counter ir.Value) ir.Value {
	counter = ctx.AutoLoad(counter)
	counterClass := ctx.Program.FindQualified(Counter).(*Class)
	object := counterClass.GetMember(ctx, counter, "object", false)
	return ctx.AutoLoad(object.Object)
}

func (c *Class) CreateInstance(ctx *Context, args *Arguments) ir.Value {
	f := c.IRFunctions[0]
	call := ir.NewCall(f)
	SetUserData(call, c.Qualified(ctx.Program.Module.Namespace))
	if args != nil {
		args.GenerateIR(ctx, call, c.AccumulatedFunctions[0])
	}
	ctx.Block.AddInstruction(call)
	return call
}

func (c *Class) DestroyInstance(b *ir.Block, instance ir.Value) ir.Value {
	f := c.IRFunctions[1]
	call := ir.NewCall(f, instance)
	b.AddInstruction(call)
	return call
}
