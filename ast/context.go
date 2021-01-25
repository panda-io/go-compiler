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
	Program  *Program
	Function *Function

	Block      *ir.Block
	LeaveBlock *ir.Block
	Terminated bool

	parent  *Context
	objects map[string]ir.Value
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		parent:  c,
		objects: make(map[string]ir.Value),
	}
}

func (c *Context) ObjectType(name string) ir.Type {
	if v, ok := c.objects[name]; ok {
		return v.Type()
	} else if c.Function.Class != nil && c.Function.Class.HasMember(name) {
		return c.Function.Class.MemberType(name)
	} else if c.parent != nil {
		return c.parent.ObjectType(name)
	}
	return nil
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
		return c.Function.Class.GetMember(c, c.FindObject(ClassThis), name)
	} else if c.parent != nil {
		return c.parent.FindObject(name)
	}
	return nil
}

func (c *Context) FindSelector(selector string, member string) (parent ir.Value, value ir.Value) {
	parent = c.FindObject(selector)
	if parent == nil {
		_, d := c.Program.FindSelector(selector, member)
		if d == nil {
			return
		}
		// TO-DO can be enum, function, variable
		// cannot be interface, class, they has no static member
		switch t := d.(type) {
		case *Enum:
			//TO-DO
			return nil, nil

		case *Variable:
			//TO-DO
			return nil, nil

		case *Function:
			return nil, t.IRFunction

		default:
			return nil, nil
		}
	} /*else {
		// TO-DO parent is class or interface, find its member then
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
	s := c.Program.IRModule.NewGlobalDef("string."+hash, ir.NewCharArray(bytes))
	s.Immutable = true
	c.Program.Strings[hash] = s
	return s
}
