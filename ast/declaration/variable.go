package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Variable struct {
	Base
	Token token.Token
	Type  types.Type
	Value expression.Expression
}

func (c *Variable) GenerateIR(*ir.Module) {
}
