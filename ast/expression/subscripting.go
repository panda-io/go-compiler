package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Subscripting struct {
	Base
	Parent  Expression
	Element Expression
}

func (e *Subscripting) Type(c *node.Context) ir.Type {
	//TO-DO
	return nil
}

func (e *Subscripting) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
