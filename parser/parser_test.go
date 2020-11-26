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

func TestSample(t *testing.T) {
	p := NewParser([]string{})

	p.ParseFile("../sample/all_types.pd")
}

func TestNamespace(t *testing.T) {
	p := NewParser([]string{})

	p.ParseBytes([]byte("@doc \"package document here\" \nnamespace first.second.third;"))
	assertEqual(t, p.root.Children["first"].Package, "first")
	assertEqual(t, p.root.Children["first"].Children["second"].Package, "second")
}
