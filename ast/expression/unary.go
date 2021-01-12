package expression

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Unary struct {
	Base
	Operator   token.Token
	Expression Expression
}

func (u *Unary) Type(c *node.Context) ir.Type {
	switch u.Operator {
	case token.Plus, token.Minus:
		t := u.Expression.Type(c)
		if ir.IsNumber(t) {
			return t
		}
		c.Error(u.Position, fmt.Sprintf("%s is not valid for '%s' operator", t.String(), u.Operator.String()))
		return nil

	case token.Not:
		// logic not
		if ir.IsInt(u.Expression.Type(c)) {
			return ir.I1
		}
		c.Error(u.Position, "only int or bool can use '!' (not) operator")
		return nil

	case token.Complement:
		// TO-DO
	}
	return nil
}

func (u *Unary) GenerateIR(c *node.Context) ir.Value {
	// token.Plus, token.Minus, token.Not, token.Complement:
	switch u.Operator {
	case token.Plus:
		return u.Expression.GenerateIR(c)

	case token.Minus:
		if ir.IsFloat(u.Expression.Type(c)) {
			value := ir.NewFNeg(u.Expression.GenerateIR(c))
			c.Block.AddInstruction(value)
			return value
		} else if ir.IsInt(u.Expression.Type(c)) {
			t := u.Expression.Type(c).(*ir.IntType)
			value := ir.NewSub(ir.NewInt(t, 0), u.Expression.GenerateIR(c))
			c.Block.AddInstruction(value)
			return value
		}
		c.Error(u.Position, "only float or int is allowed in negative unary expression")
		return nil

	case token.Not:
		// logic not
		if ir.IsInt(u.Expression.Type(c)) {
			// int or bool
			/*
				opVal, _ := createTypeCast(prog, operandValue, types.I1)

				eq := prog.Compiler.CurrentBlock().NewICmp(enum.IPredNE, opVal, constant.False)
				inv := prog.Compiler.CurrentBlock().NewXor(eq, constant.True)
				ext := prog.Compiler.CurrentBlock().NewZExt(inv, types.I32)

				return ext, nil*/
		}
		c.Error(u.Position, "only int or bool can use '!' (not) operator")
		return nil

	case token.Complement:
		// TO-DO
	}
	return nil
}
