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
	QualifiedName []string
	TypeArguments []Type
}
