package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Enum struct {
	Base
	Members []Declaration
}

func (c *Enum) GenerateIR(*node.Context) {
}

func (c *Enum) GenerateIRDeclaration() ir.Value {
	return nil
}
