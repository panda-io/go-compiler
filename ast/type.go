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
