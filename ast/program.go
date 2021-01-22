package ast

import (
	"fmt"
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

func (p *Program) AddModule(file string, m *Module) {
	p.Modules[file] = m
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
	// zero pass (register all)
	for _, m := range p.Modules {
		// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated
		if c, ok := p.Contexts[m.Namespace]; ok {
			p.Context = c
		} else {
			p.Context = p.Contexts[Global].NewContext()
			p.Contexts[m.Namespace] = p.Context
		}
		p.Context.Module = m

		for _, member := range m.Members {
			qualified := member.Qualified(m.Namespace)
			if p.Declarations[qualified] == nil {
				p.Declarations[qualified] = member
			} else {
				p.Context.Error(member.GetPosition(), fmt.Sprintf("%s redeclared", member.Identifier()))
			}

			switch t := member.(type) {
			case *Enum:
				t.GenerateIR(p.Context)

			case *Class:
				t.PreProcess(p.Context)
			}
		}
	}

	// first pass (resolve oop)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, member := range m.Members {
			switch t := member.(type) {
			case *Interface:
				// TO-DO save it then check class later
				// Generate function declaration
				// resolve parants

			case *Class:
				t.ResolveParents(p.Context)
				t.IRStruct.GenerateDeclaration(p.Context)
				t.IRVTable.GenerateDeclaration(p.Context)
			}
		}
	}

	// second pass (generate declarations)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, member := range m.Members {
			switch t := member.(type) {
			case *Variable:
				// TO-DO

			case *Function:
				t.GenerateDeclaration(p.Context)

			case *Interface:
				// TO-DO save it then check class later
				// Generate function declaration

			case *Class:
				t.IRStruct.GenerateIR(p.Context)
				t.IRVTable.GenerateIR(p.Context)
			}
		}
	}

	// third pass (generate functions)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, member := range m.Members {
			switch t := member.(type) {
			case *Function:
				t.GenerateIR(p.Context)

			case *Interface:
				// TO-DO save it then check class later

			case *Class:
				t.GenerateIR(p.Context)
			}
		}
	}

	buf := &strings.Builder{}
	_, err := p.Module.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
