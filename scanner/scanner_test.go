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

	s.ScanFile(f, []byte(`identifier 100 100.0 "string" 'c' @ ++ .5 .`))
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
	assertEqual(t, literal, `'c'`)
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.META)
	assertEqual(t, literal, "@")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.PlusPlus)
	assertEqual(t, literal, "++")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.FLOAT)
	assertEqual(t, literal, ".5")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.Dot)
	assertEqual(t, literal, ".")

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

func TestPreprocessor(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, []string{"windows", "opengl"})

	src := `
	#if windows
	windows
		#if opengl
		opengl
		#elif directx
		directx
		#else
		none
		#end
	#elif macos
	macos
		#if opengl
		opengl
		#elif directx
		directx
		#else
		none
		#end
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
	assertEqual(t, literal, "opengl")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "windows")
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "end")
}

func TestEscape(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`"hello\n\r\x1b\123\u1234\U0001FFFF"`))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.STRING)
	assertEqual(t, literal, `"hello\n\r\x1b\123\u1234\U0001FFFF"`)
}

func TestCharEscape(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`'c' '\x1b' '\123' '\U0001FFFF'`))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.CHAR)
	assertEqual(t, literal, `'c'`)
	_, tok, literal = s.Scan()
	assertEqual(t, tok, token.CHAR)
	assertEqual(t, literal, `'\x1b'`)
}

func TestUnicade(t *testing.T) {
	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("\uFEFFhello\n"))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "hello")
	_, tok, literal = s.Scan()
}

func TestNullChar(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("\\0 did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte{0})
	s.Scan()
}

func TestInvalidUTF8Char(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid utf-8 did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("\x80\uFEFF"))
	s.Scan()
}

func TestInvalidBom(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid utf bom did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("abc\uFEFF"))
	_, tok, literal := s.Scan()
	assertEqual(t, tok, token.IDENT)
	assertEqual(t, literal, "abc")
	s.Scan()
}

func TestUnterminated(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("unterminated comment did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("/* comment"))
	s.Scan()
}

func TestInvalidNumber(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid number did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("00"))
	s.Scan()
}

func TestInvalidHexNumber(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid hex did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("0x"))
	s.Scan()
}

func TestInvalidRadixPoint(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid radix point did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("0xabc.0"))
	s.Scan()
}

func TestInvalidFloat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid float did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte("100."))
	s.Scan()
}

func TestInvalidEscape(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid number did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`"\`))
	s.Scan()
}

func TestInvalidNumberEscape(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid number did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`"\0FF"`))
	s.Scan()
}

func TestInvalidUtf8Escape(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid uft escape did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`"\U00FFFFFF"`))
	s.Scan()
}

func TestUnterminatedString(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("unterminated string did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`"abc`))
	s.Scan()
}

func TestUnterminatedChar(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("unterminated char did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`'`))
	s.Scan()
}

func TestInvalidCharLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid char did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`'\abc'`))
	s.Scan()
}

func TestInvalidChar(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid char did not panic")
		}
	}()

	fs := &token.FileSet{}
	f := fs.AddFile("file.pd", 100)
	s := NewScanner(handleError, nil)

	s.ScanFile(f, []byte(`'\xxx'`))
	s.Scan()
}
