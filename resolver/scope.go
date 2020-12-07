package resolver

import (
	"github.com/panda-foundation/go-compiler/ast/types"
)

//TO-DO scope check
// use scope when start compound, if, for, switch, case, function

type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}

func NewScope(outer *Scope) *Scope {
	return &Scope{outer, make(map[string]*Object)}
}

func (s *Scope) Find(name string) *Object {
	if obj, ok := s.Objects[name]; ok {
		return obj
	} else if s.Outer != nil {
		return s.Outer.Find(name)
	}
	return nil
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
	VariableObject ObjectKind = iota
	ClassObject
	InterfaceObject
	EnumObject
	FunctionObject
)
