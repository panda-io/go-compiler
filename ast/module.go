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
