package ast

import (
	"math/big"
	"strconv"

	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Literal struct {
	ExpressionBase
	Typ   token.Token
	Value string
}

func (l *Literal) Type(c *Context, expected ir.Type) ir.Type {
	switch l.Typ {
	case token.STRING:
		return ir.NewArrayType(uint64(len(l.Value)-1), ir.I8)

	case token.CHAR:
		return ir.I32

	case token.FLOAT:
		if expected != nil && ir.IsFloat(expected) {
			return expected
		}
		return ir.Float32

	case token.INT:
		if expected != nil && ir.IsInt(expected) {
			return expected
		}
		return ir.I32

	case token.BOOL:
		return ir.I1

	case token.NULL:
		if expected != nil && ir.IsPointer(expected) {
			return expected
		}
		return nil

	default:
		return nil
	}
}

func (l *Literal) GenerateIR(c *Context, expected ir.Type) ir.Value {
	switch l.Typ {
	case token.STRING:
		if l.Value[0] == '"' {
			// string
			str, _ := strconv.Unquote(l.Value)
			return c.Program.AddString(str)
		} else {
			// `` raw string
			return c.Program.AddString(l.Value[1 : len(l.Value)-1])
		}

	case token.CHAR:
		//TO-DO convert char to i32
		return nil

	case token.FLOAT:
		if expected != nil {
			if ir.IsFloat(expected) {
				return ir.NewFloatFromString(expected.(*ir.FloatType), l.Value)
			}
			c.Program.Error(l.Position, "type mismatch")
		}
		return ir.NewFloatFromString(ir.Float32, l.Value)

	case token.INT:
		if expected != nil {
			if ir.IsInt(expected) {
				return ir.NewIntFromString(expected.(*ir.IntType), l.Value)
			}
			c.Program.Error(l.Position, "type mismatch")
		}
		return ir.NewIntFromString(ir.I32, l.Value)

	case token.BOOL:
		if expected != nil && !ir.IsBool(expected) {
			c.Program.Error(l.Position, "type mismatch")
		}
		if l.Value == "true" {
			return ir.True
		}
		return ir.False

	case token.NULL:
		if expected != nil && ir.IsPointer(expected) {
			return ir.NewNull(expected.(*ir.PointerType))
		}
		c.Program.Error(l.Position, "missing type for null")
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
		p.Error(l.Position, "type mismatch")
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
		p.Error(l.Position, "type mismatch")
		return nil

	case token.BOOL:
		if expected != nil && !ir.IsBool(expected) {
			p.Error(l.Position, "type mismatch")
			return nil
		}
		if l.Value == "true" {
			return ir.True
		}
		return ir.False

	case token.NULL:
		if expected == nil || !ir.IsPointer(expected) {
			p.Error(l.Position, "missing type or invalid for null")
			return nil
		}
		return ir.NewNull(expected.(*ir.PointerType))

	default:
		return nil
	}
}

func (l *Literal) GetValue() interface{} {
	switch l.Typ {
	case token.STRING:
		return l.Value[1 : len(l.Value)-1]

	case token.CHAR:
		//TO-DO convert char to i32
		return nil

	case token.FLOAT:
		x, _, _ := big.ParseFloat(l.Value, 10, 24, big.ToNearestEven)
		f, _ := x.Float32()
		return f

	case token.INT:
		x, _ := (&big.Int{}).SetString(l.Value, 10)
		return x.Int64()

	case token.BOOL:
		if l.Value == "true" {
			return true
		}
		return false

	default:
		return nil
	}
}
