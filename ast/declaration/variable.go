package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
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

func (c *Variable) GenerateIR(*node.Context) {
}

func (c *Variable) GenerateIRDeclaration(namespace string) ir.Value {
	return nil
}
