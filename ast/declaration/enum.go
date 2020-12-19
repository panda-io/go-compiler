package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ir"
)

type Enum struct {
	Base
	Name    *expression.Identifier
	Members []Declaration
}

func (c *Enum) GenerateIR(*ir.Module) {
}
