package ast

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Program struct {
	Modules      map[string]*Module
	Declarations map[string]declaration.Declaration
	Namespaces   map[string]bool

	context *node.Context
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) AddModule(file string, m *Module) {
	p.Modules[file] = m
	p.context.Imports = m.Imports
	p.context.Namespace = m.Namespace
	p.Namespaces[m.Namespace] = true

	// zero pass
	for _, member := range m.Members {
		qualified := member.Qualified(m.Namespace)
		if p.Declarations[qualified] != nil {
			p.context.Error(member.GetPosition(), fmt.Sprintf("%s redeclared", member.Identifier()))
			//TO-DO get redeclaration position
		}
		switch t := member.(type) {
		case *declaration.Enum:
			t.GenerateIR(p.context)

		case *declaration.Class:
			t.GenerateStructDeclaration(p.context)
		}
	}
}

// TO-DO rebuild (language engine)
func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)
	p.Declarations = make(map[string]declaration.Declaration)
	p.Namespaces = make(map[string]bool)
	p.context = node.NewContext(ir.NewModule())
}

func (p *Program) GenerateIR() string {
	// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated
	// first pass
	for _, m := range p.Modules {
		p.context.Imports = m.Imports
		p.context.Namespace = m.Namespace

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Function:
				err := p.context.AddDeclaration(t.Qualified(m.Namespace), t.GenerateDeclaration(m.Namespace))
				if err != nil {
					p.context.Error(t.Position, err.Error())
				}

			case *declaration.Interface:
				// TO-DO save it then check class later
				// Generate function declaration

			case *declaration.Class:
				t.ResolveParents(p.context, p.Declarations)
			}
		}
	}

	// second pass
	for _, m := range p.Modules {
		p.context.Imports = m.Imports
		p.context.Namespace = m.Namespace

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Variable:
				// resovle later after all class type registered

			case *declaration.Function:
				t.GenerateIR(p.context)

			case *declaration.Interface:
				// TO-DO save it then check class later

			case *declaration.Class:
				t.GenerateIR(p.context)
			}
		}
	}

	buf := &strings.Builder{}
	_, err := p.context.Program.Module.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
