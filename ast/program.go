package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

const (
	Global = "global"
)

type Package struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Members    []declaration.Declaration
}

type Program struct {
	Packages map[string]*Package
}

func NewProgram() *Program {
	return &Program{
		Packages: make(map[string]*Package),
	}
}

func (p *Program) AddSource(s *Source) {
	if s.Namespace == "" {
		s.Namespace = Global
	}
	if _, ok := p.Packages[s.Namespace]; !ok {
		p.Packages[s.Namespace] = &Package{
			Namespace: s.Namespace,
		}
	}
	pkg := p.Packages[s.Namespace]
	pkg.Attributes = append(pkg.Attributes, s.Attributes...)
	pkg.Members = append(pkg.Members, s.Members...)
}
