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
	Source *expression.Literal
}

type Expression struct {
	Base
	Expression expression.Expression
}

type Declaration struct {
	Base
	Name  *expression.Identifier
	Type  types.Type
	Value Expression
}

type Return struct {
	Base
	Expression Expression
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
	Expression Expression
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
	Position int
	Operand  Expression
	Body     []*Case
}

type Case struct {
	Base
	Case Expression
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
	Value    Statement
	Iterator Expression
	Body     Statement
}

type Compound struct {
	Base
	Statements []Statement
}
