package ast

import (
	"github.com/panda-foundation/go-compiler/token"
)

type Variable struct {
	DeclarationBase
	Token token.Token
	Type  Type
	Value Expression
}

func (c *Variable) GenerateIR(*Context) {
}
