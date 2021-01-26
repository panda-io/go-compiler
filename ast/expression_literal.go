package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Literal struct {
	ExpressionBase
	Typ   token.Token // (identifier, string, char, float, int, bool)
	Value string
}

func (l *Literal) Type(c *Context) ir.Type {
	switch l.Typ {
	case token.STRING:
		return ir.NewArrayType(uint64(len(l.Value)-1), ir.I8)

	case token.CHAR:
		return ir.I32

	case token.FLOAT:
		return ir.Float32

	case token.INT:
		return ir.I32

	case token.BOOL:
		return ir.I1

	case token.NULL:
		return ir.Void

	default:
		return nil
	}
}

func (l *Literal) GenerateIR(c *Context) ir.Value {
	// case token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL, token.NULL, token.Void:
	switch l.Typ {
	case token.STRING:
		return c.AddString(l.Value[1 : len(l.Value)-1])

	case token.CHAR:
		//TO-DO convert char to i32
		return nil

	case token.FLOAT:
		return ir.NewFloatFromString(ir.Float32, l.Value)

	case token.INT:
		return ir.NewIntFromString(ir.I32, l.Value)

	case token.BOOL:
		return ir.NewIntFromString(ir.I1, l.Value)

	case token.NULL:
		return ir.NewNull(nil)

	default:
		return nil
	}
}

func (*Literal) IsConstant() bool {
	return true
}

func (*Literal) GenerateConstIR(c Context, exprect ir.Type) ir.Value {
	//TO-DO
	return nil
}
