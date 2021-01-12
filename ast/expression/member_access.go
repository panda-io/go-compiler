package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type MemberAccess struct {
	Base
	Parent        Expression
	Member        *Identifier
	FullNamespace string
}

func (e *MemberAccess) Type(c *node.Context) ir.Type {
	//TO-DO
	return nil
}

func (e *MemberAccess) GenerateIR(c *node.Context) ir.Value {
	//TO-DO
	return nil
}
