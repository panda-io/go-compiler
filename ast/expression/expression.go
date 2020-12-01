package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

type Expression interface {
	node.Node
	expression()
}

type Base struct {
	node.Base
}

func (*Base) expression() {}

type Literal struct {
	Base
	Type  token.Token // (identifier, string, char, float, int, bool)
	Value string
}

type Identifier struct {
	Base
	Name string
}

type This struct {
	Base
}

type Super struct {
	Base
}

type Parentheses struct {
	Base
	Expression Expression
}

type MemberAccess struct {
	Base
	Parent Expression
	Member *Identifier
}

type Subscripting struct {
	Base
	Parent  Expression
	Element Expression
}

type Invocation struct {
	Base
	Function  Expression
	Arguments *types.Arguments
}

type New struct {
	Base
	Type      types.Type
	Arguments *types.Arguments
}

type Increment struct {
	Base
	Expression Expression
}

type Decrement struct {
	Base
	Expression Expression
}

type Unary struct {
	Base
	Operator   token.Token
	Expression Expression
}

type Binary struct {
	Base
	Left     Expression
	Operator token.Token
	Right    Expression
}

type Conditional struct {
	Base
	Condition Expression
	First     Expression
	Second    Expression
}
