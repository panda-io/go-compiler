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
	ResolvedParent     *Class
	ResolvedInterfaces []*Interface
}

//TO-DO parent(inheritance) later
type Struct struct {
	Parent    *Struct
	Variables []*Variable

	Types   []ir.Type
	Indexes map[string]int
}

func (s *Struct) GenerateIR(ctx *node.Context) {
	// TO-DO add vtable pointer
	structs := []*Struct{s}
	current := s
	for current.Parent != nil {
		structs = append(structs, current.Parent)
		current = current.Parent
	}
	index := 0
	for i := len(structs) - 1; i > -1; i-- {
		current = structs[i]
		for _, v := range current.Variables {
			s.Types = append(s.Types, types.TypeOf(v.Type))
			if _, ok := s.Indexes[v.Name.Name]; ok {
				ctx.Error(v.Position, fmt.Sprintf("duplicate class member: %s", v.Name.Name))
			} else {
				s.Indexes[v.Name.Name] = index
			}
			index++
		}
	}
}

func (c *Class) GenerateIR(ctx *node.Context) {
	// generate struct
	c.Struct.GenerateIR(ctx)
	ctx.Program.Module.NewGlobal(c.Qualified(ctx.Namespace), ir.NewStructType(c.Struct.Types...))

	//TO-DO generate function

	//TO-DO generate vtable
}

func (c *Class) GenerateStructDeclaration(*node.Context) *Struct {
	// resolve parent later
	s := &Struct{
		Indexes: make(map[string]int),
	}
	for _, m := range c.Members {
		if v, ok := m.(*Variable); ok {
			s.Variables = append(s.Variables, v)
		}
	}
	c.Struct = s
	return s
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
					c.Struct.Parent = t.Struct
					//TO-DO check, cannot self inherit
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
