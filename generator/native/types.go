package native

import (
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
		//TO-DO

	case *types.TypeArguments:
		//TO-DO

	case *types.TypeParameters:
		//TO-DO

	case *types.TypeParameter:
		//TO-DO

	case *types.Parameters:
		//TO-DO

	case *types.Parameter:
		//TO-DO

	case *types.Arguments:
		//TO-DO
	}
}

/*
type Scalar struct {
	Start int
	Token Token
}
func (x *Scalar) Print(buffer *bytes.Buffer) {
	x.Token.Print(buffer)
}

type EllipsisLit struct {
	Start int  // position of "..."
	Expr  Expr // ellipsis element type (parameter lists only); or nil
}

func (x *EllipsisLit) Print(buffer *bytes.Buffer) {
	buffer.WriteString("...")
	x.Expr.Print(buffer)
}

type GenericLit struct {
	Start int    // < position
	Types []Expr // <int, int> <T>
}

func (x *GenericLit) Print(buffer *bytes.Buffer) {
	buffer.WriteString("<")
	for i, v := range x.Types {
		if i != 0 {
			buffer.WriteString(", ")
		}
		v.Print(buffer)
	}
	buffer.WriteString(">")
}*/
