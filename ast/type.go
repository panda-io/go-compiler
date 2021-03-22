package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

var (
	builtinClasses = []string{"global.counter", "golbal.allocator", "global.string"}
)

type Type interface {
	Node
}

type TypeBase struct {
	NodeBase
}

func GetIRType(typ Type, p *Program, ref bool) ir.Type {
	switch t := typ.(type) {
	case *BuitinType:
		if ref {
			return ir.NewPointerType(t.Type())
		}
		return t.Type()

	case *TypeName:
		_, d := p.FindDeclaration(t)
		switch d.(type) {
		case *Class:
			return pointerType

		case *Enum:
			return ir.I32

		case *Interface:
			return pointerType
		}

	case *TypeFunction:
		return t.Type(p)
	}

	return nil
}

func GetIRParam(parameter *Parameter, p *Program) *ir.Param {
	var param *ir.Param
	switch t := parameter.Type.(type) {
	case *BuitinType:
		param = ir.NewParam(GetIRType(t, p, parameter.Ref))
		param.Ref = parameter.Ref
		param.Builtin = true

	case *TypeName:
		userData, _ := p.FindDeclaration(t)
		param = ir.NewParam(GetIRType(t, p, parameter.Ref))
		SetUserData(param, userData)

	case *TypeFunction:
		param = ir.NewParam(GetIRType(t, p, parameter.Ref))
	}
	param.LocalName = parameter.Name

	return param
}

func CreateStruct(qualified string) *ir.StructType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return t
}

func CreateStructPointer(qualified string) *ir.PointerType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return ir.NewPointerType(t)
}

func IsBuiltinClass(qualified string) bool {
	for _, str := range builtinClasses {
		if str == qualified {
			return true
		}
	}
	return false
}

func CopyUserData(source, dest ir.Value) {
	t1 := source.Type()
	t2 := dest.Type()
	if ir.IsPointer(t1) && ir.IsPointer((t2)) {
		t2.(*ir.PointerType).UserData = t1.(*ir.PointerType).UserData
	}
}

func GetUserData(value ir.Value) string {
	t := value.Type()
	if ir.IsPointer(t) {
		return t.(*ir.PointerType).UserData
	}
	return ""
}

func SetUserData(value ir.Value, userData string) {
	t := value.Type()
	if ir.IsPointer(t) {
		t.(*ir.PointerType).UserData = userData
	}
}
