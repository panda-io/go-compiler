package parser

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

func TestStatement(t *testing.T) {
	p := NewParser([]string{"cpp"})
	p.ParseCompoundStatement([]byte("{i++;}"))
	p.ParseCompoundStatement([]byte("{;var i int = 1; var j float = 1.0; i++; j = i + 1; return i;}"))
	p.ParseCompoundStatement([]byte("{{var i int = 0;}}"))
	p.ParseCompoundStatement([]byte("{for(;;){break;}}"))
	p.ParseCompoundStatement([]byte("{for{break;}}"))
	p.ParseCompoundStatement([]byte("{for(var i = 0; i < 10; i++){break;}}"))
	p.ParseCompoundStatement([]byte("{for(var i = 0; i < 10){ i++; }}"))
	p.ParseCompoundStatement([]byte("{for(var item : items) { print(item); }}"))
	p.ParseCompoundStatement([]byte("{for(var key; var value : values) { print(key); }}"))
	p.ParseCompoundStatement([]byte("{if (a == 1) {} else if (a == 2) {} else {}}"))
	p.ParseCompoundStatement([]byte("{if (true) {} else do_something();}"))
	p.ParseCompoundStatement([]byte("{switch(a.b.c){case 1: do_something(); case 2:{} default:{}}}"))
	p.ParseCompoundStatement([]byte("{try{}catch(e exception){}finally{}}"))
	p.ParseCompoundStatement([]byte("{throw \"some message\";}"))
	p.ParseCompoundStatement([]byte("{@\"raw source\"}"))
}

func TestStatementFail1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("raw statement did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseCompoundStatement([]byte("{@what}"))
}

func TestDeclaration(t *testing.T) {
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; public interface x<type> { function print(); } public class a {} public class b<type> : a, x<type> { public var e int = 100; public function print<t>() void {} function ~b(){}}"))
	p.ParseBytes([]byte("namespace; public enum test { blue, yello, red = 10 }"))
	p.ParseBytes([]byte("namespace; interface a {} interface b : a {}"))
}

func TestDeclarationFail1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("destructor did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; function ~gen_int() int { return 1; }"))
}

func TestDeclarationFail2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("destructor did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class test { function ~test1(){} }"))
}

func TestDeclarationFail3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("interface redeclare did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; interface a { b() int; b() int}"))
}

func TestDeclarationFail4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class member redeclare did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class a { function b() int; function b() int}"))
}

func TestDeclarationFail5(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class redeclare did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class a { var b int; var b int; }"))
}

func TestNamespace(t *testing.T) {
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("@doc \"package document here\" \nnamespace first.second.third;"))
}

func TestExpression(t *testing.T) {
	p := NewParser([]string{"cpp"})
	p.ParseCompoundStatement([]byte("{ this.call_back(); var a = new vector<int>(); }"))
}
