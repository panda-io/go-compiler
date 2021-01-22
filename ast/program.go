package ast

import (
	"strings"

	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Error struct {
	Position *token.Position
	Message  string
}

type Program struct {
	Modules map[string]*Module

	Module   *ir.Module
	Context  *Context
	Contexts map[string]*Context

	Declarations map[string]Declaration
	Strings      map[string]*ir.Global

	Errors []*Error
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)

	p.Module = ir.NewModule()
	p.Context = nil
	p.Contexts = make(map[string]*Context)
	p.Contexts[Global] = NewContext(p)

	p.Declarations = make(map[string]Declaration)
	p.Strings = make(map[string]*ir.Global)

	p.Errors = p.Errors[:0]
}

func (p *Program) GenerateIR() string {
	// zero pass (generate declarations)
	for _, m := range p.Modules {
		// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated
		if c, ok := p.Contexts[m.Namespace]; ok {
			p.Context = c
		} else {
			p.Context = p.Contexts[Global].NewContext()
			p.Contexts[m.Namespace] = p.Context
		}
		p.Context.Module = m

		for _, f := range m.Functions {
			f.GenerateIRDeclaration(p.Context)
		}

		for _, e := range m.Enums {
			e.GenerateIR(p.Context)
		}

		for _, i := range m.Interfaces {
			i.ResolveParents(p.Context)
		}

		for _, c := range m.Classes {
			c.ResolveParents(p.Context)
			c.PreProcess(p.Context)
			c.GenerateIRDeclaration(p.Context)
		}
	}

	// first pass (resolve oop)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, c := range m.Classes {
			c.GenerateIRStruct(p.Context)
			c.GenerateIRVTable(p.Context)
		}
	}

	// second pass (generate functions)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, f := range m.Functions {
			f.GenerateIR(p.Context)
		}

		for _, c := range m.Classes {
			c.GenerateIR(p.Context)
		}
	}

	buf := &strings.Builder{}
	_, err := p.Module.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
