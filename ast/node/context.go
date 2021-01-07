package node

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

const (
	Global      = "global"
	Entry       = "main"
	Extern      = "extern"
	Constructor = "new"
	Destructor  = "destroy"
)

type Import struct {
	Alias     string
	Namespace string
}

type Error struct {
	Position int
	Message  string
}

//TO-DO parent(inheritance) later
type Struct struct {
	Parent    *Struct
	Variables []ir.Type
	Indexes   map[string]ir.Type
}

/*

type VTable struct {
	Type  *ir.Global
	Data  *ir.Global
	Index map[string]int
}
}*/

func NewProgramData(module *ir.Module) *ProgramData {
	return &ProgramData{
		Module:       module,
		Structs:      make(map[string]*Struct),
		Declarations: make(map[string]ir.Value),
	}
}

type ProgramData struct {
	Module *ir.Module

	Structs      map[string]*Struct
	Declarations map[string]ir.Value

	Errors []*Error
}

func NewContext(module *ir.Module) *Context {
	return &Context{
		Program:   NewProgramData(module),
		parent:    nil,
		variables: make(map[string]ir.Value),
	}
}

type Context struct {
	Program *ProgramData

	Block     *ir.Block
	Imports   []*Import
	Namespace string

	parent    *Context
	variables map[string]ir.Value
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program: c.Program,

		Imports:   c.Imports,
		Namespace: c.Namespace,
		parent:    c,
		variables: make(map[string]ir.Value),
	}
}

func (c *Context) AddVariable(name string, value ir.Value) error {
	if _, ok := c.variables[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.variables[name] = value
	return nil
}

func (c *Context) FindVariable(name string) ir.Value {
	if v, ok := c.variables[name]; ok {
		return v
	} else if c.parent != nil {
		return c.parent.FindVariable(name)
	}
	return nil
}

func (c *Context) AddDeclaration(qualified string, value ir.Value) error {
	if c.Program.Declarations[qualified] != nil {
		return fmt.Errorf("redeclared function %s.", qualified)
	}
	c.Program.Declarations[qualified] = value
	return nil
}

func (c *Context) FindDelaration(name string) []ir.Value {
	declarations := []ir.Value{}
	// search global
	qualified := Global + "." + name
	if c.Program.Declarations[qualified] != nil {
		declarations = append(declarations, c.Program.Declarations[qualified])
	}
	// search current package
	if c.Namespace != Global {
		qualified = c.Namespace + "." + name
		if c.Program.Declarations[qualified] != nil {
			declarations = append(declarations, c.Program.Declarations[qualified])
		}
	}
	// search import packages
	for _, i := range c.Imports {
		qualified = i.Namespace + "." + name
		if c.Program.Declarations[qualified] != nil {
			declarations = append(declarations, c.Program.Declarations[qualified])
		}
	}
	return declarations
}

func (c *Context) Errors() []*Error {
	return c.Program.Errors
}

func (c *Context) Error(p int, message string) {
	if c.parent != nil {
		c.parent.Error(p, message)
	} else {
		c.Program.Errors = append(c.Program.Errors, &Error{
			Position: p,
			Message:  message,
		})
	}
}
