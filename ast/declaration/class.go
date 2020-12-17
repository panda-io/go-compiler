package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Class struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration
}

func (c *Class) Identifier() string {
	return c.Name.Name
}
