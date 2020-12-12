package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

type Import struct {
	Alias     string
	Namespace string
}

type Source struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Imports    []*Import
	Members    []declaration.Declaration
}
