package ast

import (
	"github.com/panda-foundation/go-compiler/token"
)

// Expression interface for all expressions
type Expression interface {
	expression()
}

// Literal is basic type
type Literal struct {
	Node
	Type  token.Token // (identifier, string, char, float, int)
	Value string
}
