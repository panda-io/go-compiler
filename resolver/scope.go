package parser

import (
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}

func NewScope(outer *Scope) *Scope {
	return &Scope{outer, make(map[string]*Object)}
}

func (s *Scope) Find(name string) *Object {
	return s.Objects[name]
}

func (s *Scope) Insert(obj *Object) (alt *Object) {
	if alt = s.Objects[obj.Name]; alt == nil {
		s.Objects[obj.Name] = obj
	}
	return
}

type Object struct {
	Kind ObjectKind
	Name string
	Type types.Type
}

func NewObject(kind ObjectKind, name string) *Object {
	return &Object{Kind: kind, Name: name}
}

type ObjectKind int

const (
	PackageObject ObjectKind = iota
	VariableObject
	ClassObject
	InterfaceObject
	EnumObject
	FunctionObject
)
