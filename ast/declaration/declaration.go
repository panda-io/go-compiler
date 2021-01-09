package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Declaration interface {
	node.Node
	Identifier() string
	Qualified(namespace string) string
}

type Modifier struct {
	Public bool
	//TO-DO Inline
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
	Attributes []*Attribute
	Modifier   *Modifier
	Name       *expression.Identifier
	ObjectName string // parent object (class|interface|enum)
}

func (b *Base) External() bool {
	for _, a := range b.Attributes {
		if a.Name == node.Extern {
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
	if b.External() {
		if b.ObjectName != "" {
			panic("object member cannot be external")
		}
	} else {
		if b.ObjectName != "" {
			name = b.ObjectName + "." + name
		}
		if !(namespace == node.Global && b.Name.Name == node.Entry) {
			name = namespace + "." + name
		}
	}
	return name
}

func FindeDeclaration(c *node.Context, declarations map[string]Declaration, t *types.TypeName) Declaration {
	if t.Selector == "" {
		// search current package
		if c.Namespace != node.Global {
			qualified := c.Namespace + "." + t.Name
			if declarations != nil {
				return declarations[qualified]
			}
		}
		// search global
		qualified := node.Global + "." + t.Name
		return declarations[qualified]
	} else {
		// search imports
		for _, i := range c.Imports {
			if i.Alias == t.Selector {
				return declarations[i.Namespace+"."+t.Name]
			}
		}
		return nil
	}
}
