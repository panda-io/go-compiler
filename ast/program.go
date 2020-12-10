package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

type Package struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Members    []declaration.Declaration
}

type Program struct {
	Global   *Package
	Packages map[string]*Package
}

func NewProgram() *Program {
	return &Program{
		Packages: make(map[string]*Package),
	}
}

func (p *Program) AddSource(s *Source) {
	if s.Namespace == "" {
		if p.Global == nil {
			p.Global = &Package{}
		}
		p.Global.Attributes = append(p.Global.Attributes, s.Attributes...)
		p.Global.Members = append(p.Global.Members, s.Members...)
	} else {
		if _, ok := p.Packages[s.Namespace]; !ok {
			p.Packages[s.Namespace] = &Package{
				Namespace: s.Namespace,
			}
		}
		pkg := p.Packages[s.Namespace]
		pkg.Attributes = append(pkg.Attributes, s.Attributes...)
		pkg.Members = append(pkg.Members, s.Members...)
	}
}
