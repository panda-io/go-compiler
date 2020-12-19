package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/token"
)

type Binary struct {
	Base
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (e *Binary) GenerateIR(c *node.Context) {
	//TO-DO
}
