package expression

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type ExpressionType int

const (
	NoneValue ExpressionType = iota
	IntValue
	FloatValue
	StringValue
	ObjectValue
)

// TO-DO add validate later
type Expression interface {
	node.Node
	Type() *ExpressionType
	TypeName() string
	GenerateIR() ir.Value
}

type Base struct {
	node.Base
}

func (*Base) GenerateIR() {}

type Literal struct {
	Base
	Type  token.Token // (identifier, string, char, float, int, bool)
	Value string
}

func (e *Literal) GenerateIR() ir.Value {
	switch e.Type {
	case token.STRING:
		//TO-DO

	case token.CHAR:
		//TO-DO

	case token.FLOAT:
		return ir.NewFloatFromString(ir.Float32, e.Value)

	case token.INT:
		return ir.NewIntFromString(ir.I32, e.Value)

	case token.BOOL:
		return ir.NewIntFromString(ir.I1, e.Value)

	case token.IDENT:
		//TO-DO
	}
	return nil
}

type Identifier struct {
	Base
	Name string
}

func (e *Identifier) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type This struct {
	Base
}

func (e *This) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Super struct {
	Base
}

func (e *Super) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Parentheses struct {
	Base
	Expression Expression
}

func (e *Parentheses) GenerateIR() ir.Value {
	return e.Expression.GenerateIR()
}

type MemberAccess struct {
	Base
	Parent        Expression
	Member        *Identifier
	FullNamespace string
}

func (e *MemberAccess) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Subscripting struct {
	Base
	Parent  Expression
	Element Expression
}

func (e *Subscripting) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Invocation struct {
	Base
	Function  Expression
	Arguments *types.Arguments
}

func (e *Invocation) GenerateIR() ir.Value {
	args := []ir.Value{}
	for _, arg := range e.Arguments.Arguments {
		args = append(args, arg.(types.Type).GenerateIR())
	}
	return ir.NewCall(e.Function.GenerateIR(), args...)
}

type New struct {
	Base
	Type      types.Type
	Arguments *types.Arguments
}

func (e *New) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Increment struct {
	Base
	Expression Expression
}

func (e *Increment) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Decrement struct {
	Base
	Expression Expression
}

func (e *Decrement) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Unary struct {
	Base
	Operator   token.Token
	Expression Expression
}

func (e *Unary) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Binary struct {
	Base
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (e *Binary) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

type Conditional struct {
	Base
	Condition Expression
	First     Expression
	Second    Expression
}

func (e *Conditional) GenerateIR() ir.Value {
	//TO-DO
	return nil
}
