package native

import (
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/token"
)

func writeStatement(s statement.Statement, indent int, w *writer) {
	switch t := s.(type) {
	case *statement.Empty:
		w.buffer.WriteString(";")

	case *statement.Raw:
		w.buffer.WriteString(t.Source)

	case *statement.Expression:
		writeExpression(t.Expression, w)

	case *statement.Declaration:
		//TO-DO

	case *statement.Return:
		w.buffer.WriteString("return")
		if t.Expression != nil {
			w.buffer.WriteString(" ")
			writeExpression(t.Expression, w)
		}

	case *statement.Continue:
		w.buffer.WriteString("continue")

	case *statement.Break:
		w.buffer.WriteString("break")

	case *statement.Try:
		w.buffer.WriteString("try")

	case *statement.Throw:
		w.buffer.WriteString("throw ")
		writeExpression(t.Expression, w)

	case *statement.If:
		writeIndent(indent, w)
		w.buffer.WriteString("if (")
		if t.Initialization != nil {
			writeStatement(t.Initialization, 0, w)
			w.buffer.WriteString("; ")
		}
		writeStatement(t.Condition, 0, w)
		w.buffer.WriteString(")\n")
		writeIndent(indent, w)
		w.buffer.WriteString("{\n")
		writeStatement(t.Body, indent+tabSize, w)
		writeIndent(indent, w)
		w.buffer.WriteString("}\n")

	case *statement.Switch:
		writeIndent(indent, w)
		w.buffer.WriteString("switch (")
		if t.Initialization != nil {
			writeStatement(t.Initialization, 0, w)
			w.buffer.WriteString("; ")
		}
		writeStatement(t.Operand, 0, w)
		w.buffer.WriteString(")\n")
		writeIndent(indent, w)
		w.buffer.WriteString("{\n")

		for _, c := range t.Body {
			if c.Token == token.Case {
				writeIndent(indent+tabSize, w)
				w.buffer.WriteString("case ")
				writeExpression(c.Case, w)
				w.buffer.WriteString(":\n")
				if c.Body != nil {
					writeStatement(c.Body, indent+tabSize*2, w)
				}
			} else {
				writeIndent(indent+tabSize, w)
				w.buffer.WriteString("default:\n")
				if c.Body != nil {
					writeStatement(c.Body, indent+tabSize*2, w)
				}
			}
			w.buffer.WriteString("\n")
		}
		writeIndent(indent, w)
		w.buffer.WriteString("}\n")

	case *statement.For:
		writeIndent(indent, w)
		w.buffer.WriteString("for (")
		if t.Initialization != nil {
			writeStatement(t.Initialization, 0, w)
		}
		w.buffer.WriteString(";")
		if t.Condition != nil {
			w.buffer.WriteString(" ")
			writeStatement(t.Condition, 0, w)
		}
		w.buffer.WriteString(";")
		if t.Post != nil {
			w.buffer.WriteString(" ")
			writeStatement(t.Post, 0, w)
		}
		w.buffer.WriteString(")\n")
		writeStatement(t.Body, indent+tabSize, w)

	case *statement.Foreach:
		/* TO-DO
		WriteIndent(buffer, indent)
		buffer.WriteString("for (")
		s.Init.Print(buffer, 0)
		buffer.WriteString(" : ")
		s.Iterator.Print(buffer)
		buffer.WriteString(" )\n")
		s.Body.Print(buffer, indent)*/

	case *statement.Compound:
		writeIndent(indent, w)
		w.buffer.WriteString("{\n")
		for _, v := range t.Statements {
			writeIndent(indent+tabSize, w)
			writeStatement(v, indent+tabSize, w)
			w.buffer.WriteString(";\n")
		}
		writeIndent(indent, w)
		w.buffer.WriteString("}\n")
	}
}
