package statement

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Try struct {
	Base
	Try     Statement
	Operand *types.Parameters
	Catch   Statement
	Finally Statement
}

func (*Try) GenerateIR(*node.Context) {
	//TO-DO
}
