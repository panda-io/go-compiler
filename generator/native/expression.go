package cpp

import "github.com/panda-foundation/go-compiler/token"

var (
	cppTypes = [...]string{
		token.Bool:    "bool",
		token.Char:    "uint32_t",
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

func CppType(t token.Token) string {
	if t.IsScalar() {
		return cppTypes[t]
	}
	panic("invalid type:" + t.String())
}
