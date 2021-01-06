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

func NewContext(declarations map[string]ir.Value, module *ir.Module) *Context {
	return &Context{
		Declarations: declarations,
		Module:       module,
		parent:       nil,
		variables:    make(map[string]ir.Value),
	}
}

type Context struct {
	Block        *ir.Block
	Module       *ir.Module
	Declarations map[string]ir.Value

	Imports   []*Import
	Namespace string

	parent    *Context
	variables map[string]ir.Value
	errors    []*Error
}

func (c *Context) NewContext() *Context {
	ctx := NewContext(c.Declarations, c.Module)
	ctx.Imports = c.Imports
	ctx.Namespace = c.Namespace
	ctx.parent = c
	return ctx
}

func (c *Context) SetImports(namespace string, imports []*Import) {
	c.Namespace = namespace
	c.Imports = imports
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

func (c *Context) FindDelaration(name string) []ir.Value {
	declarations := []ir.Value{}
	// search global
	qualified := Global + "." + name
	if c.Declarations[qualified] != nil {
		declarations = append(declarations, c.Declarations[qualified])
	}
	// search current package
	if c.Namespace != Global {
		qualified = c.Namespace + "." + name
		if c.Declarations[qualified] != nil {
			declarations = append(declarations, c.Declarations[qualified])
		}
	}
	// search import packages
	for _, i := range c.Imports {
		qualified = i.Namespace + "." + name
		if c.Declarations[qualified] != nil {
			declarations = append(declarations, c.Declarations[qualified])
		}
	}
	return declarations
}

func (c *Context) Errors() []*Error {
	return c.errors
}

func (c *Context) Error(p int, message string) {
	if c.parent != nil {
		c.parent.Error(p, message)
	} else {
		c.errors = append(c.errors, &Error{
			Position: p,
			Message:  message,
		})
	}
}
