package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Variable struct {
	DeclarationBase
	Token token.Token
	Type  Type
	Value Expression
	Const bool

	IRVariable *ir.Global
}

func (v *Variable) GenerateIR(p *Program) {
	if v.Value != nil {
		v.IRVariable = p.IRModule.NewGlobal(v.Qualified(p.Module.Namespace), v.Type.Type(p))
	} else {
		/* TO-DO
		value := v.Value.GenerateIR(c) // GenerateConstantIR
		if constValue, ok := value.(ir.Constant); ok {
			v.IRVariable = p.IRModule.NewGlobalDef(v.Qualified(c.Module.Namespace), constValue)
		} else {
			p.Error(v.Position, "variable initialize value must be constant type")
		}*/
	}
}
