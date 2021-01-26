package ast

import (
	"fmt"

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
	switch l.Typ {
	case token.STRING:
		return c.Program.AddString(l.Value[1 : len(l.Value)-1])

	case token.CHAR:
		//TO-DO convert char to i32
		return nil

	case token.FLOAT:
		return ir.NewFloatFromString(ir.Float32, l.Value)

	case token.INT:
		return ir.NewIntFromString(ir.I32, l.Value)

	case token.BOOL:
		var b int64 = 0
		if l.Value == "true" {
			b = 1
		}
		return ir.NewInt(ir.I1, b)

	case token.NULL:
		return ir.NewNull(nil)

	default:
		return nil
	}
}

func (*Literal) IsConstant(p *Program) bool {
	return true
}

func (l *Literal) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	switch l.Typ {
	case token.STRING:
		//TO-DO check expected
		return p.AddString(l.Value[1 : len(l.Value)-1])

	case token.CHAR:
		//TO-DO convert char to i32
		return nil

	case token.FLOAT:
		if expected == nil {
			return ir.NewFloatFromString(ir.Float32, l.Value)
		} else if ir.IsFloat(expected) {
			if expected.(*ir.FloatType).Kind == ir.FloatKindFloat {
				return ir.NewFloatFromString(ir.Float32, l.Value)
			}
			return ir.NewFloatFromString(ir.Float64, l.Value)
		}
		p.Error(l.Position, fmt.Sprintf("cannot convert float to %s", expected.String()))
		return nil

	case token.INT:
		if expected == nil {
			return ir.NewIntFromString(ir.I32, l.Value)
		} else if ir.IsInt(expected) {
			i := expected.(*ir.IntType)
			t := ir.NewIntFromString(i, l.Value)
			t.Typ.Unsigned = i.Unsigned
			return t
		}
		p.Error(l.Position, fmt.Sprintf("cannot convert int to %s", expected.String()))
		return nil

	case token.BOOL:
		if expected != nil && !ir.IsBool(expected) {
			p.Error(l.Position, fmt.Sprintf("cannot convert bool type to %s", expected.String()))
			return nil
		}
		return ir.NewIntFromString(ir.I1, l.Value)

	case token.NULL:
		if expected == nil || !ir.IsPointer(expected) {
			p.Error(l.Position, "type is missing or invalid for null")
			return nil
		}
		return ir.NewNull(expected.(*ir.PointerType))

	default:
		return nil
	}
}
