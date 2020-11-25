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

func TestTokens(t *testing.T) {
	s := &token.FileSet{}
	f := s.AddFile("file1.pd", 100)
	f.AddLine(10)
	f.AddLine(20)
}
