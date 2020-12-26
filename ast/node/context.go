package node

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

const (
	Global = "global"
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

	imports   []*Import
	namespace string
	parent    *Context
	variables map[string]ir.Value
	errors    []*Error
}

func (c *Context) NewContext() *Context {
	ctx := NewContext(c.Declarations, c.Module)
	ctx.imports = c.imports
	ctx.namespace = c.namespace
	ctx.parent = c
	return ctx
}

func (c Context) SetImports(namespace string, imports []*Import) {
	c.namespace = namespace
	c.imports = imports
}

func (c Context) AddVariable(name string, value ir.Value) error {
	if _, ok := c.variables[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.variables[name] = value
	return nil
}

func (c Context) FindVariable(name string) ir.Value {
	if v, ok := c.variables[name]; ok {
		return v
	} else if c.parent != nil {
		return c.parent.FindVariable(name)
	}
	return nil
}

func (c Context) FindDelaration(name string) []ir.Value {
	declarations := []ir.Value{}
	// search global
	if c.Declarations[name] != nil {
		declarations = append(declarations, c.Declarations[name])
	}
	// search current package
	if c.namespace != Global {
		qualified := c.namespace + "." + name
		if c.Declarations[qualified] != nil {
			declarations = append(declarations, c.Declarations[qualified])
		}
	}
	// search import packages
	for _, i := range c.imports {
		qualified := i.Namespace + "." + name
		if c.Declarations[qualified] != nil {
			declarations = append(declarations, c.Declarations[qualified])
		}
	}
	return declarations
}

func (c Context) Errors() []*Error {
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
