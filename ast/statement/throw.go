package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
)

type Throw struct {
	Base
	Expression expression.Expression
}
