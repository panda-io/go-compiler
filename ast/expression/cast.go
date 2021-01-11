package expression

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

func NumberSize(typ ir.Type) int {
	switch t := t.(type) {
	case *types.IntType:
		return int(t.BitSize)
	case *types.FloatType:
		return gtypes.FloatBitSize(t)
	}

	return -1
}
*/
