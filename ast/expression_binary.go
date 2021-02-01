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
func (b *Binary) Type(c *Context, expected ir.Type) ir.Type {
	t1 := b.Left.Type(c, expected)
	t2 := b.Right.Type(c, expected)

	switch b.Operator {
	// bitwise code
	case token.LeftShift, token.RightShift, token.BitXor, token.BitOr, token.BitAnd:
		if ir.IsInt(t1) && ir.IsInt(t2) {
			return t1
		}

	// assign
	case token.Assign, token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		// TO-DO assert right type // equal or convertable
		return t1

	// logic operator
	case token.Or, token.And:
		if ir.IsBool(t1) && ir.IsBool(t2) {
			return ir.I1
		}

	// compare
	case token.Less, token.LessEqual, token.Greater, token.GreaterEqual, token.Equal, token.NotEqual:
		if ir.IsNumber(t1) && ir.IsNumber(t2) {
			return ir.I1
		}

	// arithmetic operator
	case token.Plus, token.Minus, token.Mul, token.Div, token.Rem:
		t, err := PromoteNumberType(t1, t2)
		if err == nil {
			return t
		}
	}

	c.Program.Error(b.Position, "invalid type for binary expression")
	return nil
}

//TO-DO operator overload
func (b *Binary) GenerateIR(c *Context, expected ir.Type) ir.Value {
	t1 := b.Left.Type(c, expected)
	t2 := b.Right.Type(c, expected)
	c1 := b.Left.IsConstant(c.Program)
	c2 := b.Right.IsConstant(c.Program)
	if c1 && c2 {
		c.Program.Error(b.Position, "unexpected const expression")
		return nil
	}
	var v1 ir.Value
	var v2 ir.Value
	if c1 {
		if expected == nil {
			v1 = b.Left.GenerateConstIR(c.Program, b.Right.Type(c, nil))
		} else {
			v1 = b.Left.GenerateConstIR(c.Program, expected)
		}
	} else {
		v1 = c.AutoLoad(b.Left.GenerateIR(c, expected))
	}
	if c2 {
		if expected == nil {
			v2 = b.Right.GenerateConstIR(c.Program, b.Left.Type(c, nil))
		} else {
			v2 = b.Right.GenerateConstIR(c.Program, expected)
		}
	} else {
		v2 = c.AutoLoad(b.Right.GenerateIR(c, expected))
	}

	var inst ir.Instruction
	switch b.Operator {
	case token.Assign:
		//TO-DO
		//Left cannot be const

	case token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		//TO-DO
		//Left cannot be const

	case token.Or, token.And:
		if ir.IsBool(t1) && ir.IsBool(t2) {
			if b.Operator == token.Or {
				inst = ir.NewOr(v1, v2)
			} else {
				inst = ir.NewAnd(v1, v2)
			}
		}

	default:
		t, e := PromoteNumberType(t1, t2)
		if e == nil {
			if !t1.Equal(t) {
				v1 = CastNumber(c, v1, t)
			}
			if !t2.Equal(t) {
				v2 = CastNumber(c, v2, t)
			}
			switch b.Operator {
			case token.BitOr:
				if ir.IsInt(t) {
					inst = ir.NewOr(v1, v2)
				}

			case token.BitXor:
				if ir.IsInt(t) {
					inst = ir.NewXor(v1, v2)
				}

			case token.BitAnd:
				if ir.IsInt(t) {
					inst = ir.NewAnd(v1, v2)
				}

			case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
				if ir.IsInt(t) {
					var icmp ir.IPred
					if t.(*ir.IntType).Unsigned {
						icmp = UICMP[b.Operator]
					} else {
						icmp = ICMP[b.Operator]
					}
					inst = ir.NewICmp(icmp, v1, v2)
				} else {
					fmp := FCMP[b.Operator]
					inst = ir.NewFCmp(fmp, v1, v2)
				}

			case token.LeftShift:
				if ir.IsInt(t) {
					inst = ir.NewShl(v1, v2)
				}

			case token.RightShift:
				if ir.IsInt(t) {
					if t.(*ir.IntType).Unsigned {
						inst = ir.NewLShr(v1, v2)
					} else {
						inst = ir.NewAShr(v1, v2)
					}
				}

			case token.Plus:
				if ir.IsInt(t) {
					inst = ir.NewAdd(v1, v2)
				} else if ir.IsFloat(t) {
					inst = ir.NewFAdd(v1, v2)
				}

			case token.Minus:
				if ir.IsInt(t) {
					inst = ir.NewSub(v1, v2)
				} else if ir.IsFloat(t) {
					inst = ir.NewFSub(v1, v2)
				}

			case token.Mul:
				if ir.IsInt(t) {
					inst = ir.NewMul(v1, v2)
				} else if ir.IsFloat(t) {
					inst = ir.NewFMul(v1, v2)
				}

			case token.Div:
				if ir.IsInt(t) {
					if t.(*ir.IntType).Unsigned {
						inst = ir.NewUDiv(v1, v2)
					} else {
						inst = ir.NewSDiv(v1, v2)
					}
				} else if ir.IsFloat(t) {
					inst = ir.NewFDiv(v1, v2)
				}

			case token.Rem:
				if ir.IsInt(t) {
					if t.(*ir.IntType).Unsigned {
						inst = ir.NewURem(v1, v2)
					} else {
						inst = ir.NewSRem(v1, v2)
					}
				}
			}
		}
	}
	if inst == nil {
		c.Program.Error(b.Position, "invalid type for binary expression")
		return nil
	}
	c.Block.AddInstruction(inst)
	return inst.(ir.Value)
}

