package declaration

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Enum struct {
	Base
	Members []Declaration
}

func (c *Enum) GenerateIR(*ir.Module) {
}
