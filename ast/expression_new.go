package ast

import (
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type New struct {
	ExpressionBase
	Typ       types.Type
	Arguments *Arguments
}

func (n *New) Type(c *Context) ir.Type {
	//TO-DO
	return nil
}

func (n *New) GenerateIR(c *Context) ir.Value {
	//TO-DO
	return nil
}
