package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Module struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Imports    []*node.Import
	Members    []declaration.Declaration
}

func (m *Module) GenerateIR(c *node.Context) {
	for _, member := range m.Members {
		member.GenerateIR(c)
	}
}
