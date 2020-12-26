package ast

import (
	"io"

	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Module struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Imports    []*node.Import
	Members    []declaration.Declaration
}

func (m *Module) GenerateIR(c *node.Context, w io.Writer) {
	for _, member := range m.Members {
		member.GenerateIR(c)
	}
	_, err := c.Module.WriteTo(w)
	if err != nil {
		panic(err)
	}
}
