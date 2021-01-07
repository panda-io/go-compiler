package ast

import (
	"strings"

	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Program struct {
	Modules map[string]*Module

	//interfaces map[string]string

	context *node.Context
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) AddModule(file string, m *Module) {
	//TO-DO check redeclare and merge
	p.Modules[file] = m
	p.context.Imports = m.Imports
	p.context.Namespace = m.Namespace
	for _, member := range m.Members {
		switch t := member.(type) {
		case *declaration.Variable:
			// resovle later after all class type registered

		case *declaration.Function:
			// resovle later after all class type registered

		case *declaration.Enum:
			t.GenerateIR(p.context)

		case *declaration.Interface:
			// TO-DO save it then check class later

		case *declaration.Class:
			p.context.Program.Structs[t.Qualified(m.Namespace)] = t.GenerateStructDeclaration(p.context)
		}
	}
}

// TO-DO rebuild (language engine)
func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)
	//interfaces: make(map[string]string),
	p.context = node.NewContext(ir.NewModule())
}

func (p *Program) GenerateIR() string {
	// first pass
	for _, m := range p.Modules {
		p.context.Imports = m.Imports
		p.context.Namespace = m.Namespace

		for _, member := range m.Members {
			switch t := member.(type) {
			case *declaration.Variable:
				// resovle later after all class type registered

			case *declaration.Function:
				err := p.context.AddDeclaration(t.Qualified(m.Namespace), t.GenerateDeclaration(m.Namespace))
				if err != nil {
					p.context.Error(t.Position, err.Error())
				}

			case *declaration.Interface:
				// TO-DO save it then check class later

			case *declaration.Class:
				// TO-DO resolve class inheritance (struct, vtable)
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
