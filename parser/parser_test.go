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
	p.ParseStatementBlock([]byte("{;var i int = 1; var j float = 1.0; i++; j = i + 1; return i;}"))
	p.ParseStatementBlock([]byte("{{var i int = 0;}}"))
	p.ParseStatementBlock([]byte("{for(;;){break;}}"))
	p.ParseStatementBlock([]byte("{for(var i = 0; i < 10; i++){break;}}"))
	p.ParseStatementBlock([]byte("{if (a == 1) {} else if (a == 2) {} else {}}"))
	p.ParseStatementBlock([]byte("{while(true){}}"))
	p.ParseStatementBlock([]byte("{switch(a.b.c){case 1:{} case 2:{} default:{}}}"))
	p.ParseStatementBlock([]byte("{foreach(var item : data){}}"))
	p.ParseStatementBlock([]byte("{foreach(var key; var value : data){}}"))
	p.ParseStatementBlock([]byte("{try{}catch(e exception){}finally{}}"))
	p.ParseStatementBlock([]byte("{throw \"some message\";}"))
	p.ParseStatementBlock([]byte("{@\"raw source\"}"))
	p.ParseStatementBlock([]byte("{yield}")) //TO-DO
	p.ParseStatementBlock([]byte("{await}")) //TO-DO
}

func TestStatementFail1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no effect statement did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseStatementBlock([]byte("{1+1;}"))
}

func TestStatementFail2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("none declare statement did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseStatementBlock([]byte("{for(1; i < 10; i++){break;}}"))
}

func TestStatementFail3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no effect statement did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseStatementBlock([]byte("{a.b.c;}"))
}

func TestStatementFail4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("raw statement did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseStatementBlock([]byte("{@what}"))
}

func TestStatementFail5(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("if else statement did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseStatementBlock([]byte("{if (true) {} else do_something();}"))
}

func TestDeclaration(t *testing.T) {
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; public interface x<type> { print(); } public class a {} public class b<type> : a, x<type> { public var e int = 100; public function print<t>() void {} function ~b(){}}"))
	p.ParseBytes([]byte("namespace; public enum test { blue, yello, red = 10 }"))
	p.ParseBytes([]byte("namespace; interface a {} interface b : a {}"))
}

func TestDeclarationFail1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("var init did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; var a int = gen_int(); function gen_int() int { return 1; }"))
}

func TestDeclarationFail2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("destructor did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; function ~gen_int() int { return 1; }"))
}

func TestDeclarationFail3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("destructor did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class test { function ~test1(){} }"))
}

func TestDeclarationFail4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("enum did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; public enum test { blue, yello, red = \"red\" }"))
}

func TestDeclarationFail5(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("interface redeclare did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; interface a { b() int; b() int}"))
}

func TestDeclarationFail6(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class redeclare did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class a { function b() int; function b() int}"))
}

func TestDeclarationFail7(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class redeclare did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class a { var b int; var b int; }"))
}

func TestDeclarationFail8(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("class unexpected did not panic")
		}
	}()
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("namespace; class a { enum b {} }"))
}
func TestNamespace(t *testing.T) {
	p := NewParser([]string{"cpp"})
	p.ParseBytes([]byte("@doc \"package document here\" \nnamespace first.second.third;"))
	assertEqual(t, p.root.Children["first"].Package, "first")
	assertEqual(t, p.root.Children["first"].Children["second"].Package, "second")
}
