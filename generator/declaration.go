package generator

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

func writeDeclaration(p *ast.Package, d declaration.Declaration, indent int, w *writer) {
	writeIndent(indent, w)
	switch t := d.(type) {
	case *declaration.Variable:
		_, isTypeName := t.Type.(*types.TypeName)
		if isTypeName {
			w.buffer.WriteString("std::shared_ptr<")
		}
		writeType(t.Type, w)
		if isTypeName {
			w.buffer.WriteString(">")
		}
		w.buffer.WriteString(" ")
		w.buffer.WriteString(t.Identifier())
		if t.Value != nil {
			w.buffer.WriteString(" = ")
			writeExpression(t.Value, w)
		} else if b, ok := t.Type.(*types.BuitinType); ok {
			switch b.Token {
			case token.Int8, token.Int16, token.Int32, token.Int64,
				token.Uint8, token.Uint16, token.Uint32, token.Uint64,
				token.SByte, token.Short, token.Int, token.Long,
				token.Byte, token.Ushort, token.Uint, token.Ulong,
				token.Float32, token.Float64, token.Float, token.Double:
				w.buffer.WriteString(" = 0")

			case token.Char:
				w.buffer.WriteString(" = U''")

			case token.Bool:
				w.buffer.WriteString(" = false")

			case token.String:
				w.buffer.WriteString(` = ""`)
			}
		}
		w.buffer.WriteString(";\n")

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
				_, isTypeName := t.ReturnType.(*types.TypeName)
				if isTypeName {
					w.buffer.WriteString("std::shared_ptr<")
				}
				writeType(t.ReturnType, w)
				if isTypeName {
					w.buffer.WriteString(">")
				}
				w.buffer.WriteString(" ")
			}
		}
		w.buffer.WriteString(t.Identifier())
		writeType(t.Parameters, w)
		if t.ClassName != "" && t.Body == nil {
			w.buffer.WriteString(" = 0")
		}
		w.buffer.WriteString(";\n")

	case *declaration.Enum:
		w.buffer.WriteString("enum class " + t.Identifier() + "\n")
		writeIndent(indent, w)
		w.buffer.WriteString("{\n")
		for i, m := range t.Members {
			writeIndent(indent+tabSize, w)
			v := m.(*declaration.Variable)
			w.buffer.WriteString(v.Identifier())
			if v.Value != nil {
				w.buffer.WriteString(" = ")
				writeExpression(v.Value, w)
			}
			if i == len(t.Members)-1 {
				w.buffer.WriteString("\n")
			} else {
				w.buffer.WriteString(",\n")
			}
		}
		w.buffer.WriteString("};\n")

	case *declaration.Interface:
		if t.TypeParameters != nil {
			writeType(t.TypeParameters, w)
			writeIndent(indent, w)
		}
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
			writeDeclaration(p, m, indent+tabSize, w)
		}
		w.buffer.WriteString("};\n")

	case *declaration.Class:
		if w.replaceClasses[p.Namespace+"."+t.Identifier()] != nil {
			break
		}
		if t.TypeParameters != nil {
			writeType(t.TypeParameters, w)
			writeIndent(indent, w)
		}
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
		hasDefaultConstructor := false
		for _, m := range t.Members {
			if f, ok := m.(*declaration.Function); ok {
				if f.Identifier() == t.Identifier() {
					if f.Parameters == nil || len(f.Parameters.Parameters) == 0 {
						hasDefaultConstructor = true
						break
					}
				}
			}
		}
		//TO-DO sort and add destructor
		if !hasDefaultConstructor {
			writeIndent(indent+tabSize, w)
			w.buffer.WriteString(t.Identifier() + "();\n")
		}
		for _, m := range t.Members {
			writeDeclaration(p, m, indent+tabSize, w)
		}
		w.buffer.WriteString("};\n")
	}
}

func writeClass(c *declaration.Class, w *writer) {
	first := true
	hasDefaultConstructor := false
	for _, m := range c.Members {
		if f, ok := m.(*declaration.Function); ok {
			if f.Identifier() == c.Identifier() {
				if f.Parameters == nil || len(f.Parameters.Parameters) == 0 {
					hasDefaultConstructor = true
					break
				}
			}
		}
	}
	if !hasDefaultConstructor {
		if c.TypeParameters != nil {
			writeType(c.TypeParameters, w)
		}
		w.buffer.WriteString(c.Identifier())
		if c.TypeParameters != nil {
			writeTypeArguments(c.TypeParameters, w)
		}
		w.buffer.WriteString("::" + c.Identifier() + "()\n{\n}\n")
		first = false
	}
	for _, m := range c.Members {
		if f, ok := m.(*declaration.Function); ok {
			if !first {
				w.buffer.WriteString("\n")
			}
			writeFunction(f, c, w)
			first = false
		}
	}
}

func writeFunction(f *declaration.Function, c *declaration.Class, w *writer) {
	if f.Body == nil {
		// TO-DO only declare, replace with cpp function
		return
	}
	if f.TypeParameters != nil {
		writeType(f.TypeParameters, w)
	}
	if c != nil && c.TypeParameters != nil {
		writeType(c.TypeParameters, w)
	}
	if !(f.ClassName != "" && (f.Identifier() == f.ClassName || f.Identifier() == "~"+f.ClassName)) {
		if f.ReturnType == nil {
			w.buffer.WriteString("void ")
		} else {
			_, isTypeName := f.ReturnType.(*types.TypeName)
			if isTypeName {
				w.buffer.WriteString("std::shared_ptr<")
			}
			writeType(f.ReturnType, w)
			if isTypeName {
				w.buffer.WriteString(">")
			}
			w.buffer.WriteString(" ")
		}
	}
	if f.ClassName != "" {
		if c != nil && c.TypeParameters != nil {
			writeType(c.TypeParameters, w)
		}
		w.buffer.WriteString(f.ClassName)
		if c.TypeParameters != nil {
			writeTypeArguments(c.TypeParameters, w)
		}
		w.buffer.WriteString("::")
	}
	w.buffer.WriteString(f.Identifier())
	writeType(f.Parameters, w)
	w.buffer.WriteString("\n")
	writeStatement(f.Body, 0, w)
}
