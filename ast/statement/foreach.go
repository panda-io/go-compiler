package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
)

type Foreach struct {
	Base
	Key      Statement
	Item     Statement
	Iterator expression.Expression
	Body     Statement
}
