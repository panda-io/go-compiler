package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Literal struct {
	Base
	Type  token.Token // (identifier, string, char, float, int, bool)
	Value string
}

func (e *Literal) GenerateIR() ir.Value {
	switch e.Type {
	case token.STRING:
		//TO-DO

	case token.CHAR:
		//TO-DO

	case token.FLOAT:
		return ir.NewFloatFromString(ir.Float32, e.Value)

	case token.INT:
		return ir.NewIntFromString(ir.I32, e.Value)

	case token.BOOL:
		return ir.NewIntFromString(ir.I1, e.Value)

	case token.IDENT:
		//TO-DO
	}
	return nil
}

func (e *Literal) IsConstant() bool {
	switch e.Type {
	case token.IDENT:
		//TO-DO
		return false

	default:
		return true
	}
}
