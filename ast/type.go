package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

var (
	builtinTypes = []string{"global.counter", "golbal.allocator", "global.string"}
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

func IsBuiltinType(qualified string) bool {
	for _, str := range builtinTypes {
		if str == qualified {
			return true
		}
	}
	return false
}
