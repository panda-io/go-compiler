package ast

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast/node"
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
		if c, ok := p.Data.Contexts[m.Namespace]; ok {
			p.Context = c
		} else {
			p.Context = p.Data.Contexts[node.Global].NewContext()
			p.Contexts[m.Namespace] = p.Data.Context
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
			case *declaration.Enum:
				t.GenerateIR(p.Data.Context)

			case *declaration.Class:
				t.PreProcess(p.Data.Context)
			}
		}
	}

	// first pass (resolve oop)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Interface:
				// TO-DO save it then check class later
				// Generate function declaration
				// resolve parants

			case *declaration.Class:
				t.ResolveParents(p.Data.Context, p.Declarations)
				t.IRStruct.GenerateDeclaration(p.Data.Context, p.Declarations)
				t.IRVTable.GenerateDeclaration(p.Data.Context, p.Declarations)
			}
		}
	}

	// second pass (generate declarations)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Variable:
				// TO-DO

			case *declaration.Function:
				t.GenerateDeclaration(p.Data.Context, p.Declarations)

			case *declaration.Interface:
				// TO-DO save it then check class later
				// Generate function declaration

			case *declaration.Class:
				t.IRStruct.GenerateIR(p.Data.Context)
				t.IRVTable.GenerateIR(p.Data.Context)
			}
		}
	}

	// third pass (generate functions)
	for _, m := range p.Modules {
		p.Context = p.Contexts[m.Namespace]
		p.Context.Module = m

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Function:
				t.GenerateIR(p.Data.Context)

			case *declaration.Interface:
				// TO-DO save it then check class later

			case *declaration.Class:
				t.GenerateIR(p.Data.Context)
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

func (p *Program) Errors() []*Error {
	return p.Errors
}
