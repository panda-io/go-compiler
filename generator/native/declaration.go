package native

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

func writeDeclaration(d declaration.Declaration, indent int, w *writer) {
	writeIndent(indent, w)
	switch t := d.(type) {
	case *declaration.Function:
		if t.TypeParameters != nil {
			writeType(t.TypeParameters, w)
		}
		if t.ClassName != "" && t.ClassName != t.Identifier() {
			w.buffer.WriteString("virtual ")
		}
		if !(t.ClassName != "" && (t.Identifier() == t.ClassName || t.Identifier() == "~"+t.ClassName)) {
			if t.ReturnType == nil {
				w.buffer.WriteString("void ")
			} else {
				writeType(t.ReturnType, w)
				w.buffer.WriteString(" ")
			}
		}
		w.buffer.WriteString(t.Identifier())
		w.buffer.WriteString("(")
		writeType(t.Parameters, w)
		w.buffer.WriteString(");\n")

	case *declaration.Enum:
		w.buffer.WriteString("enum class " + t.Identifier() + "\n")
		writeIndent(indent, w)
		w.buffer.WriteString("{\n")
		for _, m := range t.Members {
			writeDeclaration(m, indent+tabSize, w)
		}
		w.buffer.WriteString("};\n")

	case *declaration.Interface:
		w.buffer.WriteString("class " + t.Identifier())
		if len(t.Parents) > 0 {
			w.buffer.WriteString(" : public ")
			for i, n := range t.Parents {
				if i > 0 {
					w.buffer.WriteString(", ")
				}
				writeType(n, w)
			}
		}
		w.buffer.WriteString("\n")
		writeIndent(indent, w)
		w.buffer.WriteString("{\npublic:\n")
		for _, m := range t.Members {
			writeDeclaration(m, indent+tabSize, w)
		}
		w.buffer.WriteString("};\n")

	case *declaration.Class:
		w.buffer.WriteString("class " + t.Identifier())
		if len(t.Parents) > 0 {
			w.buffer.WriteString(" : public ")
			for i, n := range t.Parents {
				if i > 0 {
					w.buffer.WriteString(", ")
				}
				writeType(n, w)
			}
		}
		w.buffer.WriteString("\n")
		writeIndent(indent, w)
		w.buffer.WriteString("{\npublic:\n")
		for _, m := range t.Members {
			writeDeclaration(m, indent+tabSize, w)
		}
		w.buffer.WriteString("};\n")
	}
}

func writeClass(c *declaration.Class, w *writer) {
	for _, m := range c.Members {
		if f, ok := m.(*declaration.Function); ok {
			writeFunction(f, w)
		}
	}
}

func writeVariable(v *declaration.Variable, w *writer) {
	writeType(v.Type, w)
	w.buffer.WriteString(" ")
	w.buffer.WriteString(v.Identifier())
	if v.Value != nil {
		w.buffer.WriteString(" = ")
		writeExpression(v.Value, w)
	}
	w.buffer.WriteString(";\n")

}

func writeFunction(f *declaration.Function, w *writer) {
	if f.TypeParameters != nil {
		writeType(f.TypeParameters, w)
	}
	if !(f.ClassName != "" && (f.Identifier() == f.ClassName || f.Identifier() == "~"+f.ClassName)) {
		if f.ReturnType == nil {
			w.buffer.WriteString("void ")
		} else {
			writeType(f.ReturnType, w)
			w.buffer.WriteString(" ")
		}
	}
	if f.ClassName != "" {
		w.buffer.WriteString(f.ClassName + "::")
	}
	w.buffer.WriteString(f.Identifier())
	w.buffer.WriteString("(")
	writeType(f.Parameters, w)
	w.buffer.WriteString(")\n")
	writeStatement(f.Body, 0, w)
}
