package expression

import "github.com/panda-foundation/go-compiler/ast/node"

type Conditional struct {
	Base
	Condition Expression
	First     Expression
	Second    Expression
}

func (e *Conditional) GenerateIR(c *node.Context) {
	//TO-DO
}
