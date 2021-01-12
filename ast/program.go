package ast

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Program struct {
	Modules      map[string]*Module
	Data         *node.ProgramData
	Declarations map[string]declaration.Declaration
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) AddModule(file string, m *Module) {
	p.Modules[file] = m
}

// TO-DO rebuild (language server)
func (p *Program) Reset() {
	p.Data = node.NewProgramData()
	p.Modules = make(map[string]*Module)
	p.Declarations = make(map[string]declaration.Declaration)
}

func (p *Program) GenerateIR() string {
	// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated

	// zero pass (register all)
	for _, m := range p.Modules {
		if c, ok := p.Data.Contexts[m.Namespace]; ok {
			p.Data.Context = c
		} else {
			p.Data.Context = p.Data.Contexts[node.Global].NewContext()
			p.Data.Contexts[m.Namespace] = p.Data.Context
		}
		p.Data.Context.Imports = m.Imports
		p.Data.Context.Namespace = m.Namespace

		for _, member := range m.Members {
			qualified := member.Qualified(m.Namespace)
			if p.Declarations[qualified] == nil {
				p.Declarations[qualified] = member
			} else {
				p.Data.Context.Error(member.GetPosition(), fmt.Sprintf("%s redeclared", member.Identifier()))
				//TO-DO get another redeclaration position
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
		p.Data.Context = p.Data.Contexts[m.Namespace]
		p.Data.Context.Imports = m.Imports
		p.Data.Context.Namespace = m.Namespace

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Interface:
				// TO-DO save it then check class later
				// Generate function declaration
				// resolve parants

			case *declaration.Class:
				t.ResolveParents(p.Data.Context, p.Declarations)
			}
		}
	}

	// second pass (generate declarations)
	for _, m := range p.Modules {
		p.Data.Context = p.Data.Contexts[m.Namespace]
		p.Data.Context.Imports = m.Imports
		p.Data.Context.Namespace = m.Namespace

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Function:
				t.GenerateDeclaration(p.Data.Context)

			case *declaration.Interface:
				// TO-DO save it then check class later
				// Generate function declaration

			case *declaration.Class:
				t.GenerateDeclaration(p.Data.Context)
			}
		}
	}

	// third pass (generate functions)
	for _, m := range p.Modules {
		p.Data.Context = p.Data.Contexts[m.Namespace]
		p.Data.Context.Imports = m.Imports
		p.Data.Context.Namespace = m.Namespace

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Variable:
				// resovle later after all class type registered

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
	_, err := p.Data.Context.Program.Module.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (p *Program) Errors() []*node.Error {
	return p.Data.Context.Program.Errors
}
