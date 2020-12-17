package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Function struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Compound
	ClassName      string
}

func (f *Function) Identifier() string {
	return f.Name.Name
}
