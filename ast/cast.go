package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

func PromoteNumberType(c *Context, t1 ir.Type, t2 ir.Type) (ir.Type, error) {
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

//TO-DO expected type and constexpr
func PromoteNumberValue(c *Context, e1 Expression, e2 Expression) (t ir.Type, v1 ir.Value, v2 ir.Value, e error) {
	v1 = c.AutoLoad(e1.GenerateIR(c))
	v2 = c.AutoLoad(e2.GenerateIR(c))
	t1 := e1.Type(c)
	t2 := e2.Type(c)
	t, e = PromoteNumberType(c, t1, t2)
	if e != nil {
		return
	}
	// cast
	if !t1.Equal(t) {
		v1 = Cast(c, v1, t)
	}
	if !t2.Equal(t) {
		v2 = Cast(c, v2, t)
	}
	return
}

// also string here
func Cast(c *Context, in ir.Value, outType ir.Type) ir.Value {
	inType := in.Type()
	if ir.IsPointer(inType) && ir.IsPointer(outType) {
		if outType.Equal(ir.NewPointerType(ir.I8)) {
			// convert to raw pointer
			//gep := ir.NewGetElementPtr(inType.(*ir.PointerType).ElemType, in, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			cast := ir.NewBitCast(in, ir.NewPointerType(ir.I8))
			c.Block.AddInstruction(cast)
			return cast
		} /*else {
			// TO-DO convert
			check if convertable, oop inheritance
			// Check OO inheritance
		}*/
		return nil
	}

	if ir.IsPointer(inType) || ir.IsPointer(outType) {
		// cannot convert between pointer and value
		return nil
	}
	/*
		inType := in.Type()
		fromInt := types.IsInt(inType)
		fromFloat := types.IsFloat(inType)

		toInt := types.IsInt(to)
		toFloat := types.IsFloat(to)

		inSize := sizeof(in.Type())
		outSize := sizeof(out)

		// If the cast would not change the type, just return the in value, nil
		if types.Equal(inType, to) {
			return in, nil
		}

		if c, ok := in.(*constant.Int); ok && types.IsInt(to) {
			c.Typ = to.(*types.IntType)
			return c, nil
		}

		if c, ok := in.(*constant.Float); ok && types.IsFloat(to) {
			c.Typ = to.(*types.FloatType)
			return c, nil
		}

		if types.Equal(to, types.Void) {
			return nil, nil
		}

		if types.IsPointer(inType) && types.IsPointer(to) {
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
		if types.Equal(inType, to) {
			return in, nil
		}

		if types.IsPointer(inType) && types.IsInt(to) {
			return prog.Compiler.CurrentBlock().NewPtrToInt(in, to), nil
		}

		if types.IsInt(inType) && types.IsPointer(to) {
			return prog.Compiler.CurrentBlock().NewIntToPtr(in, to), nil
		}

		return nil, fmt.Errorf("Failed to typecast type %s to %s", inType.String(), to)*/
	return nil
}

/*
func CastNumber(c *Context, in ir.Value, out ir.Type) ir.Value {
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
}
*/
