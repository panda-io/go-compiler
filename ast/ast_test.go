package ast

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/panda-foundation/go-compiler/token"
)

func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if isNil(a) && isNil(b) {
		return
	}
	if a != b {
		panic(fmt.Sprintf("expected %v, but got %v", b, a))
	}
}

func TestAll(t *testing.T) {
	(&Variable{}).declaration()
	(&Function{}).declaration()
	(&Enum{}).declaration()
	(&Interface{}).declaration()
	(&Class{}).declaration()

	(&BuitinType{}).expression()
	(&Literal{}).expression()
	(&Identifier{}).expression()
	(&ClassReference{}).expression()
	(&ParenExpression{}).expression()
	(&MemberAccessExpression{}).expression()
	(&ElementAccessExpression{}).expression()
	(&InvocationExpression{}).expression()
	(&ClassCreateExpression{}).expression()
	(&UnaryExpression{}).expression()
	(&BinaryExpression{}).expression()
	(&TernaryExpression{}).expression()

	(&EmptyStatement{}).statement()
	(&BlockStatement{}).statement()
	(&ReturnStatement{}).statement()
	(&BranchStatement{}).statement()
	(&RawStatement{}).statement()
	(&TryStatement{}).statement()
	(&ThrowStatement{}).statement()
	(&DeclarationStatement{}).statement()
	(&IfStatement{}).statement()
	(&AssignStatement{}).statement()
	(&IncreaseDecreaseStatement{}).statement()
	(&SwitchStatement{}).statement()
	(&CaseStatement{}).statement()
	(&WhileStatement{}).statement()
	(&ForStatement{}).statement()
	(&ForeachStatement{}).statement()
}

func TestTypeEqual(t *testing.T) {
	t0 := &BuitinType{Token: token.Int32}
	t1 := &BuitinType{Token: token.Int32}
	t2 := &TypeName{}
	assertEqual(t, t0.Equal(t1), true)
	assertEqual(t, t0.Equal(t2), false)
	assertEqual(t, t2.Equal(t0), false)
}
