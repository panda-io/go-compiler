package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Interface struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration
}

func (i *Interface) Identifier() string {
	return i.Name.Name
}
