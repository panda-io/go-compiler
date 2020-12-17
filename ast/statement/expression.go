package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
)

type Expression struct {
	Base
	Expression expression.Expression
}
