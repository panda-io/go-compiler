package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type MemberAccess struct {
	Base
	Parent        Expression
	Member        *Identifier
	FullNamespace string
}

func (e *MemberAccess) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *MemberAccess) IsConstant() bool {
	//TO-DO search scope
	return false
}
