package expression

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

func PromoteNumberType(c *node.Context, e1 Expression, e2 Expression) ir.Type {
	t1 := e1.Type(c)
	t2 := e2.Type(c)
	if ir.IsInt(t1) {
		if ir.IsInt(t2) {
			if Sizeof(t1) > Sizeof(t2) {
				return t1
			}
			return t2
		} else if ir.IsFloat(t2) {
			return t2
		}
		c.Error(e2.GetPosition(), "invalid number")
		return nil
	} else if ir.IsFloat(t1) {
		if ir.IsInt(t2) {
			return t1
		} else if ir.IsFloat(t2) {
			if Sizeof(t1) > Sizeof(t2) {
				return t1
			}
			return t2
		}
		c.Error(e2.GetPosition(), "invalid number")
		return nil
	}
	c.Error(e1.GetPosition(), "invalid number")
	return nil
}

// also string here
func Cast(c *node.Context, in ir.Value, outType ir.Type) ir.Value {
	inType := in.Type()
	if ir.IsPointer(inType) && ir.IsPointer(outType) {
		if outType.Equal(ir.NewPointerType(ir.I8)) {
			// convert to raw pointer
			gep := ir.NewGetElementPtr(inType.(*ir.PointerType).ElemType, in, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			c.Block.AddInstruction(gep)
			return gep
		} /*else {
			// TO-DO convert
			check if convertable, oop inheritance
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

//? TO-DO pointer size?
func Sizeof(typ ir.Type) int {
	switch t := typ.(type) {
	case *ir.IntType:
		return int(t.BitSize)

	case *ir.FloatType:
		switch t.Kind {
		case ir.FloatKindFloat:
			return 32

		case ir.FloatKindDouble:
			return 64

		default:
			panic(fmt.Errorf("floating-point kind %q is not implemented", t.Kind))
		}
	}
	return -1
}

/*
func CastNumber(c *node.Context, in ir.Value, out ir.Type) ir.Value {
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
