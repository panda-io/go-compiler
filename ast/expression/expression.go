package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
)

/*
type ExpressionType int

const (
	NoneValue ExpressionType = iota
	IntValue
	FloatValue
	StringValue
	ObjectValue
)*/

// TO-DO add validate later
type Expression interface {
	node.Node
	//Type() *ExpressionType //TO-DO
	//TypeName() string //TO-DO
	GenerateIR() ir.Value
	IsConstant() bool
}

type Base struct {
	node.Base
}
