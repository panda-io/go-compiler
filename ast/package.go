package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

func NewPackage(packageName string, parent *Package) *Package {
	return &Package{
		Name:     packageName,
		Parent:   parent,
		Children: make(map[string]*Package),
	}
}

type Package struct {
	declaration.Attributes
	Name     string
	Members  []declaration.Declaration
	Parent   *Package
	Children map[string]*Package
}
