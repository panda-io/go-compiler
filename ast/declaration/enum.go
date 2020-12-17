package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
)

type Enum struct {
	Base
	Name    *expression.Identifier
	Members []Declaration
}

func (e *Enum) Identifier() string {
	return e.Name.Name
}
