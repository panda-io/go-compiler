package ast

import "github.com/panda-foundation/go-compiler/token"

// Type struct
type Type interface {
	types()
}

func (*BuitinType) types() {}
func (*TypeName) types()   {}

// BuitinType basic type
type BuitinType struct {
	Position int
	Token    token.Token
}

// TypeName qualified name
type TypeName struct {
	Position      int
	QualifiedName []*Identifier
	TypeArguments *TypeArguments
}

// TypeArguments type arguments
type TypeArguments struct {
	Position  int
	Arguments []Type
}

// TypeParameters type parameters
type TypeParameters struct {
	Position   int
	Parameters []*TypeParameter
}

// TypeParameter type parameter
type TypeParameter struct {
	Name *Identifier
	Type Type
}
