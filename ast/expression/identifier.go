package expression

import "github.com/panda-foundation/go-compiler/ast/node"

type Identifier struct {
	Base
	Name string
}

func (e *Identifier) GenerateIR(c *node.Context) {
	//TO-DO
}
