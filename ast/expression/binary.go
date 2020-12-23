package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Binary struct {
	Base
	Left     Expression
	Operator token.Token
	Right    Expression
}

//TO-DO operator overload
func (b *Binary) GenerateIR(c *node.Context) ir.Value {
	switch b.Operator {
	case token.Assign, token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
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

	case token.Equal, token.NotEqual:
		//TO-DO
		return nil

	case token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		//TO-DO
		return nil

	case token.LeftShift, token.RightShift:
		//TO-DO
		return nil

	case token.Plus:
		l := b.Left.GenerateIR(c)
		r := b.Right.GenerateIR(c)

		if ir.IsInt(l.Type()) && ir.IsInt(r.Type()) {
			add := ir.NewAdd(l, r)
			c.Block.AddInstruction(add)
			return add
		} else if ir.IsFloat(l.Type()) && ir.IsFloat(r.Type()) {
			add := ir.NewFAdd(l, r)
			c.Block.AddInstruction(add)
			return add
		} else {
			c.Error(b.Left.GetPosition(), "type mismatch for add")
		}

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
