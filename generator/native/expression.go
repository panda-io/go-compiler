package native

import (
	"strconv"

	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/token"
)

func writeExpression(e expression.Expression, w *writer) {
	switch t := e.(type) {
	case *expression.Literal:
		switch t.Type {
		case token.STRING:
			if t.Value[0] == '"' {
				w.buffer.WriteString(t.Value)
			} else if t.Value[0] == '`' {
				w.buffer.WriteString(strconv.Quote(t.Value[1 : len(t.Value)-1]))
			}

		case token.INT, token.FLOAT:
			w.buffer.WriteString(t.Value)

		case token.CHAR:
			// char32_t
			w.buffer.WriteString(t.Value)

		case token.BOOL, token.Void, token.NULL:
			w.buffer.WriteString(t.Value)
		}

	case *expression.Identifier:
		w.buffer.WriteString(t.Name)

	case *expression.This:
		w.buffer.WriteString("this")

	case *expression.Super:
		//TO-DO // check parent type

	case *expression.Parentheses:
		w.buffer.WriteString("(")
		writeExpression(t.Expression, w)
		w.buffer.WriteString(")")

	case *expression.MemberAccess:
		writeExpression(t.Parent, w)
		w.buffer.WriteString(".")
		writeExpression(t.Member, w)

	case *expression.Subscripting:
		writeExpression(t.Parent, w)
		w.buffer.WriteString("[")
		writeExpression(t.Element, w)
		w.buffer.WriteString("]")

	case *expression.Invocation:
		//TO-DO
		/*
					x.Func.Print(buffer)
			buffer.WriteString("(")
			for i, v := range x.Args {
				if i != 0 {
					buffer.WriteString(", ")
				}
				v.Print(buffer)
			}
			buffer.WriteString(")")*/

	case *expression.New:
		//TO-DO
		/*
					buffer.WriteString("std::make_shared<//TO-DO>(")
			for i, v := range x.Args {
				if i != 0 {
					buffer.WriteString(", ")
					v.Print(buffer)
				}
			}
			buffer.WriteString(")")*/

	case *expression.Increment:
		writeExpression(t.Expression, w)
		w.buffer.WriteString("++")

	case *expression.Decrement:
		writeExpression(t.Expression, w)
		w.buffer.WriteString("--")

	case *expression.Unary:
		w.buffer.WriteString(t.Operator.String())
		writeExpression(t.Expression, w)

	case *expression.Binary:
		writeExpression(t.Left, w)
		w.buffer.WriteString(" ")
		w.buffer.WriteString(t.Operator.String())
		w.buffer.WriteString(" ")
		writeExpression(t.Right, w)

	case *expression.Conditional:
		writeExpression(t.Condition, w)
		w.buffer.WriteString(" ? ")
		writeExpression(t.First, w)
		w.buffer.WriteString(" : ")
		writeExpression(t.Second, w)
	}
}
