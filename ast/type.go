package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

var (
	builtinClasses = []string{"global.counter", "golbal.allocator", "global.string"}
)

type Type interface {
	Node
	Type(*Program) ir.Type
}

type TypeBase struct {
	NodeBase
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
