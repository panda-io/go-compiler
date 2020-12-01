package ast

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

func NewPackage(packageName string, parent *Package) *Package {
	return &Package{
		Name: packageName,

		Variables:  make(map[string]*declaration.Variable),
		Functions:  make(map[string]*declaration.Function),
		Enums:      make(map[string]*declaration.Enum),
		Interfaces: make(map[string]*declaration.Interface),
		Classes:    make(map[string]*declaration.Class),

		Parent:   parent,
		Children: make(map[string]*Package),
	}
}

type Package struct {
	declaration.Attributes

	Name string

	Variables  map[string]*declaration.Variable
	Functions  map[string]*declaration.Function
	Enums      map[string]*declaration.Enum
	Interfaces map[string]*declaration.Interface
	Classes    map[string]*declaration.Class

	Parent   *Package
	Children map[string]*Package
}
