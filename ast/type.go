package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Type interface {
	Node
	Type(*Program) ir.Type
}

type TypeBase struct {
	NodeBase
}

func CreateStructType(qualified string) *ir.StructType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return t
}

func CreateStructPointerType(qualified string) *ir.PointerType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return ir.NewPointerType(t)
}
