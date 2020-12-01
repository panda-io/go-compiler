package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Declaration interface {
	node.Node
	declaration()
}

type Modifier struct {
	Public bool
	Static bool
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
	Modifier
	Attributes
	Namespace []string
}

func (*Base) declaration() {}

type Variable struct {
	Base
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
	Parents        []types.Type
	Functions      map[string]*Function
}

type Class struct {
	Base
	Name           *expression.Identifier
	TypeParameters *types.TypeParameters
	Parents        []types.Type
	Variables      map[string]*Variable
	Functions      map[string]*Function
}
