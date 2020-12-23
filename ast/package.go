package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Package struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Members    []declaration.Declaration
}

func (p *Package) GenerateIR(c *node.Context) string {
	for _, member := range p.Members {
		member.GenerateIR(c)
	}
	return c.Module.String()
}
