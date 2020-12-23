package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Enum struct {
	Base
	Members []Declaration
}

func (c *Enum) GenerateIR(*node.Context) {
}
