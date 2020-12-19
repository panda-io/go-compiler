package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ir"
)

type Package struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Members    []declaration.Declaration
}

func (p *Package) GenerateIR() string {
	m := ir.NewModule()
	for _, member := range p.Members {
		member.GenerateIR(m)
	}
	return m.String()
}
