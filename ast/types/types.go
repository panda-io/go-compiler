package types

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Type interface {
	node.Node
	GenerateIR() ir.Value
}

type Base struct {
	node.Base
}

func (*Base) GenerateIR() ir.Value { return nil }

type BuitinType struct {
	Base
	Token token.Token
}

type TypeName struct {
	Base
	Name          string
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
