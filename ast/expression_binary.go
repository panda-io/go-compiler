package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

var (
	ICMP = map[token.Token]ir.IPred{
		token.Equal:        ir.IPredEQ,
		token.NotEqual:     ir.IPredNE,
		token.Less:         ir.IPredSLT,
		token.LessEqual:    ir.IPredSLE,
		token.Greater:      ir.IPredSGT,
		token.GreaterEqual: ir.IPredSGE,
	}

	UICMP = map[token.Token]ir.IPred{
		token.Equal:        ir.IPredEQ,
		token.NotEqual:     ir.IPredNE,
		token.Less:         ir.IPredULT,
		token.LessEqual:    ir.IPredULE,
		token.Greater:      ir.IPredUGT,
		token.GreaterEqual: ir.IPredUGE,
	}

	FCMP = map[token.Token]ir.FPred{
		token.Equal:        ir.FPredOEQ,
		token.NotEqual:     ir.FPredONE,
		token.Less:         ir.FPredOLT,
		token.LessEqual:    ir.FPredOLE,
		token.Greater:      ir.FPredOGT,
		token.GreaterEqual: ir.FPredOGE,
	}
)

type Binary struct {
	ExpressionBase
	Left     Expression
	Operator token.Token
	Right    Expression
}

//TO-DO operator override
//TO-DO check if const
func (b *Binary) Type(c *Context) ir.Type {
	switch b.Operator {
	// bitwise code
	case token.LeftShift, token.RightShift, token.BitXor, token.BitOr, token.BitAnd, token.Not:
		t, err := PromoteNumberType(c, b.Left.Type(c), b.Right.Type(c))
		if err != nil {
			c.Program.Error(b.Position, err.Error())
			return nil
		}
		if ir.IsInt(t) {
			t1 := b.Left.Type(c)
			if ir.IsPointer(t1) {
				return t1.(*ir.PointerType).ElemType
			}
			return t1
		}
		c.Program.Error(b.Position, "only int are valid for bitwise operator")
		return nil

	// assign
	case token.Assign, token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		// TO-DO assert left type
		return b.Right.Type(c)

	// logic operator
	case token.Or, token.And:
		if ir.IsBool(b.Left.Type(c)) && ir.IsBool(b.Right.Type(c)) {
			return ir.I1
		}
		c.Program.Error(b.Position, "invalid type for binary operator")
		return nil

	case token.Less, token.LessEqual, token.Greater, token.GreaterEqual, token.Equal, token.NotEqual:
		t, err := PromoteNumberType(c, b.Left.Type(c), b.Right.Type(c))
		if err != nil {
			c.Program.Error(b.Position, err.Error())
			return nil
		}
		if ir.IsNumber(t) {
			return ir.I1
		}
		c.Program.Error(b.Position, "invalid type for binary operator")
		return nil

	//arithmetic operator
	case token.Plus, token.Minus, token.Mul, token.Div, token.Rem:
		t, err := PromoteNumberType(c, b.Left.Type(c), b.Right.Type(c))
		if err != nil {
			c.Program.Error(b.Position, err.Error())
		}
		return t

	default:
		c.Program.Error(b.Position, "invalid type for binary expression")
		return nil
	}
}

//TO-DO operator overload
func (b *Binary) GenerateIR(c *Context) ir.Value {
	//shiftOperation := false
	//integerOperation := false

	switch b.Operator {
	case token.LeftShift, token.RightShift:
		//shiftOperation = true

	case token.BitXor, token.BitOr, token.BitAnd, token.Not:
		//integerOperation = true
	}

	switch b.Operator {
	case token.Assign:
		/*
			t := l.Type()
			prog.Compiler.PushType(targetType)

			val, err := n.Value.GenAccess(prog)
			if err != nil {
				return nil, err
			}

			if targetType != nil && !ir.Equal(val.Type(), targetType) {
				val, err = createTypeCast(prog, val, targetType)
				if err != nil {
					return nil, err
				}
			}

			// fmt.Println(val)
			n.Assignee.GenAssign(prog, val)
			return val, nil
			return nil*/

	case token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		//TO-DO
		return nil

	case token.Question:
		//TO-DO
		return nil

	case token.Or:
		//TO-DO
		return nil

	case token.And:
		//TO-DO
		return nil

	case token.BitOr:
		//TO-DO
		return nil

	case token.BitXor:
		//TO-DO
		return nil

	case token.BitAnd:
		//TO-DO
		return nil

	case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		t, v1, v2, err := PromoteNumberValue(c, b.Left, b.Right)
		if err != nil {
			c.Program.Error(b.Position, err.Error())
		}
		if t == nil {
			return nil
		}

		if ir.IsInt(t) {
			var icmp ir.IPred
			if t.(*ir.IntType).Unsigned {
				icmp = UICMP[b.Operator]
			} else {
				icmp = ICMP[b.Operator]
			}
			cmp := ir.NewICmp(icmp, v1, v2)
			c.Block.AddInstruction(cmp)
			return cmp
		}

		fmp := FCMP[b.Operator]
		cmp := ir.NewFCmp(fmp, v1, v2)
		c.Block.AddInstruction(cmp)
		return cmp

	case token.LeftShift, token.RightShift:
		//TO-DO
		return nil

	case token.Plus:
		/*
			if ir.IsInt(l.Type()) && ir.IsInt(r.Type()) {
				add := ir.NewAdd(l, r)
				c.Block.AddInstruction(add)
				return add
			} else if ir.IsFloat(l.Type()) && ir.IsFloat(r.Type()) {
				add := ir.NewFAdd(l, r)
				c.Block.AddInstruction(add)
				return add
			} else {
				c.Program.Error(b.Left.GetPosition(), "type mismatch for add")
			}*/

		//TO-DO convert int to float
		//TO-DO vector add
		//TO-DO operator overload and string
		return nil

	case token.Minus:
		//TO-DO
		return nil

	case token.Mul:
		//TO-DO
		return nil

	case token.Div:
		//TO-DO
		return nil

	case token.Rem:
		//TO-DO
		return nil
	}
	return nil
}

