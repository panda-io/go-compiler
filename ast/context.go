package ast

import (
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
	Program  *Program
	Function *Function

	Block      *ir.Block
	LeaveBlock *ir.Block
	LoopBlock  *ir.Block
	Returned   bool

	parent  *Context
	objects map[string]ir.Value
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		LeaveBlock: c.LeaveBlock,
		LoopBlock:  c.LoopBlock,

		parent:  c,
		objects: make(map[string]ir.Value),
	}
}

func (c *Context) ObjectType(name string) ir.Type {
	if v, ok := c.objects[name]; ok {
		return c.ContentType(v)
	} else if c.Function.Class != nil && c.Function.Class.HasMember(name) {
		return c.Function.Class.MemberType(name)
	} else if c.parent != nil {
		return c.parent.ObjectType(name)
	}
	return nil
}

func (c *Context) ContentType(value ir.Value) ir.Type {
	switch t := value.(type) {
	// global define
	case *ir.Global:
		return t.ContentType

	// global define
	case *ir.Func:
		return t.Sig

	// alloca in function
	case *ir.InstAlloca:
		return t.ElemType

	// class member
	case *ir.InstGetElementPtr:
		return t.Type().(*ir.PointerType).ElemType
	}
	return nil
}

func (c *Context) AutoLoad(value ir.Value) ir.Value {
	return AutoLoad(value, c.Block)
}

func (c *Context) AddObject(name string, value ir.Value) error {
	if _, ok := c.objects[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.objects[name] = value
	return nil
}

func (c *Context) FindObject(name string) ir.Value {
	if v, ok := c.objects[name]; ok {
		return v
	} else if c.Function.Class != nil && c.Function.Class.HasMember(name) {
		this := c.FindObject(ClassThis)
		v, _ := c.Function.Class.GetMember(c, this, name, true)
		return v
	} else if c.parent != nil {
		return c.parent.FindObject(name)
	}
	return nil
}

func (c *Context) FindSelector(selector string, member string) (parent ir.Value, value ir.Value, isMemberFunction bool) {
	parent = c.FindObject(selector)
	if parent == nil {
		_, d := c.Program.FindSelector(selector, member)
		if d == nil {
			// could be an enum
			_, e := c.Program.FindSelector("", selector)
			if enum, ok := e.(*Enum); ok {
				value = enum.GetMember(member)
			} else {
				return
			}
		}
		switch t := d.(type) {
		case *Enum:
			value = t.IRStructData

		case *Variable:
			value = t.IRVariable

		case *Function:
			value = t.IRFunction
		}

	} else if p, ok := parent.Type().(*ir.PointerType); ok {
		// find declaration
		if d, ok := c.Program.Declarations[p.UserData]; ok {
			if class, ok := d.(*Class); ok {
				if IsBuiltinClass(p.UserData) {
					value, isMemberFunction = class.GetMember(c, parent, member, false)
				} else {
					parent, value, isMemberFunction = class.GetMemberFromCounter(c, parent, member)
				}
			}
		}
	}
	return
}

func AutoLoad(value ir.Value, b *ir.Block) ir.Value {
	switch t := value.(type) {
	// global define
	case *ir.Global:
		load := ir.NewLoad(t.ContentType, t)
		CopyUserData(t, load)
		b.AddInstruction(load)
		return load

	// global define
	case *ir.Func:
		return t

	// alloca in function
	case *ir.InstAlloca:
		load := ir.NewLoad(t.ElemType, t)
		CopyUserData(t, load)
		b.AddInstruction(load)
		return load

	// class member
	case *ir.InstGetElementPtr:
		typ := t.Type().(*ir.PointerType)
		load := ir.NewLoad(typ.ElemType, t)
		CopyUserData(t, load)
		b.AddInstruction(load)
		return load

	// ref param
	case *ir.Param:
		if t.Ref {
			if ir.IsPointer(t.Typ) && t.Typ.(*ir.PointerType).UserData == "" {
				typ := t.Type().(*ir.PointerType)
				load := ir.NewLoad(typ.ElemType, t)
				b.AddInstruction(load)
				return load
			}
		}
	}

	return value
}
