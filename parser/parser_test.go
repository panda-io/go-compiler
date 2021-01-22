package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/panda-foundation/go-compiler/ast"
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

func TestStatement(t *testing.T) {
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseStatements([]byte("{i++;}"))
	p.ParseStatements([]byte("{;var i int = 1; var j float = 1.0; i++; j = i + 1; return i;}"))
	p.ParseStatements([]byte("{{var i int = 0;}}"))
	p.ParseStatements([]byte("{for(;;){break;}}"))
	p.ParseStatements([]byte("{for{break;}}"))
	p.ParseStatements([]byte("{for(var i = 0; i < 10; i++){break;}}"))
	p.ParseStatements([]byte("{for(var i = 0; i < 10){ i++; }}"))
	p.ParseStatements([]byte("{for(var item : items) { print(item); }}"))
	p.ParseStatements([]byte("{for(var key; var value : values) { print(key); }}"))
	p.ParseStatements([]byte("{if (a == 1) {} else if (a == 2) {} else {}}"))
	p.ParseStatements([]byte("{if (true) {} else do_something();}"))
	p.ParseStatements([]byte("{switch(a.b.c){case 1: do_something(); case 2:{} default:{}}}"))
	p.ParseStatements([]byte("{try{}catch(e exception){}finally{}}"))
	p.ParseStatements([]byte("{throw \"some message\";}"))
	p.ParseStatements([]byte("{@\"raw source\"}"))
}

func TestStatementFail1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("raw statement did not panic")
		}
	}()
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseStatements([]byte("{@what}"))
}

func TestDeclaration(t *testing.T) {
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("@doc \"some package doc\" \nnamespace; var i int = 1; function print() { console.print_line(i); }"))
	p.ParseBytes([]byte("namespace; @meta(a = true, b = \"yes\", c = 1) \npublic enum test { blue, yello, red = 10 }"))
	p.ParseBytes([]byte("namespace; interface ia {} interface ib : ia {}"))
	p.ParseBytes([]byte("namespace; public interface x<type> { function print(); }"))
	p.ParseBytes([]byte("namespace; public class a {} public class b<type> : a, x<type> { public var e int = 100; public function print<t>() void {} function ~b(){}}"))
}

func TestDeclarationFail1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("destructor did not panic")
		}
	}()
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("@doc(\"some package doc\")\nnamespace; function ~gen_int() int { return 1; }"))
}

func TestDeclarationFail2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("destructor did not panic")
		}
	}()
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("namespace; class test { function ~test1(){} }"))
}

func TestDeclarationFail3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("interface redeclare did not panic")
		}
	}()
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("namespace; interface a { b() int; b() int}"))
}

func TestDeclarationFail4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class member redeclare did not panic")
		}
	}()
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("namespace; class a { function b() int; function b() int}"))
}

func TestDeclarationFail5(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class redeclare did not panic")
		}
	}()
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("namespace; class a { var b int; var b int; }"))
}

func TestNamespace(t *testing.T) {
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseBytes([]byte("@doc \"package document here\" \nnamespace first.second.third;"))
}

func TestExpression(t *testing.T) {
	p := NewParser([]string{}, ast.NewProgram())
	p.ParseStatements([]byte("{ this.call_back(); var a = new vector<int>(); }"))
}
