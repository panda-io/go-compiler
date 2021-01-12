package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type New struct {
	Base
	Typ       types.Type
	Arguments *Arguments
}

func (n *New) Type(c *node.Context) ir.Type {
	//TO-DO
	return nil
}

func (n *New) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
