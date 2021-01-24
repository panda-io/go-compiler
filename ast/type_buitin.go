package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type BuitinType struct {
	TypeBase
	Token token.Token
}

func (b *BuitinType) Type(*Program) ir.Type {
	switch b.Token {
	case token.Bool:
		return ir.I1

	case token.Char:
		return ir.I32

	case token.Int8, token.SByte:
		return ir.I8

	case token.Uint8, token.Byte:
		return ir.UI8

	case token.Int16, token.Short:
		return ir.I16

	case token.Uint16, token.Ushort:
		return ir.UI16

	case token.Int32, token.Int:
		return ir.I32

	case token.Uint32, token.Uint:
		return ir.UI32

	case token.Int64, token.Long:
		return ir.I64

	case token.Uint64, token.Ulong:
		return ir.UI64

	case token.Float32, token.Float:
		return ir.Float32

	case token.Float64, token.Double:
		return ir.Float64

	case token.Void:
		return ir.Void

	case token.Pointer:
		return ir.NewPointerType(ir.I8)

	default:
		return nil
	}
}
