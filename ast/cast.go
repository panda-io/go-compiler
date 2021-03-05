package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

func PromoteNumberType(t1 ir.Type, t2 ir.Type) (ir.Type, error) {
	if ir.IsInt(t1) {
		if ir.IsInt(t2) {
			i1 := t1.(*ir.IntType)
			i2 := t2.(*ir.IntType)
			if i1.Unsigned != i2.Unsigned {
				return nil, fmt.Errorf("implicit conversion between signed and unsigned integers is not allowed")
			}
			if i1.BitSize > i2.BitSize {
				return i1, nil
			}
			return i2, nil
		} else if ir.IsFloat(t2) {
			return t2, nil
		}
	} else if ir.IsFloat(t1) {
		if ir.IsInt(t2) {
			return t1, nil
		} else if ir.IsFloat(t2) {
			f1 := t1.(*ir.FloatType)
			f2 := t2.(*ir.FloatType)
			if f1.Kind == ir.FloatKindDouble {
				return f1, nil
			}
			if f2.Kind == ir.FloatKindDouble {
				return f2, nil
			}
			return f1, nil
		}
	}
	return nil, fmt.Errorf("invalid number")
}

func CastToPointer(b *ir.Block, value ir.Value) ir.Value {
	if ir.IsPointer(value.Type()) {
		cast := ir.NewBitCast(value, pointerType)
		b.AddInstruction(cast)
		return cast
	}
	return nil
}

func CastFromPointer(b *ir.Block, value ir.Value, typ ir.Type) ir.Value {
	t := value.Type()
	if ir.IsPointer(t) && t.(*ir.PointerType).ElemType.Equal(ir.I8) {
		cast := ir.NewBitCast(value, typ)
		b.AddInstruction(cast)
		return cast
	}
	return nil
}

func CastNumberExpr(from ir.Constant, to ir.Type) ir.Constant {
	t := from.Type()
	if ir.IsInt(t) {
		i := t.(*ir.IntType)
		if ir.IsInt(to) {
			if i.Unsigned {
				return ir.NewExprZExt(from, to)
			}
			return ir.NewExprSExt(from, to)
		}
		if i.Unsigned {
			return ir.NewExprUIToFP(from, to)
		}
		return ir.NewExprSIToFP(from, to)
	}
	return ir.NewExprFPExt(from, to)
}

func CastNumber(c *Context, from ir.Value, to ir.Type) ir.Value {
	// will not check if valid, since PromoteNumberType is called
	// before this which already checked error
	var cast ir.Instruction
	t := from.Type()
	if ir.IsInt(t) {
		i := t.(*ir.IntType)
		if ir.IsInt(to) {
			if i.Unsigned {
				cast = ir.NewZExt(from, to)
			} else {
				cast = ir.NewSExt(from, to)
			}
		} else {
			if i.Unsigned {
				cast = ir.NewUIToFP(from, to)
			} else {
				cast = ir.NewSIToFP(from, to)
			}
		}
	} else {
		cast = ir.NewFPExt(from, to)
	}

	c.Block.AddInstruction(cast)
	return cast.(ir.Value)
}
