package ast

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Program struct {
	Modules      map[string]*Module
	Declarations map[string]ir.Value
}

func NewProgram() *Program {
	return &Program{
		Modules:      make(map[string]*Module),
		Declarations: make(map[string]ir.Value),
	}
}

func (p *Program) Declare(name string, value ir.Value) {
	if p.Declarations[name] != nil {
		panic(fmt.Errorf("redeclared: %s", name))
	}
	p.Declarations[name] = value
}

func (p *Program) AddModule(file string, m *Module) {
	p.Modules[file] = m
	for _, member := range m.Members {
		name := m.Namespace + "." + member.Identifier()
		p.Declare(name, member.GenerateIRDeclaration(m.Namespace))
	}
}

func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)
	p.Declarations = make(map[string]ir.Value)
	//TO-DO build from source again
}

func (p *Program) GenerateIR() string {
	c := node.NewContext(p.Declarations, ir.NewModule())
	buf := &strings.Builder{}
	for _, m := range p.Modules {
		c.SetImports(m.Namespace, m.Imports)
		m.GenerateIR(c)
	}
	_, err := c.Module.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
