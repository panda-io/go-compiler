package resolver

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/types"
)

//TO-DO scope check
// open scope when start compound, if, for, switch, case, function

type Scope struct {
	outer   *Scope
	objects map[string]*Object
}

func NewScope(outer *Scope) *Scope {
	return &Scope{outer, make(map[string]*Object)}
}

func (s *Scope) OpenScope() *Scope {
	return NewScope(s)
}

func (s *Scope) CloseScope() *Scope {
	return s.outer
}

func (s *Scope) Find(name string) *Object {
	if obj, ok := s.objects[name]; ok {
		return obj
	} else if s.outer != nil {
		return s.outer.Find(name)
	}
	return nil
}

func (s *Scope) Insert(obj *Object) bool {
	if _, ok := s.objects[obj.Name]; ok {
		return false
	}
	s.objects[obj.Name] = obj
	return true
}

type Object struct {
	Kind        ObjectKind
	Name        string
	Public      bool
	Declaration declaration.Declaration
	Type        types.Type
}

func NewObject(kind ObjectKind, name string) *Object {
	return &Object{Kind: kind, Name: name}
}

type ObjectKind int

const (
	VariableObject ObjectKind = iota
	FunctionObject
	EnumObject
	InterfaceObject
	ClassObject
)
