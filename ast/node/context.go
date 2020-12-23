package node

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type IRObjectType int

const (
	IRVariable IRObjectType = iota
	IRStruct
	IRFunction
)

type IRObject struct {
	Type  IRObjectType
	Value ir.Value
}

type Error struct {
	Position int
	Message  string
}

func NewContext(declarations map[string]*IRObject, module *ir.Module) *Context {
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
	Declarations map[string]*IRObject

	parent    *Context
	variables map[string]ir.Value
	errors    []*Error
}

func (c *Context) NewContext() *Context {
	ctx := NewContext(c.Declarations, c.Module)
	ctx.parent = c
	return ctx
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
