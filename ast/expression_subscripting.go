package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Subscripting struct {
	ExpressionBase
	Parent  Expression
	Element Expression
}

func (e *Subscripting) Type(c *Context) ir.Type {
	//TO-DO
	return nil
}

func (e *Subscripting) GenerateIR(c *Context) ir.Value {
	//TO-DO
	return nil
}