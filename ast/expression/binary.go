package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Binary struct {
	Base
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (e *Binary) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Binary) IsConstant() bool {
	switch e.Operator {
	case token.Assign, token.MulAssign, token.DivAssign, token.ModAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		return false

	default:
		return e.Left.IsConstant() && e.Right.IsConstant()
	}
}
