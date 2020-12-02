package statement

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Statement interface {
	node.Node
	statement()
}

type Base struct {
	node.Base
}

func (*Base) statement() {}

type Empty struct {
	Base
}

type Raw struct {
	Base
	Source string
}

type Expression struct {
	Base
	Expression expression.Expression
}

type Declaration struct {
	Base
	Name  *expression.Identifier
	Type  types.Type
	Value expression.Expression
}

type Return struct {
	Base
	Expression expression.Expression
}

type Continue struct {
	Base
}

type Break struct {
	Base
}

type Try struct {
	Base
	Try     Statement
	Operand *types.Parameters
	Catch   Statement
	Finally Statement
}

type Throw struct {
	Base
	Expression expression.Expression
}

type If struct {
	Base
	Initialization Statement
	Condition      Statement
	Body           Statement
	Else           Statement
}

type Switch struct {
	Base
	Initialization Statement
	Operand        Statement
	Body           []*Case
}

type Case struct {
	Base
	Case expression.Expression
	Body Statement
}

type For struct {
	Base
	Initialization Statement
	Condition      Statement
	Post           Statement
	Body           Statement
}

type Foreach struct {
	Base
	Key      Statement
	Item     Statement
	Iterator expression.Expression
	Body     Statement
}

type Compound struct {
	Base
	Statements []Statement
}
