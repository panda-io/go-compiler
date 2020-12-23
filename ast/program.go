package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

const (
	Global = "global"
)

type Program struct {
	Packages     map[string]*Package
	Declarations map[string]*node.IRObject
}

func NewProgram() *Program {
	return &Program{
		Packages:     make(map[string]*Package),
		Declarations: make(map[string]*node.IRObject),
	}
}

func (p *Program) Declare(name string, typ node.IRObjectType, value ir.Value) error {
	if p.Declarations[name] != nil {
		return fmt.Errorf("redeclared: %s", name)
	}
	p.Declarations[name] = &node.IRObject{
		Type:  typ,
		Value: value,
	}
	return nil
}

func (p *Program) AddSource(s *Source) {
	if _, ok := p.Packages[s.Namespace]; !ok {
		p.Packages[s.Namespace] = &Package{
			Namespace: s.Namespace,
		}
	}
	pkg := p.Packages[s.Namespace]
	pkg.Attributes = append(pkg.Attributes, s.Attributes...)
	pkg.Members = append(pkg.Members, s.Members...)
}

func (p *Program) Reset() {
	p.Packages = make(map[string]*Package)
	p.Declarations = make(map[string]*node.IRObject)
	//TO-DO build from source again
}
