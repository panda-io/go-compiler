package declaration

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Class struct {
	Base
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration

	ResolvedParent     *Class
	ResolvedInterfaces []*Interface

	IRStruct *Struct
	IRVTable *VTable
}

type Struct struct {
	Class     *Class
	Variables []*Variable
	Members   []ir.Type

	Type          *ir.StructType
	MergedMembers []ir.Type
	Indexes       map[string]int
}

func (s *Struct) GenerateDeclaration(ctx *node.Context, declarations map[string]Declaration) {
	for _, v := range s.Variables {
		s.Members = append(s.Members, TypeOf(ctx, declarations, v.Type))
	}
}

func (s *Struct) GenerateIR(ctx *node.Context) {
	t := ir.NewStructType()
	t.TypeName = ctx.Namespace + "." + s.Class.Name.Name + ".vtable.type"
	s.MergedMembers = append(s.MergedMembers, ir.NewPointerType(t))

	structs := []*Struct{s}
	current := s
	for current.Class.ResolvedParent != nil {
		structs = append(structs, current.Class.ResolvedParent.IRStruct)
		current = current.Class.ResolvedParent.IRStruct
	}
	index := 0
	for i := len(structs) - 1; i > -1; i-- {
		current = structs[i]
		for i, v := range current.Variables {
			s.MergedMembers = append(s.MergedMembers, current.Members[i])
			if _, ok := s.Indexes[v.Name.Name]; ok {
				ctx.Error(v.Position, fmt.Sprintf("duplicate class member: %s", v.Name.Name))
			} else {
				s.Indexes[v.Name.Name] = index
			}
			index++
		}
	}

	qualified := s.Class.Qualified(ctx.Namespace)
	s.Type = ir.NewStructType(s.MergedMembers...)
	v := ctx.Program.Module.NewGlobal(qualified, s.Type)
	err := ctx.AddObject(s.Class.Name.Name, v)
	if err != nil {
		ctx.Error(s.Class.Position, fmt.Sprintf("%s redeclared", s.Class.Name.Name))
	}
}

type VTable struct {
	Class     *Class
	Functions []*Function
	Members   []*ir.Func

	Type            *ir.StructType
	Data            *ir.Struct
	MergedFunctions []*ir.Func
	Indexes         map[string]int
}

//%Foo_vtable_type = type { i32(%Foo*)* }
func (t *VTable) GenerateDeclaration(ctx *node.Context, declarations map[string]Declaration) {
	for _, v := range t.Functions {
		t.Members = append(t.Members, v.GenerateDeclaration(ctx, declarations))
	}
}

func (t *VTable) GenerateIR(ctx *node.Context) {
	vtables := []*VTable{t}
	current := t
	for current.Class.ResolvedParent != nil {
		vtables = append(vtables, current.Class.ResolvedParent.IRVTable)
		current = current.Class.ResolvedParent.IRVTable
	}
	index := 0

	for i := len(vtables) - 1; i > -1; i-- {
		current = vtables[i]
		for i, v := range current.Functions {
			if existing, ok := t.Indexes[v.Name.Name]; ok {
				// existing function
				f := t.MergedFunctions[existing]
				if !CompareMemberFunction(f.Sig, current.Members[i].Sig) {
					ctx.Error(v.Position, fmt.Sprintf("member function %s does not match its parent class", v.Name.Name))
					//TO-DO print more params details here
				} else {
					t.MergedFunctions[existing] = current.Members[i]
				}
			} else {
				// new function
				t.MergedFunctions = append(t.MergedFunctions, current.Members[i])
				t.Indexes[v.Name.Name] = index
				index++
			}
		}
	}

	var types []ir.Type
	var constants []ir.Constant
	for _, f := range t.MergedFunctions {
		types = append(types, f.Sig)
		constants = append(constants, f)
	}
	t.Type = ir.NewStructType(types...)
	v := ctx.Program.Module.NewGlobal(t.Class.Qualified(ctx.Namespace)+".vtable.type", t.Type)
	err := ctx.AddObject(t.Class.Name.Name+".vtable.type", v)
	if err != nil {
		ctx.Error(t.Class.Position, fmt.Sprintf("%s redeclared", t.Class.Name.Name))
	}

	vtableType := ir.NewStructType()
	vtableType.TypeName = t.Class.Qualified(ctx.Namespace) + ".vtable.type"
	t.Data = ir.NewStruct(vtableType, constants...)
	v = ctx.Program.Module.NewGlobalDef(t.Class.Qualified(ctx.Namespace)+".vtable.data", t.Data)
	err = ctx.AddObject(t.Class.Name.Name+".vtable.data", v)
	if err != nil {
		ctx.Error(t.Class.Position, fmt.Sprintf("%s redeclared", t.Class.Name.Name))
	}
}

func (c *Class) GenerateIR(ctx *node.Context) {
	for _, v := range c.IRVTable.Functions {
		v.GenerateIR(ctx)
	}
	//TO-DO constructor and destructor
}

func (c *Class) PreProcess(*node.Context) {
	s := &Struct{
		Class:   c,
		Indexes: make(map[string]int),
	}
	for _, m := range c.Members {
		if v, ok := m.(*Variable); ok {
			s.Variables = append(s.Variables, v)
		}
	}
	c.IRStruct = s

	t := &VTable{
		Class:   c,
		Indexes: make(map[string]int),
	}

	// first is constructor, second is destructor
	t.Functions = append(t.Functions, nil, nil)
	for _, m := range c.Members {
		if v, ok := m.(*Function); ok {
			if v.Name.Name == node.Constructor {
				t.Functions[0] = v
			} else if v.Name.Name == node.Destructor {
				t.Functions[1] = v
			} else {
				t.Functions = append(t.Functions, v)
			}
		}
	}
	if t.Functions[0] == nil {
		t.Functions[0] = &Function{}
		t.Functions[0].ObjectName = c.Name.Name
		t.Functions[0].Name = &expression.Identifier{
			Name: node.Constructor,
		}
	}
	t.Functions[0].ReturnType = &types.TypeName{
		Name: c.Name.Name,
	}
	if t.Functions[1] == nil {
		t.Functions[1] = &Function{}
		t.Functions[1].ObjectName = c.Name.Name
		t.Functions[1].Name = &expression.Identifier{
			Name: node.Destructor,
		}
	}
	c.IRVTable = t
}

func (c *Class) ResolveParents(ctx *node.Context, declarations map[string]Declaration) {
	for _, p := range c.Parents {
		_, d := FindDeclaration(ctx, declarations, p)
		if d == nil {
			ctx.Error(p.Position, fmt.Sprintf("%s undefined", p.Name))
		} else {
			switch t := d.(type) {
			case *Class:
				if c.ResolvedParent == nil {
					c.ResolvedParent = t
					//TO-DO check, cannot self inherit, cycle inherit
				} else {
					ctx.Error(p.Position, "class can only inherit 1 other class")
				}
			case *Interface:
				c.ResolvedInterfaces = append(c.ResolvedInterfaces, t)
			default:
				ctx.Error(p.Position, fmt.Sprintf("invalid parent type: %s", t.Identifier()))
			}
		}
	}
}
