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
	FuncType   *ir.Func
}

func (v *Variable) GenerateIR(p *Program) {
	if v.Value != nil {
		value := v.Value.GenerateConstIR(p, GetIRType(v.Type, p, false))
		if value == nil {
			p.Error(v.Name.Position, "only constant expression is allowed to initialize value")
		} else {
			v.IRVariable = p.IRModule.NewGlobalDef(v.Qualified(p.Module.Namespace), value)
		}
	} else {
		// zero initalize
		// TO-DO class type conversion with auto pointer
		v.IRVariable = p.IRModule.NewGlobalDef(v.Qualified(p.Module.Namespace), ir.NewZeroInitializer(pointerType))
	}
	SetUserData(v.IRVariable, v.Qualified(p.Module.Namespace))
}
