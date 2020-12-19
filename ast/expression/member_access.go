package expression

import "github.com/panda-foundation/go-compiler/ast/node"

type MemberAccess struct {
	Base
	Parent        Expression
	Member        *Identifier
	FullNamespace string
}

func (e *MemberAccess) GenerateIR(c *node.Context) {
	//TO-DO
}
