package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/token"
)

type Type interface {
	node.Node
	Equal(Type) bool
}

type Base struct {
	node.Base
}

func (*Base) Equal(*Base) bool {
	return false
}

type BuitinType struct {
	Base
	Token token.Token
}

func (t0 *BuitinType) Equal(t Type) bool {
	if t1, ok := t.(*BuitinType); ok {
		return t0.Token == t1.Token
	}
	return false
}

type TypeName struct {
	Base
	QualifiedName []string
	TypeArguments *TypeArguments
}

func (t0 *TypeName) Equal(t Type) bool {
	t1, ok := t.(*TypeName)
	if !ok {
		return false
	}
	if len(t0.QualifiedName) != len(t1.QualifiedName) {
		return false
	}
	for i, name := range t0.QualifiedName {
		if name != t1.QualifiedName[i] {
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

type TypeArguments struct {
	Base
	Arguments []Type
	Ellipsis  int
}

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
	return t0.Ellipsis == t1.Ellipsis
}

type TypeParameters struct {
	Base
	Parameters []*TypeParameter
	Ellipsis   bool
}

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
	return t0.Ellipsis == t1.Ellipsis
}

type TypeParameter struct {
	Base
	Name string
	Type Type
}

func (t0 *TypeParameter) Equal(t Type) bool {
	t1, ok := t.(*TypeParameter)
	if !ok {
		return false
	}
	if t0.Name != t1.Name {
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

type Parameters struct {
	Base
	Parameters []*Parameter
	Ellipsis   bool
}

func (t0 *Parameters) Equal(t Type) bool {
	t1, ok := t.(*Parameters)
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
	return t0.Ellipsis == t1.Ellipsis
}

type Parameter struct {
	Base
	Name string
	Type Type
}

func (t0 *Parameter) Equal(t Type) bool {
	t1, ok := t.(*Parameter)
	if !ok {
		return false
	}
	if t0.Name != t1.Name {
		return false
	}
	return t0.Type.Equal(t1.Type)
}

type Arguments struct {
	Base
	Arguments []node.Node
	Ellipsis  int
}

func (t0 *Arguments) Equal(t Type) bool {
	return false
}
