package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

type Using struct {
	Alias     string
	Namespace string
}

type Source struct {
	Namespace  string
	Attributes []*declaration.Attribute
	Using      []*Using
	Members    []declaration.Declaration
}
