package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
)

type Return struct {
	Base
	Expression expression.Expression
}
