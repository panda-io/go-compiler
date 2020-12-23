package types

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type BuitinType struct {
	Base
	Token token.Token
}

func (t *BuitinType) GenerateIR() ir.Type {
	switch t.Token {
	case token.Bool:
		return ir.I1

	case token.Char:
		return ir.I32

	case token.Int8, token.Uint8, token.SByte, token.Byte:
		return ir.I8

	case token.Int16, token.Uint16, token.Short, token.Ushort:
		return ir.I16

	case token.Int32, token.Uint32, token.Int, token.Uint:
		return ir.I32

	case token.Int64, token.Uint64, token.Long, token.Ulong:
		return ir.I64

	case token.Float32, token.Float:
		return ir.Float32

	case token.Float64, token.Double:
		return ir.Float64

	case token.String:
		panic("TO-DO string is not implemented")

	case token.Void:
		return ir.Void

	default:
		panic("invalid buitin type: " + t.Token.String())
	}
}
