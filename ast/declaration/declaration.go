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
}

type Modifier struct {
	Public bool
	Static bool
}

func (m0 *Modifier) Equal(m1 *Modifier) bool {
	return m0.Public == m1.Public && m0.Static == m1.Static
}

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*expression.Literal
}

type Attributes struct {
	Resolved []*Attribute
	Custom   []*Attribute
}

type Base struct {
	node.Base
	Attributes
	Modifier  *Modifier
	Namespace []string
}

func (*Base) declaration() {}

type Variable struct {
	Base
	Token token.Token
	Name  *expression.Identifier
	Type  types.Type
	Value expression.Expression
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

type Enum struct {
	Base
	Name    *expression.Identifier
	Members map[string]*Variable
}

type Interface struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Functions      map[string]*Function
}

type Class struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Variables      map[string]*Variable
	Functions      map[string]*Function
}
