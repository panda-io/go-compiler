package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

type Declaration interface {
	node.Node
	declaration()
	Identifier() string
	IsPublic() bool
}

type Modifier struct {
	Public bool
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
	Namespace  []string
}

func (*Base) declaration() {}

func (b *Base) IsPublic() bool {
	if b.Modifier == nil {
		return false
	}
	return b.Modifier.Public
}

type Variable struct {
	Base
	Token token.Token
	Name  *expression.Identifier
	Type  types.Type
	Value expression.Expression
}

func (v *Variable) Identifier() string {
	return v.Name.Name
}

type Function struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Compound

	Class *Class
}

func (f *Function) Identifier() string {
	return f.Name.Name
}

type Enum struct {
	Base
	Name    *expression.Identifier
	Members []Declaration
}

func (e *Enum) Identifier() string {
	return e.Name.Name
}

type Interface struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration
}

func (i *Interface) Identifier() string {
	return i.Name.Name
}

type Class struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration
}

func (c *Class) Identifier() string {
	return c.Name.Name
}
