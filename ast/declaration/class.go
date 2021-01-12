package declaration

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Class struct {
	Base
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration

	Struct             *Struct
	VTable             *VTable
	ResolvedParent     *Class
	ResolvedInterfaces []*Interface
}

type Struct struct {
	Class     *Class
	Variables []*Variable

	Type    *ir.StructType
	Members []ir.Type
	Indexes map[string]int
}

func (s *Struct) GenerateIR(ctx *node.Context) {
	structs := []*Struct{s}
	current := s
	for current.Class.ResolvedParent != nil {
		structs = append(structs, current.Class.ResolvedParent.Struct)
		current = current.Class.ResolvedParent.Struct
	}
	index := 0
	for i := len(structs) - 1; i > -1; i-- {
		current = structs[i]
		for _, v := range current.Variables {
			s.Members = append(s.Members, types.TypeOf(v.Type))
			if _, ok := s.Indexes[v.Name.Name]; ok {
				ctx.Error(v.Position, fmt.Sprintf("duplicate class member: %s", v.Name.Name))
			} else {
				s.Indexes[v.Name.Name] = index
			}
			index++
		}
	}
	// TO-DO add vtable pointer, add interface vtable
	qualified := s.Class.Qualified(ctx.Namespace)
	s.Type = ir.NewStructType(s.Members...)
	v := ctx.Program.Module.NewGlobal(qualified, s.Type)
	err := ctx.AddObject(s.Class.Name.Name, v)
	if err != nil {
		ctx.Error(s.Class.Position, fmt.Sprintf("%s redeclared", s.Class.Name.Name))
	}
}

type VTable struct {
	Class     *Class
	Functions []*Function

	//Table   []*ir.Func
	Indexes map[string]int
}

func (t *VTable) GenerateIR(ctx *node.Context) {
	for _, v := range t.Functions {
		v.GenerateDeclaration(ctx)
	}
	/*
		vtables := []*VTable{t}
		current := t
		for current.Parent != nil {
			vtables = append(vtables, current.Parent)
			current = current.Parent
		}
		index := 0
		for i := len(vtables) - 1; i > -1; i-- {
			current = vtables[i]
			for _, v := range current.Functions {
				v.GenerateDeclaration(ctx)
				if _, ok := t.Indexes[v.Name.Name]; ok {
					ctx.Error(v.Position, fmt.Sprintf("duplicate class member: %s", v.Name.Name))
				} else {
					t.Indexes[v.Name.Name] = index
				}
				index++
			}
		}*/
	//TO-DO compare function sig
}

func (c *Class) GenerateDeclaration(ctx *node.Context) {
	c.Struct.GenerateIR(ctx)

	c.VTable.GenerateIR(ctx)

	//TO-DO generate vtable
}

func (c *Class) GenerateIR(ctx *node.Context) {
	for _, v := range c.VTable.Functions {
		v.GenerateIR(ctx)
	}
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
	c.Struct = s

	t := &VTable{
		Class:   c,
		Indexes: make(map[string]int),
	}
	for _, m := range c.Members {
		if v, ok := m.(*Function); ok {
			t.Functions = append(t.Functions, v)
		}
	}
	c.VTable = t
}

func (c *Class) ResolveParents(ctx *node.Context, declarations map[string]Declaration) {
	for _, p := range c.Parents {
		d := FindeDeclaration(ctx, declarations, p)
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
