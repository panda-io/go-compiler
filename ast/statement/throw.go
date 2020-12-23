package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
)

type Throw struct {
	Base
	Expression expression.Expression
}

func (*Throw) GenerateIR(*node.Context) {
	//TO-DO
}
