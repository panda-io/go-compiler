package ast

import (
	"crypto/md5"
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

func NewContext(p *Program) *Context {
	return &Context{
		Program: p,
		objects: make(map[string]ir.Value),
	}
}

type Context struct {
	Program *Program
	Module  *Module

	Block    *ir.Block
	Class    *Class
	Function *Function

	parent  *Context
	objects map[string]ir.Value
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program:  c.Program,
		Module:   c.Module,
		Class:    c.Class,
		Function: c.Function,

		parent:  c,
		objects: make(map[string]ir.Value),
	}
}

func (c *Context) StructType(qualified string) *ir.StructType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return t
}

func (c *Context) StructPointer(qualified string) *ir.PointerType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return ir.NewPointerType(t)
}

func (c *Context) AddObject(name string, value ir.Value) error {
	if _, ok := c.objects[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.objects[name] = value
	return nil
}

func (c *Context) ObjectType(name string) ir.Type {
	if v, ok := c.objects[name]; ok {
		return v.Type()
	} else if c.Class != nil && c.Class.HasMember(name) {
		return c.Class.MemberType(name)
	} else if c.parent != nil {
		return c.parent.ObjectType(name)
	}
	return nil
}

func (c *Context) FindObject(name string) ir.Value {
	if v, ok := c.objects[name]; ok {
		return v
	} else if c.Class != nil && c.Class.HasMember(name) {
		return c.Class.GetMember(c, c.FindObject(ClassThis), name)
	} else if c.parent != nil {
		return c.parent.FindObject(name)
	}
	return nil
}

func (c *Context) FindSelector(parent string, member string) (parentValue ir.Value, memberValue ir.Value) {
	parentValue = c.FindObject(parent)
	if parentValue == nil {
		// find from imports
		for _, i := range c.Module.Imports {
			if i.Alias == parent {
				ctx := c.Program.Contexts[i.Namespace]
				if ctx == nil {
					c.Error(i.Position, "invalid import")
					return
				}
				memberValue = ctx.objects[member]
				return
			}
		}
	} /*else {
		// TO-DO parent is an object, find its member then
		//parent is "this", "base"
		//parent is an object //class instance
	}*/
	return
}

func (c *Context) AddString(value string) *ir.Global {
	bytes := []byte(value)
	bytes = append(bytes, 0)
	hash := fmt.Sprintf("%x", md5.Sum(bytes))
	if s, ok := c.Program.Strings[hash]; ok {
		return s
	}
	s := c.Program.Module.NewGlobalDef("string."+hash, ir.NewCharArray(bytes))
	s.Immutable = true
	c.Program.Strings[hash] = s
	return s
}

func (c *Context) FindDeclaration(t *TypeName) (string, Declaration) {
	if t.Selector == "" {
		// search current package
		if c.Module.Namespace != Global {
			qualified := c.Module.Namespace + "." + t.Name
			d := c.Program.Declarations[qualified]
			if d != nil {
				return qualified, d
			}
		}
		// search global
		qualified := Global + "." + t.Name
		return qualified, c.Program.Declarations[qualified]
	} else {
		// search imports
		for _, i := range c.Module.Imports {
			if i.Alias == t.Selector {
				qualified := i.Namespace + "." + t.Name
				return qualified, c.Program.Declarations[qualified]
			}
		}
		return "", nil
	}
}

func (c *Context) Error(offset int, message string) {
	c.Program.Errors = append(c.Program.Errors, &Error{
		Position: c.Module.File.Position(offset),
		Message:  message,
	})
}
