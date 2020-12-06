package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/token"
)

type Type interface {
	node.Node
	types()
}

type Base struct {
	node.Base
}

func (*Base) types() {}

type BuitinType struct {
	Base
	Token token.Token
}

type TypeName struct {
	Base
	QualifiedName string
	TypeArguments *TypeArguments
}

type TypeArguments struct {
	Base
	Arguments []Type
	Ellipsis  int
}

type TypeParameters struct {
	Base
	Parameters []*TypeParameter
	Ellipsis   bool
}

type TypeParameter struct {
	Base
	Name string
	Type Type
}

type Parameters struct {
	Base
	Parameters []*Parameter
	Ellipsis   bool
}

type Parameter struct {
	Base
	Name string
	Type Type
}

type Arguments struct {
	Base
	Arguments []node.Node
	Ellipsis  int
}
