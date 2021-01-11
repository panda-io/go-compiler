package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Literal struct {
	Base
	Type  token.Token // (identifier, string, char, float, int, bool)
	Value string
}

func (e *Literal) GenerateIR(c *node.Context) ir.Value {
	//	case token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL, token.NULL, token.Void:
	switch e.Type {
	case token.STRING:
		return ir.NewCharArray([]byte(e.Value))

	case token.CHAR:
		//TO-DO
		return nil

	case token.FLOAT:
		return ir.NewFloatFromString(ir.Float32, e.Value)

	case token.INT:
		return ir.NewIntFromString(ir.I32, e.Value)

	case token.BOOL:
		return ir.NewIntFromString(ir.I1, e.Value)

	case token.NULL:
		return ir.NewNull(nil)

	default:
		return nil
	}
}
