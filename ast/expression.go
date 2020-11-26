package ast

import (
	"github.com/panda-foundation/go-compiler/token"
)

// Expression interface for all expressions
type Expression interface {
	expression()
}

func (*Literal) expression()    {}
func (*Identifier) expression() {}

// Literal is basic type
type Literal struct {
	Position int
	Type     token.Token // (identifier, string, char, float, int)
	Value    string
}

// Identifier expr
type Identifier struct {
	Position int
	Name     string
}
