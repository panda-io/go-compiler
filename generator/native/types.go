package native

import (
	"strings"

	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

var (
	cppTypes = [...]string{
		token.Bool:    "bool",
		token.Char:    "char32_t",
		token.Int8:    "int8_t",
		token.Int16:   "int16_t",
		token.Int32:   "int32_t",
		token.Int64:   "int64_t",
		token.Uint8:   "uint8_t",
		token.Uint16:  "uint16_t",
		token.Uint32:  "uint32_t",
		token.Uint64:  "uint64_t",
		token.SByte:   "int8_t",
		token.Short:   "int16_t",
		token.Int:     "int32_t",
		token.Long:    "int64_t",
		token.Byte:    "uint8_t",
		token.Ushort:  "uint16_t",
		token.Uint:    "uint32_t",
		token.Ulong:   "uint64_t",
		token.Float32: "float",
		token.Float64: "double",
		token.Float:   "float",
		token.Double:  "double",
		token.String:  "std::string",
		token.Void:    "void",
	}
)

func writeType(t types.Type, w *writer) {
	switch v := t.(type) {
	case *types.BuitinType:
		w.buffer.WriteString(cppTypes[v.Token])

	case *types.TypeName:
		w.buffer.WriteString(strings.ReplaceAll(v.QualifiedName, ".", "::"))
		if v.TypeArguments != nil {
			writeType(v.TypeArguments, w)
		}

	case *types.TypeArguments:
		w.buffer.WriteString("<")
		for i, arg := range v.Arguments {
			if i != 0 {
				w.buffer.WriteString(", ")
			}
			writeType(arg, w)
			if i == v.Ellipsis {
				w.buffer.WriteString("...")
			}
		}
		w.buffer.WriteString(">")

	case *types.TypeParameters:
		w.buffer.WriteString("template <")
		for i, t := range v.Parameters {
			if i > 0 {
				w.buffer.WriteString(", ")
			}
			if v.Ellipsis && i == len(v.Parameters)-1 {
				w.buffer.WriteString("class ... " + t.Name)
			} else {
				w.buffer.WriteString("class " + t.Name)
			}
		}
		w.buffer.WriteString(">\n")

	case *types.Parameters:
		w.buffer.WriteString("(")
		if v != nil {
			for i, param := range v.Parameters {
				if i != 0 {
					w.buffer.WriteString(", ")
				}
				writeType(param.Type, w)
				if v.Ellipsis && i == len(v.Parameters)-1 {
					w.buffer.WriteString("...")
				}
				w.buffer.WriteString(" " + param.Name)
			}
		}
		w.buffer.WriteString(")")

	case *types.Arguments:
		w.buffer.WriteString("(")
		if v != nil {
			for i, arg := range v.Arguments {
				if i != 0 {
					w.buffer.WriteString(", ")
				}
				writeExpression(arg.(expression.Expression), w)
				if i == v.Ellipsis {
					w.buffer.WriteString("...")
				}
			}
		}
		w.buffer.WriteString(")")
	}
}
