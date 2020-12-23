package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Foreach struct {
	Base
	Key      Statement
	Item     Statement
	Iterator expression.Expression
	Body     Statement
}

func (*Foreach) GenerateIR(*node.Context) {
	//TO-DO
}
