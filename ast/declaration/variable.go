package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

type Variable struct {
	Base
	Token token.Token
	Name  *expression.Identifier
	Type  types.Type
	Value expression.Expression
}

func (v *Variable) Identifier() string {
	return v.Name.Name
}
