package ast

import (
	"github.com/panda-foundation/go-compiler/token"
)

// Expression interface for all expressions
type Expression interface {
	expression()
}

func (*BuitinType) expression()              {}
func (*Literal) expression()                 {}
func (*Identifier) expression()              {}
func (*ClassReference) expression()          {}
func (*ParenExpression) expression()         {}
func (*MemberAccessExpression) expression()  {}
func (*ElementAccessExpression) expression() {}
func (*InvocationExpression) expression()    {}

// Literal expr
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

// ClassReference this or base
type ClassReference struct {
	Position int
	Type     token.Token
}

// ParenExpression expr
type ParenExpression struct {
	Position   int
	Expression Expression
}

// MemberAccessExpression expr
type MemberAccessExpression struct {
	Parent Expression
	Member *Identifier
}

// ElementAccessExpression expr
type ElementAccessExpression struct {
	Parent  Expression
	Element Expression
}

// InvocationExpression expr
type InvocationExpression struct {
	Function  Expression
	Arguments *Arguments
}