/*
func (b *Binary) GenerateIR1(c *Context) ir.Value {
	if (isShiftOperation) // Shift operations are specials
	{
		switch (operand2.StackType)
		{
			case StackValueType.Int32:
				break;
			case StackValueType.NativeInt:
				value2 = LLVM.BuildPtrToInt(builder, value2, nativeIntLLVM, string.Empty);
				break;
			default:
				goto InvalidBinaryOperation;
		}

		// Check first operand, and convert second operand to match first one
		switch (operand1.StackType)
		{
			case StackValueType.Int32:
				value2 = LLVM.BuildIntCast(builder, value2, int32LLVM, string.Empty);
				break;
			case StackValueType.Int64:
				value2 = LLVM.BuildIntCast(builder, value2, int64LLVM, string.Empty);
				break;
			case StackValueType.NativeInt:
				value1 = LLVM.BuildPtrToInt(builder, value1, nativeIntLLVM, string.Empty);
				value2 = LLVM.BuildIntCast(builder, value2, nativeIntLLVM, string.Empty);
				break;
			default:
				goto InvalidBinaryOperation;
		}

		// Output type is determined by first operand
		outputOperandType = operand1;
	}
	else if (operand1.StackType == operand2.StackType) // Diagonal
	{
		// Check type
		switch (operand1.StackType)
		{
			case StackValueType.Int32:
			case StackValueType.Int64:
			case StackValueType.Float:
				outputOperandType = operand1;
				break;
			case StackValueType.NativeInt:
				value1 = LLVM.BuildPtrToInt(builder, value1, nativeIntLLVM, string.Empty);
				value2 = LLVM.BuildPtrToInt(builder, value2, nativeIntLLVM, string.Empty);
				outputOperandType = operand1;
				break;
			case StackValueType.Reference:
				if (opcode != Code.Sub && opcode != Code.Sub_Ovf_Un)
					goto InvalidBinaryOperation;
				value1 = LLVM.BuildPtrToInt(builder, value1, nativeIntLLVM, string.Empty);
				value2 = LLVM.BuildPtrToInt(builder, value2, nativeIntLLVM, string.Empty);
				outputOperandType = new StackValue(StackValueType.NativeInt, intPtr, ValueRef.Empty);
				break;
			default:
				throw new InvalidOperationException(string.Format("Binary operations are not allowed on {0}.", operand1.StackType));
		}
	}
	else if (operand1.StackType == StackValueType.NativeInt && operand2.StackType == StackValueType.Int32)
	{
		value1 = LLVM.BuildPtrToInt(builder, value1, nativeIntLLVM, string.Empty);
		value2 = LLVM.BuildIntCast(builder, value2, nativeIntLLVM, string.Empty);
		outputOperandType = operand1;
	}
	else if (operand1.StackType == StackValueType.Int32 && operand2.StackType == StackValueType.NativeInt)
	{
		value2 = LLVM.BuildPtrToInt(builder, value2, nativeIntLLVM, string.Empty);
		value1 = LLVM.BuildIntCast(builder, value1, nativeIntLLVM, string.Empty);
		outputOperandType = operand2;
	}
	else if (!isIntegerOperation
			 && (operand1.StackType == StackValueType.Reference || operand2.StackType == StackValueType.Reference)) // ref + [i32, nativeint] or [i32, nativeint] + ref
	{
		StackValue operandRef, operandInt;
		ValueRef valueRef, valueInt;

		if (operand2.StackType == StackValueType.Reference)
		{
			operandRef = operand2;
			operandInt = operand1;
			valueRef = value2;
			valueInt = value1;
		}
		else
		{
			operandRef = operand1;
			operandInt = operand2;
			valueRef = value1;
			valueInt = value2;
		}

		switch (operandInt.StackType)
		{
			case StackValueType.Int32:
				break;
			case StackValueType.NativeInt:
				valueInt = LLVM.BuildPtrToInt(builder, valueInt, nativeIntLLVM, string.Empty);
				break;
			default:
				goto InvalidBinaryOperation;
		}

		switch (opcode)
		{
			case Code.Add:
			case Code.Add_Ovf_Un:
				break;
			case Code.Sub:
			case Code.Sub_Ovf:
				if (operand2.StackType == StackValueType.Reference)
					goto InvalidBinaryOperation;

				valueInt = LLVM.BuildNeg(builder, valueInt, string.Empty);
				break;
			default:
				goto InvalidBinaryOperation;
		}

		// If necessary, cast to i8*
		var valueRefType = LLVM.TypeOf(valueRef);
		if (valueRefType != intPtrLLVM)
			valueRef = LLVM.BuildPointerCast(builder, valueRef, intPtrLLVM, string.Empty);

		valueRef = LLVM.BuildGEP(builder, valueRef, new[] {valueInt}, string.Empty);

		// Cast back to original type
		if (valueRefType != intPtrLLVM)
			valueRef = LLVM.BuildPointerCast(builder, valueRef, valueRefType, string.Empty);

		stack.Add(new StackValue(StackValueType.Reference, operandRef.Type, valueRef));

		// Early exit
		return;
	}
	else
	{
		goto InvalidBinaryOperation;
	}

	ValueRef result;

	// Perform binary operation
	if (operand1.StackType == StackValueType.Float)
	{
		switch (opcode)
		{
			case Code.Add: result = LLVM.BuildFAdd(builder, value1, value2, string.Empty); break;
			case Code.Sub: result = LLVM.BuildFSub(builder, value1, value2, string.Empty); break;
			case Code.Mul: result = LLVM.BuildFMul(builder, value1, value2, string.Empty); break;
			case Code.Div: result = LLVM.BuildFDiv(builder, value1, value2, string.Empty); break;
			case Code.Rem: result = LLVM.BuildFRem(builder, value1, value2, string.Empty); break;
			default:
				goto InvalidBinaryOperation;
		}
	}
	else
	{
		// Special case: char is size 1, not 2!
		if (CharUsesUTF8)
		{
			if (opcode == Code.Add && operand1.Type.TypeReferenceCecil.FullName == typeof(char*).FullName)
			{
				value2 = LLVM.BuildLShr(builder, value2, LLVM.ConstInt(int32LLVM, 1, false), string.Empty);
			}
			else if (opcode == Code.Add && operand2.Type.TypeReferenceCecil.FullName == typeof(char*).FullName)
			{
				value1 = LLVM.BuildLShr(builder, value1, LLVM.ConstInt(int32LLVM, 1, false), string.Empty);
			}
		}

		switch (opcode)
		{
			case Code.Add:          result = LLVM.BuildAdd(builder, value1, value2, string.Empty); break;
			case Code.Sub:          result = LLVM.BuildSub(builder, value1, value2, string.Empty); break;
			case Code.Mul:          result = LLVM.BuildMul(builder, value1, value2, string.Empty); break;
			case Code.Div:          result = LLVM.BuildSDiv(builder, value1, value2, string.Empty); break;
			case Code.Div_Un:       result = LLVM.BuildUDiv(builder, value1, value2, string.Empty); break;
			case Code.Rem:          result = LLVM.BuildSRem(builder, value1, value2, string.Empty); break;
			case Code.Rem_Un:       result = LLVM.BuildURem(builder, value1, value2, string.Empty); break;
			case Code.Shl:          result = LLVM.BuildShl(builder, value1, value2, string.Empty); break;
			case Code.Shr:          result = LLVM.BuildAShr(builder, value1, value2, string.Empty); break;
			case Code.Shr_Un:       result = LLVM.BuildLShr(builder, value1, value2, string.Empty); break;
			case Code.And:          result = LLVM.BuildAnd(builder, value1, value2, string.Empty); break;
			case Code.Or:           result = LLVM.BuildOr(builder, value1, value2, string.Empty); break;
			case Code.Xor:          result = LLVM.BuildXor(builder, value1, value2, string.Empty); break;
			case Code.Add_Ovf:
			case Code.Add_Ovf_Un:
			case Code.Sub_Ovf:
			case Code.Sub_Ovf_Un:
			case Code.Mul_Ovf:
			case Code.Mul_Ovf_Un:
			{
				Intrinsics intrinsicId;
				switch (opcode)
				{
					case Code.Add_Ovf:
						intrinsicId = Intrinsics.sadd_with_overflow;
						break;
					case Code.Add_Ovf_Un:
						intrinsicId = Intrinsics.uadd_with_overflow;
						break;
					case Code.Sub_Ovf:
						intrinsicId = Intrinsics.ssub_with_overflow;
						break;
					case Code.Sub_Ovf_Un:
						intrinsicId = Intrinsics.usub_with_overflow;
						break;
					case Code.Mul_Ovf:
						intrinsicId = Intrinsics.smul_with_overflow;
						break;
					case Code.Mul_Ovf_Un:
						intrinsicId = Intrinsics.umul_with_overflow;
						break;
					default:
						throw new ArgumentOutOfRangeException();
				}
				var intrinsic = LLVM.IntrinsicGetDeclaration(module, (uint)intrinsicId, new[] {LLVM.TypeOf(value1)});
				var overflowResult = LLVM.BuildCall(builder, intrinsic, new[] {value1, value2}, string.Empty);
				var hasOverflow = LLVM.BuildExtractValue(builder, overflowResult, 1, string.Empty);

				var nextBlock = LLVM.AppendBasicBlockInContext(context, functionGlobal, string.Empty);
				var overflowBlock = LLVM.AppendBasicBlockInContext(context, functionGlobal, "overflow");
				LLVM.MoveBasicBlockAfter(overflowBlock, LLVM.GetInsertBlock(builder));
				LLVM.MoveBasicBlockAfter(nextBlock, overflowBlock);

				LLVM.BuildCondBr(builder, hasOverflow, overflowBlock, nextBlock);

				LLVM.PositionBuilderAtEnd(builder, overflowBlock);

				// Create OverflowException object
				var overflowExceptionClass = GetClass(corlib.MainModule.GetType(typeof(OverflowException).FullName));
				EmitNewobj(functionContext, overflowExceptionClass.Type, overflowExceptionClass.Functions.Single(x => x.MethodReference.Name == ".ctor" && x.MethodReference.Parameters.Count == 0));
				var overflowException = stack.Pop();
				GenerateInvoke(functionContext, throwExceptionFunctionLLVM, new[] {LLVM.BuildPointerCast(builder, overflowException.Value, LLVM.TypeOf(LLVM.GetParam(throwExceptionFunctionLLVM, 0)), string.Empty)});
				LLVM.BuildUnreachable(builder);

				functionContext.BasicBlock = nextBlock;
				LLVM.PositionBuilderAtEnd(builder, nextBlock);
				result = LLVM.BuildExtractValue(builder, overflowResult, 0, string.Empty);

				break;
			}
			default:
				goto InvalidBinaryOperation;
		}

		if (CharUsesUTF8)
		{
			if (opcode == Code.Sub
				&& operand1.Type.TypeReferenceCecil.FullName == typeof(char*).FullName
				&& operand2.Type.TypeReferenceCecil.FullName == typeof(char*).FullName)
			{
				result = LLVM.BuildLShr(builder, result, LLVM.ConstInt(int32LLVM, 1, false), string.Empty);
			}
		}
	}

	Type outputType;

	switch (outputOperandType.StackType)
	{
		case StackValueType.Int32:
		case StackValueType.Int64:
		case StackValueType.Float:
			// No output conversion required, as it could only have been from same input types (non-shift) or operand 1 (shift)
			outputType = operand1.Type;
			break;
		case StackValueType.NativeInt:
			outputType = intPtr;
			result = LLVM.BuildIntToPtr(builder, result, intPtrLLVM, string.Empty);
			break;
		case StackValueType.Reference:
			result = LLVM.BuildIntToPtr(builder, result, intPtrLLVM, string.Empty);

			// Get type from one of its operand (if output is reference type, one of the two operand must be too)
			if (operand1.StackType == StackValueType.Reference)
				outputType = operand1.Type;
			else if (operand2.StackType == StackValueType.Reference)
				outputType = operand2.Type;
			else
				goto InvalidBinaryOperation;
			break;
		default:
			goto InvalidBinaryOperation;
	}

	stack.Add(new StackValue(outputOperandType.StackType, outputOperandType.Type, result));

	return;

}*/

func (*Binary) IsConstant(p *Program) bool {
	//TO-DO
	return false
}

func (*Binary) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	//TO-DO
	return nil
}
