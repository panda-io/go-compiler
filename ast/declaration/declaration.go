package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

type Declaration interface {
	node.Node
	GenerateIR(*node.Context)
	GenerateIRDeclaration() ir.Value
	Identifier() string
	SetQualifiedName(string)
}

type Modifier struct {
	Public bool
	//Inline
}

func (m0 *Modifier) Equal(m1 *Modifier) bool {
	return m0.Public == m1.Public
}

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*expression.Literal
}

type Base struct {
	node.Base
	Attributes     []*Attribute
	Modifier       *Modifier
	Name           *expression.Identifier
	QualifinedName string
}

func (b *Base) SetQualifiedName(qualifinedName string) {
	b.QualifinedName = qualifinedName
}

func (b *Base) Identifier() string {
	return b.Name.Name
}
