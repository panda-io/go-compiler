package expression

import "github.com/panda-foundation/go-compiler/ast/node"

type Subscripting struct {
	Base
	Parent  Expression
	Element Expression
}

func (e *Subscripting) GenerateIR(c *node.Context) {
	//TO-DO
}
