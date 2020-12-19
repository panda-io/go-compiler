package expression

import "github.com/panda-foundation/go-compiler/ast/node"

type Parentheses struct {
	Base
	Expression Expression
}

func (e *Parentheses) GenerateIR(c *node.Context) {
	//TO-DO
}
