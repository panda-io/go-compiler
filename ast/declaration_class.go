package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type Class struct {
	DeclarationBase
	TypeParameters *TypeParameters
	Parents        []*TypeName
	Functions      []*Function
	Variables      []*Variable

	Parent     *Class
	Interfaces []*Interface

	IRStruct        *ir.StructType
	IRVariables     []ir.Type
	VariableIndexes map[string]int

	IRVTable        *ir.StructType
	IRFunctions     []*ir.Func
	IRVTableData    *ir.Global
	FunctionIndexes map[string]int
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
		c.IRVariables = append(c.IRVariables, v.Type.Type(p))
	}
	for _, f := range c.Functions {
		c.IRFunctions = append(c.IRFunctions, f.GenerateIRDeclaration(p))
	}
}

func (c *Class) GenerateIRStruct(p *Program) {
	c.VariableIndexes = make(map[string]int)

	variables := []ir.Type{CreateStructPointerType(p.Module.Namespace + "." + c.Name.Name + ".vtable.type")}
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
			variables = append(variables, current.IRVariables[j])
			if _, ok := c.VariableIndexes[v.Name.Name]; ok {
				p.Error(v.Position, fmt.Sprintf("duplicate class member: %s", v.Name.Name))
			} else {
				c.VariableIndexes[v.Name.Name] = index
			}
			index++
		}
	}

	qualified := c.Qualified(p.Module.Namespace)
	c.IRStruct = ir.NewStructType(variables...)
	p.IRModule.NewTypeDef(qualified, c.IRStruct)
}

func (c *Class) GenerateIRVTable(p *Program) {
	c.FunctionIndexes = make(map[string]int)

	functions := []*ir.Func{}
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
				function := functions[existing]
				if !c.CompareFunction(function, current.IRFunctions[j], f.Name.Name == Constructor) {
					p.Error(f.Position, fmt.Sprintf("member function %s does not match its parent class", f.Name.Name))
					//TO-DO print more params details here
				} else {
					functions[existing] = current.IRFunctions[j]
				}
			} else {
				// new function
				functions = append(functions, current.IRFunctions[j])
				c.FunctionIndexes[f.Name.Name] = index
				index++
			}
		}
	}

	var types []ir.Type
	var constants []ir.Constant
	for _, f := range functions {
		types = append(types, ir.NewPointerType(f.Sig))
		constants = append(constants, f)
	}
	c.IRVTable = ir.NewStructType(types...)
	p.IRModule.NewTypeDef(c.Qualified(p.Module.Namespace)+".vtable.type", c.IRVTable)

	data := ir.NewStruct(CreateStructType(c.Qualified(p.Module.Namespace)+".vtable.type"), constants...)
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
	c.Functions[0].ReturnType = &TypeName{
		Name: c.Name.Name,
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

func (c *Class) CompareFunction(f0 *ir.Func, f1 *ir.Func, isConstructor bool) bool {
	sig0 := f0.Sig
	sig1 := f1.Sig
	if isConstructor {
		return sig0.Equal(sig1)
	}
	if !sig0.RetType.Equal(sig1.RetType) {
		return false
	}
	if len(sig0.Params) != len(sig1.Params) {
		return false
	}
	for i := 1; i < len(sig0.Params); i++ {
		if !sig0.Params[i].Equal(sig1.Params[i]) {
			return false
		}
	}
	return sig0.Variadic == sig1.Variadic
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
		return ir.GepInstType(c.IRStruct, []ir.Value{ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index))})
	} else if index, ok := c.FunctionIndexes[member]; ok {
		return ir.GepInstType(c.IRVTable, []ir.Value{ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index))})
	}
	return nil
}

func (c *Class) GetMember(ctx *Context, this ir.Value, member string) ir.Value {
	if index, ok := c.VariableIndexes[member]; ok {
		v := ir.NewGetElementPtr(c.IRStruct, ctx.FindObject(ClassThis), ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
		ctx.Block.AddInstruction(v)
		return v
	} else if index, ok := c.FunctionIndexes[member]; ok {
		vtable := ir.NewGetElementPtr(c.IRStruct, ctx.FindObject(ClassThis), ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
		ctx.Block.AddInstruction(vtable)
		f := ir.NewGetElementPtr(c.IRVTable, vtable, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
		ctx.Block.AddInstruction(f)
		return f
	}
	return nil
}

//TO-DO set member
