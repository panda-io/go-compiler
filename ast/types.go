package ast

import "github.com/panda-foundation/go-compiler/token"

// Type struct
type Type interface {
	Equal(Type) bool
}

// BuitinType basic type
type BuitinType struct {
	Position int
	Token    token.Token
}

// Equal compare if two type are same
func (t0 *BuitinType) Equal(t Type) bool {
	if t1, ok := t.(*BuitinType); ok {
		return t0.Token == t1.Token
	}
	return false
}

// TypeName qualified name
type TypeName struct {
	Position      int
	QualifiedName []*Identifier
	TypeArguments *TypeArguments
}

// Equal compare if two type are same
func (t0 *TypeName) Equal(t Type) bool {
	t1, ok := t.(*TypeName)
	if !ok {
		return false
	}
	if len(t0.QualifiedName) != len(t1.QualifiedName) {
		return false
	}
	for i, n := range t0.QualifiedName {
		if n.Name != t1.QualifiedName[i].Name {
			return false
		}
	}
	if t0.TypeArguments != nil && t1.TypeArguments != nil {
		return t0.TypeArguments.Equal(t1.TypeArguments)
	}
	if t0.TypeArguments == nil && t1.TypeArguments == nil {
		return true
	}
	return false
}

// TypeArguments type arguments
type TypeArguments struct {
	Position  int
	Arguments []Type
}

// Equal compare if two type are same
func (t0 *TypeArguments) Equal(t Type) bool {
	t1, ok := t.(*TypeArguments)
	if !ok {
		return false
	}
	if len(t0.Arguments) != len(t1.Arguments) {
		return false
	}
	for i, arg := range t0.Arguments {
		if !arg.Equal(t1.Arguments[i]) {
			return false
		}
	}
	return true
}

// TypeParameters type parameters
type TypeParameters struct {
	Position   int
	Parameters []*TypeParameter
}

// Equal compare if two type are same
func (t0 *TypeParameters) Equal(t Type) bool {
	t1, ok := t.(*TypeParameters)
	if !ok {
		return false
	}
	if len(t0.Parameters) != len(t1.Parameters) {
		return false
	}
	for i, p := range t0.Parameters {
		if !p.Equal(t1.Parameters[i]) {
			return false
		}
	}
	return true
}

// TypeParameter type parameter
type TypeParameter struct {
	Name *Identifier
	Type Type
}

// Equal compare if two type parameter are same
func (t0 *TypeParameter) Equal(t Type) bool {
	t1, ok := t.(*TypeParameter)
	if !ok {
		return false
	}
	if t0.Name.Name != t1.Name.Name {
		return false
	}
	if t0.Type != nil && t1.Type != nil {
		return t0.Type.Equal(t1.Type)
	}
	if t0.Type == nil && t1.Type == nil {
		return true
	}
	return false
}
