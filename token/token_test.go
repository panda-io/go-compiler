package token

import (
	"fmt"
	"reflect"
	"testing"
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

func TestTokens(t *testing.T) {
	assertEqual(t, Assign.String(), "=")
	assertEqual(t, Token(1000).String(), "")
	assertEqual(t, ReadToken("i16"), Int16)
	assertEqual(t, ReadToken("Ident"), IDENT)
	assertEqual(t, ReadToken("await"), Await)
	assertEqual(t, ReadToken("++"), PlusPlus)
	assertEqual(t, ReadToken("true"), BOOL)
	assertEqual(t, ReadToken("false"), BOOL)
	assertEqual(t, ReadToken("null"), NULL)
}

func TestTypes(t *testing.T) {
	assertEqual(t, Async.IsKeyword(), true)
	assertEqual(t, IDENT.IsLiteral(), true)
	assertEqual(t, String.IsScalar(), true)
	assertEqual(t, Or.IsOperator(), true)
	assertEqual(t, String.CppType(), "std::string")
	assertEqual(t, META.Precedence(), 0)
	assertEqual(t, Or.Precedence(), 1)
	assertEqual(t, And.Precedence(), 2)
	assertEqual(t, Equal.Precedence(), 3)
	assertEqual(t, Plus.Precedence(), 4)
	assertEqual(t, Mul.Precedence(), 5)
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

func TestPosition(t *testing.T) {
	s := &FileSet{}
	f1 := s.AddFile("file1.pd", 100)
	f1.AddLine(10)
	f1.AddLine(20)
	f2 := s.AddFile("file2.pd", 100)
	f2.AddLine(10)
	f2.AddLine(20)
	f3 := s.AddFile("input1", 100)
	f3.AddLine(10)
	f3.AddLine(20)

	assertEqual(t, f1.Position(12).String(), "file1.pd:2:3")
	assertEqual(t, s.File(112), f2)
	assertEqual(t, s.Position(112).file, f2)
	assertEqual(t, s.File(1000), nil)
	assertEqual(t, s.Position(1000), nil)
}

func TestAddSameFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("add same files did not panic")
		}
	}()

	s := &FileSet{}
	f1 := s.AddFile("file.pd", 100)
	f1.AddLine(10)
	f1.AddLine(20)
	s.AddFile("file.pd", 100)
}
