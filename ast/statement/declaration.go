package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Declaration struct {
	Base
	Name  *expression.Identifier
	Type  types.Type
	Value expression.Expression
}

func (*Declaration) GenerateIR(*node.Context) {
	//TO-DO
}
