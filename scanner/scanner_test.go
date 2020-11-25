package scanner

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

func handleError(p *token.Position, message string) {
	fmt.Println(p.String())
	panic(message)
}

func TestBasic(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`identifier 100 100.0 "string" 'c' @`))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "identifier")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.INT)
	assertEqual(t, literal, "100")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.FLOAT)
	assertEqual(t, literal, "100.0")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.STRING)
	assertEqual(t, literal, `"string"`)
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.CHAR)
	assertEqual(t, literal, "'c'")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.META)
	assertEqual(t, literal, "@")

	s.ScanFile(f, []byte("`raw string`"))
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.STRING)
	assertEqual(t, literal, "`raw string`")

	s.ScanFile(f, []byte("/* block comment *///comment\nnewline"))
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "newline")
}

func TestNumber(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("0xabcd 0o1234 0b1010 0"))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.INT)
	assertEqual(t, literal, "0xabcd")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.INT)
	assertEqual(t, literal, "0o1234")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.INT)
	assertEqual(t, literal, "0b1010")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.INT)
	assertEqual(t, literal, "0")
}

func TestUnicode(t *testing.T) {
}

func TestUnicodeBom(t *testing.T) {
}

func TestPreprocessor(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, []string{"windows", "opengl"})

	src := `
	#if windows
	windows
	#elif macos
	macos
	#elif linux
	linux
	#end
	
	#if none_windows
	none_windows
	#else
	windows
	#end

	end
	`
	s.ScanFile(f, []byte(src))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "windows")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "windows")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "end")
}

func TestEscape(t *testing.T) {

}
func TestOperator(t *testing.T) {

}
