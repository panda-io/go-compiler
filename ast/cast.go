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

func CastToPointer(c *Context, value ir.Value) ir.Value {
	if ir.IsPointer(value.Type()) {
		cast := ir.NewBitCast(value, ir.NewPointerType(ir.I8))
		c.Block.AddInstruction(cast)
		return cast
	}
	return nil
}

func CastFromPointer(c *Context, value ir.Value, typ ir.Type) ir.Value {
	t := value.Type()
	if ir.IsPointer(t) && t.(*ir.PointerType).ElemType.Equal(ir.I8) {
		cast := ir.NewBitCast(value, typ)
		c.Block.AddInstruction(cast)
		return cast
	}
	return nil
}

func CastNumberExpr(in ir.Constant, out ir.Type) ir.Constant {
	return in
	//TO-DO implement
}

func CastNumber(c *Context, in ir.Value, out ir.Type) ir.Value {
	return in
	//TO-DO implement
	/*
		inSize := NumberSize(in.Type())
		outSize := NumberSize(out)

		if ir.Equal(in.Type(), out) {
			return in
		}

		if i, ok := in.(*ir.Int); ok && ir.IsInt(out) {
			i.Typ = out.(*ir.IntType)
			return i
		}

		if f, ok := in.(*ir.Float); ok && ir.IsFloat(out) {
			f.Typ = out.(*ir.FloatType)
			return f
		}

		if ir.Equal(to, ir.Void) {
			return nil, nil
		}

		if ir.IsPointer(inType) && ir.IsPointer(to) {
			return prog.Compiler.CurrentBlock().NewBitCast(in, to), nil
		}

		if fromFloat && toInt {
			return prog.Compiler.CurrentBlock().NewFPToSI(in, to), nil
		}

		if fromInt && toFloat {
			return prog.Compiler.CurrentBlock().NewSIToFP(in, to), nil
		}

		if fromInt && toInt {
			if inSize < outSize {
				return prog.Compiler.CurrentBlock().NewSExt(in, to), nil
			}
			if inSize == outSize {
				return in, nil
			}
			return prog.Compiler.CurrentBlock().NewTrunc(in, to), nil
		}

		if fromFloat && toFloat {
			if inSize < outSize {
				return prog.Compiler.CurrentBlock().NewFPExt(in, to), nil
			}
			if inSize == outSize {
				return in, nil
			}
			return prog.Compiler.CurrentBlock().NewFPTrunc(in, to), nil
		}

		// If the cast would not change the type, just return the in value, nil
		if ir.Equal(inType, to) {
			return in, nil
		}

		if ir.IsPointer(inType) && ir.IsInt(to) {
			return prog.Compiler.CurrentBlock().NewPtrToInt(in, to), nil
		}

		if ir.IsInt(inType) && ir.IsPointer(to) {
			return prog.Compiler.CurrentBlock().NewIntToPtr(in, to), nil
		}

		return nil, fmt.Errorf("Failed to typecast type %s to %s", inType.String(), to)
	*/
}