func (b *Binary) IsConstant(p *Program) bool {
	return b.Left.IsConstant(p) && b.Right.IsConstant(p)
}

func (b *Binary) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	v1 := b.Left.GenerateConstIR(p, expected)
	v2 := b.Right.GenerateConstIR(p, expected)
	t1 := v1.Type()
	t2 := v2.Type()

	switch b.Operator {
	case token.Assign, token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		p.Error(b.Position, "invalid constant expression")
		return nil

	case token.Or, token.And:
		if ir.IsBool(t1) && ir.IsBool(t2) {
			if b.Operator == token.Or {
				return ir.NewExprOr(v1, v2)
			} else {
				return ir.NewExprAnd(v1, v2)
			}
		}

	default:
		t, e := PromoteNumberType(t1, t2)
		if e == nil {
			if !t1.Equal(t) {
				v1 = CastNumberExpr(v1, t)
			}
			if !t2.Equal(t) {
				v2 = CastNumberExpr(v2, t)
			}
			switch b.Operator {
			case token.BitOr:
				if ir.IsInt(t) {
					return ir.NewExprOr(v1, v2)
				}

			case token.BitXor:
				if ir.IsInt(t) {
					return ir.NewExprXor(v1, v2)
				}

			case token.BitAnd:
				if ir.IsInt(t) {
					return ir.NewExprAnd(v1, v2)
				}

			case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
				if ir.IsInt(t) {
					var icmp ir.IPred
					if t.(*ir.IntType).Unsigned {
						icmp = UICMP[b.Operator]
					} else {
						icmp = ICMP[b.Operator]
					}
					return ir.NewExprICmp(icmp, v1, v2)
				} else {
					fmp := FCMP[b.Operator]
					return ir.NewExprFCmp(fmp, v1, v2)
				}

			case token.LeftShift:
				if ir.IsInt(t) {
					return ir.NewExprShl(v1, v2)
				}

			case token.RightShift:
				if ir.IsInt(t) {
					if t.(*ir.IntType).Unsigned {
						return ir.NewExprLShr(v1, v2)
					} else {
						return ir.NewExprAShr(v1, v2)
					}
				}

			case token.Plus:
				if ir.IsInt(t) {
					return ir.NewExprAdd(v1, v2)
				} else if ir.IsFloat(t) {
					return ir.NewExprFAdd(v1, v2)
				}

			case token.Minus:
				if ir.IsInt(t) {
					return ir.NewExprSub(v1, v2)
				} else if ir.IsFloat(t) {
					return ir.NewExprFSub(v1, v2)
				}

			case token.Mul:
				if ir.IsInt(t) {
					return ir.NewExprMul(v1, v2)
				} else if ir.IsFloat(t) {
					return ir.NewExprFMul(v1, v2)
				}

			case token.Div:
				if ir.IsInt(t) {
					if t.(*ir.IntType).Unsigned {
						return ir.NewExprUDiv(v1, v2)
					} else {
						return ir.NewExprSDiv(v1, v2)
					}
				} else if ir.IsFloat(t) {
					return ir.NewExprFDiv(v1, v2)
				}

			case token.Rem:
				if ir.IsInt(t) {
					if t.(*ir.IntType).Unsigned {
						return ir.NewExprURem(v1, v2)
					} else {
						return ir.NewExprSRem(v1, v2)
					}
				}
			}
		}
	}
	p.Error(b.Position, "invalid constant expression")
	return nil
}
