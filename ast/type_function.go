package ast

import "github.com/panda-foundation/go-compiler/ir"

type TypeFunction struct {
	TypeBase
	ReturnType Type
	Parameters []Type
	Ref        []bool
}

func (f *TypeFunction) Type(p *Program) ir.Type {
	var types []ir.Type
	for i, param := range f.Parameters {
		types = append(types, GetIRType(param, p, f.Ref[i]))
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = GetIRType(f.ReturnType, p, false)
	}
	return ir.NewPointerType(ir.NewFuncType(t, types...))
}

func (f *TypeFunction) Define(p *Program) *ir.Func {
	var params []*ir.Param
	if len(f.Parameters) > 0 {
		for i, parameter := range f.Parameters {
			var param *ir.Param
			switch t := parameter.(type) {
			case *BuitinType:
				param = ir.NewParam(GetIRType(parameter, p, f.Ref[i]))
				param.Ref = f.Ref[i]
				param.Builtin = true

			case *TypeName:
				userData, d := p.FindDeclaration(t)
				switch d.(type) {
				case *Class:
					param = ir.NewParam(pointerType)

				case *Enum:
					param = ir.NewParam(ir.I32)

				case *Interface:
					// TO-DO interface
					//TO-DO need to be some convert
					param = ir.NewParam(pointerType)
				}
				SetUserData(param, userData)

			case *TypeFunction:
				// TO-DO testing~
				param = ir.NewParam(t.Type(p))
			}

			params = append(params, param)
		}
	}

	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = GetIRType(f.ReturnType, p, false)
	}

	return ir.NewFunc("", t, params...)
}
