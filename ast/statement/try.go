package statement

import (
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Try struct {
	Base
	Try     Statement
	Operand *types.Parameters
	Catch   Statement
	Finally Statement
}
