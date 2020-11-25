package token

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		panic(fmt.Sprintf("expected %v, but got %v", b, a))
	}
}

func TestTokens(t *testing.T) {
	assertEqual(t, Assign.String(), "=")
	assertEqual(t, Token(1000).String(), "")
	assertEqual(t, ReadToken("i16"), Int16)
	assertEqual(t, ReadToken("Ident"), IDENT)
	assertEqual(t, ReadToken("await"), Await)
	assertEqual(t, ReadToken("++"), PlusPlus)
}

func TestTypes(t *testing.T) {
	assertEqual(t, Async.IsKeyword(), true)
	assertEqual(t, IDENT.IsLiteral(), true)
	assertEqual(t, String.IsScalar(), true)
	assertEqual(t, Or.IsOperator(), true)
	assertEqual(t, String.CppType(), "std::string")
	assertEqual(t, META.Precedence(), 0)
	assertEqual(t, OrOr.Precedence(), 1)
	assertEqual(t, AndAnd.Precedence(), 2)
	assertEqual(t, Equal.Precedence(), 3)
	assertEqual(t, Plus.Precedence(), 4)
	assertEqual(t, Star.Precedence(), 5)
}

func TestOperators(t *testing.T) {
	o, _ := ReadOperator([]byte("+100"))
	assertEqual(t, o, Plus)

	o, _ = ReadOperator([]byte("-="))
	assertEqual(t, o, MinusAssign)

	o, _ = ReadOperator([]byte("abc"))
	assertEqual(t, o, ILLEGAL)

	o, _ = ReadOperator([]byte("+++++"))
	assertEqual(t, o, PlusPlus)
}

func TestCppTypePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("cpp type did not panic")
		}
	}()

	IDENT.CppType()
}
