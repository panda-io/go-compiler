package node

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

const (
	Global      = "global"
	Entry       = "main"
	This        = "this"
	Constructor = "new"
	Destructor  = "destroy"

	Extern = "extern"
)

type Import struct {
	Alias     string
	Namespace string
}

type Error struct {
	Position int
	Message  string
}

func NewProgramData() *ProgramData {
	p := &ProgramData{
		Module:   ir.NewModule(),
		Contexts: make(map[string]*Context),
	}
	p.Contexts[Global] = NewContext(p)
	return p
}

type ProgramData struct {
	Module *ir.Module

	Context  *Context
	Contexts map[string]*Context

	Errors []*Error
}

func NewContext(data *ProgramData) *Context {
	return &Context{
		Program: data,
		objects: make(map[string]ir.Value),
	}
}

type Context struct {
	Program *ProgramData

	Block     *ir.Block
	Imports   []*Import
	Namespace string

	parent  *Context
	objects map[string]ir.Value
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program: c.Program,

		Imports:   c.Imports,
		Namespace: c.Namespace,
		parent:    c,
		objects:   make(map[string]ir.Value),
	}
}

func (c *Context) CloseContext() *Context {
	return c.parent
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
	} else if c.parent != nil {
		return c.parent.FindObject(name)
	}
	//TO-DO find from this
	return nil
}

func (c *Context) FindSelector(object string, member string) ir.Value {
	//This
	//Object//class instance, not define
	//Import
	return nil
}

/*
func (c *Context) AddDeclaration(qualified string, value ir.Value) error {
	if c.Program.Declarations[qualified] != nil {
		return fmt.Errorf("redeclared function %s.", qualified)
	}
	c.Program.Declarations[qualified] = value
	return nil
}

func (c *Context) FindDelaration(name string) ir.Value {
	// search current package
	if c.Namespace != Global {
		qualified := c.Namespace + "." + name
		if c.Program.Declarations[qualified] != nil {
			return c.Program.Declarations[qualified]
		}
	}
	// search global
	qualified := Global + "." + name
	if c.Program.Declarations[qualified] != nil {
		return c.Program.Declarations[qualified]
	}

	//TO-DO find with selector

		// search import packages
		for _, i := range c.Imports {
			qualified = i.Namespace + "." + name
			if c.Program.Declarations[qualified] != nil {
				declarations = append(declarations, c.Program.Declarations[qualified])
			}
		}
	return nil
}*/

func (c *Context) Errors() []*Error {
	return c.Program.Errors
}

func (c *Context) Error(p int, message string) {
	c.Program.Errors = append(c.Program.Errors, &Error{
		Position: p,
		Message:  message,
	})
}
