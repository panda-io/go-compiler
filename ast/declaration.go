package ast

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Declaration interface {
	Node
	Identifier() string
	HasAttribute(attribute string) bool
	Qualified(namespace string) string
}

type Modifier struct {
	Public bool
	Weak   bool
}

func (m *Modifier) Equal(target *Modifier) bool {
	return m.Public == target.Public && m.Weak = target.Weak
}

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*expression.Literal
}

type DeclarationBase struct {
	NodeBase
	Attributes []*Attribute
	Modifier   *Modifier
	Name       *Identifier
	ObjectName string // parent object (class|interface|enum)
}

func (b *Base) HasAttribute(attribute string) bool {
	for _, a := range b.Attributes {
		if a.Name == attribute {
			return true
		}
	}
	return false
}

func (b *Base) Identifier() string {
	return b.Name.Name
}

func (b *Base) Qualified(namespace string) string {
	name := b.Name.Name
	if b.HasAttribute(Extern) {
		if b.ObjectName != "" {
			panic("object member cannot be external")
		}
	} else {
		if b.ObjectName != "" {
			name = b.ObjectName + "." + name
		}
		if !(namespace == Global && b.Name.Name == Entry) {
			name = namespace + "." + name
		}
	}
	return name
}